package com.psinder.myapplication.ui.home

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.entity.Info
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.provideApi
import com.psinder.myapplication.network.safeApiCall
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.CoroutineExceptionHandler
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

class HomeViewModel : ViewModel() {
    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()

    init {
        loadInfo(AuthRepository.token)
    }

    private fun loadInfo(token: String) {
        viewModelScope.launch {
            _viewState.emit(ViewState.Loading)
            val response = safeApiCall(Dispatchers.IO) {
                provideApi().getInfo(
                    bearerToken = "Bearer $token"
                )
            }

            when (response) {
                is ResultWrapper.Success -> {
                    val info = Info(
                        response.value.users,
                        response.value.shelters,
                        response.value.dogs
                    )
                    _viewState.emit(ViewState.Data(info))
                }
                is ResultWrapper.NetworkError -> {
                    Log.d("Home", "net error")
                }
                is ResultWrapper.GenericError -> {
                    Log.d("Home",  response.code.toString() + response.error)
                }
            }
        }
    }

    fun signOut(coroutineExceptionHandler: CoroutineExceptionHandler) {
        viewModelScope.launch(coroutineExceptionHandler) {
            AuthRepository.logout()
        }
    }

    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val info: Info) : ViewState()
    }
}