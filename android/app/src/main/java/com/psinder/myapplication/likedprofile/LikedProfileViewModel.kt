package com.psinder.myapplication.likedprofile

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.network.*
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch


class LikedProfileViewModel : ViewModel() {

    companion object {
        val LOG_TAG = "DogListViewModel"
    }

    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()
    val token: MutableStateFlow<String> = MutableStateFlow("")
    var psynaId: Int = 0


    init {
        viewModelScope.launch {
            _viewState.emit(ViewState.Loading)

            val shelterInfo = getInfo(AuthRepository.token) //loadPsynas()
            Log.d(LOG_TAG, "End loading users")
            _viewState.emit(ViewState.Data(shelterInfo))
        }
    }

    private fun mockInfo(token: String): Shelter {
        return Shelter(0, "nizhny novgorod", "da", "da")
    }

    private suspend fun getInfo(token: String): Shelter? {
        val shelter = safeApiCall(Dispatchers.IO) {
            provideApi("").getShelterInfo(
                bearerToken = "Bearer $token",
                psynasRequest = LikeRequest(psynaId = psynaId)
            )
        }



        return when (shelter) {
            is ResultWrapper.Success -> {
                Log.d("Psynas", shelter.value.city)
                shelter.value

            }
            is ResultWrapper.NetworkError -> {
                Log.d("Psynas", "net error")

                null
            }
            is ResultWrapper.GenericError -> {
                Log.d("Psynas", shelter.code.toString() + shelter.error)
                null
            }
        }
    }


    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val shelter: Shelter?) : ViewState()
    }
}