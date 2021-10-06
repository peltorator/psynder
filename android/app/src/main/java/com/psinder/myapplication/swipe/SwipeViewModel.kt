package com.psinder.myapplication.swipe

import android.util.Log
import android.widget.Toast
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.MainActivity
import com.psinder.myapplication.network.LoadPsynasRequest
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.safeApiCall
import com.squareup.moshi.Moshi
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory

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
                    bearerToken = "Bearer $token",
                    psynasData = LoadPsynasRequest(count = 100)
                )
            }

           when (psynas) {
                is ResultWrapper.Success -> {
                    return psynas.value.psynas.map {
                            Profile(
                                age = 5,
                                distance = 100,
                                name = it.name,
                                profile_pic = it.photoLink,
                                id = it.id
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


    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val userList: List<Profile>) : ViewState()
    }
}