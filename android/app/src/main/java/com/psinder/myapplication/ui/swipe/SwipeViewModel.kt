package com.psinder.myapplication.ui.swipe

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.entity.Profile
import com.psinder.myapplication.network.LikeRequest
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.SwipeApi
import com.psinder.myapplication.repository.OffsetsRepository
import com.psinder.myapplication.util.safeApiCall
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class SwipeViewModel @Inject constructor(
    private val api: SwipeApi,
    private val offsetsRepository: OffsetsRepository
) : ViewModel() {


    companion object {
        val LOG_TAG = "SwipeViewModel"
    }

    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()

    val token: MutableStateFlow<String> = MutableStateFlow("")

    init {
        viewModelScope.launch {
            val users = loadPsynas()
            _viewState.emit(ViewState.Data(users))
        }
    }


    private suspend fun loadPsynas(): List<Profile> {
        val offset = offsetsRepository.getSwipeOffsetFlow().first()
        val psynas = safeApiCall(Dispatchers.IO) {
            api.loadpsynas(offset ?: 0)
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
            val result = safeApiCall(Dispatchers.IO) {
                api.like(LikeRequest(psynaId))
            }
            if (result is ResultWrapper.Success) {
                offsetsRepository.incrementSwipeOffset()
            }
        }
    }

    fun dislikePsyna(psynaId: Int) {
        Log.d(LOG_TAG, "Dislike $psynaId")
        viewModelScope.launch {
            offsetsRepository.incrementSwipeOffset()
        }
    }

    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val userList: List<Profile>) : ViewState()
    }
}