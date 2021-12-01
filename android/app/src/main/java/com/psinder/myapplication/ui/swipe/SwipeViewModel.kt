package com.psinder.myapplication.ui.swipe

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.entity.Profile
import com.psinder.myapplication.network.LikeRequest
import com.psinder.myapplication.network.LoadPsynasRequest
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.safeApiCall
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext

class SwipeViewModel: ViewModel() {


    companion object {
        val LOG_TAG = "SwipeViewModel"
    }

    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()

    val token: MutableStateFlow<String> = MutableStateFlow("")

    init {
        viewModelScope.launch {
            _viewState.emit(ViewState.Loading)
            Log.d(LOG_TAG, "Start loading psynas")
            token.collect { token ->
                Log.d(LOG_TAG, "Token recieved ${token}")
                val users = loadPsynas(token)
                Log.d(LOG_TAG, "End loading psynas")
                _viewState.emit(ViewState.Data(users))
            }
        }
    }


    private suspend fun loadPsynas(token : String): List<Profile> {
            val psynas = safeApiCall(Dispatchers.IO) {
                com.psinder.myapplication.network.provideApi().loadpsynas(
                    bearerToken = "Bearer $token"
                )
            }

           when (psynas) {
                is ResultWrapper.Success -> {
                    return psynas.value.map {
                            Profile(
                                age = 5,
                                distance = 100,
                                name = it.name,
                                profile_pic = it.photoLink,
                                id = it.id,
                                description = it.description
                            )
                    }
                }
                is ResultWrapper.NetworkError -> {
                    Log.d("Psynas", "net error")
                    return emptyList()
                }
                is ResultWrapper.GenericError -> {
                    Log.d("Psynas", psynas.code.toString() + psynas.error)
                    return emptyList()
                }
           }
    }


    fun likePsyna(psynaId: Int) {
        Log.d(LOG_TAG, "Like $psynaId")
        viewModelScope.launch {
            safeApiCall(Dispatchers.IO) {
                com.psinder.myapplication.network.provideApi().like(
                    bearerToken = "Bearer ${token.value}",
                    LikeRequest(psynaId)
                )
            }
        }
    }

    fun dislikePsyna(psynaId: Int) {
        Log.d(LOG_TAG, "Dislike $psynaId")

    }

    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val userList: List<Profile>) : ViewState()
    }
}