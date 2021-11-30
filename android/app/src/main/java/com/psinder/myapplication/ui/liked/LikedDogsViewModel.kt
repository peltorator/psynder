package com.psinder.myapplication.ui.liked

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.network.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch


class LikedDogsViewModel : ViewModel() {

    companion object {
        val LOG_TAG = "LikedDogsViewModel"
    }

    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()
    val token: MutableStateFlow<String> = MutableStateFlow("")


    init {
        viewModelScope.launch {
            _viewState.emit(ViewState.Loading)
            Log.d(LOG_TAG, "Start loading liked dogs")
            token.collect { token ->
                Log.d(LOG_TAG, "Token recieved $token")
                val psynas = liked(token)
                Log.d(LOG_TAG, "End loading liked dogs")
                _viewState.emit(ViewState.Data(psynas))
            }
        }
    }

    private suspend fun liked(token: String): List<Psyna> {
        val psynas = safeApiCall(Dispatchers.IO) {
            provideApi().liked(
                bearerToken = "Bearer $token"
            )
        }

        return when (psynas) {
            is ResultWrapper.Success -> {
                psynas.value
            }
            is ResultWrapper.NetworkError -> {
                Log.d("Psynas", "net error")
                emptyList()
            }
            is ResultWrapper.GenericError -> {
                Log.d("Psynas", psynas.code.toString() + psynas.error)
                emptyList()
            }
        }
    }

    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val psynaList: List<Psyna>) : ViewState()
    }
}