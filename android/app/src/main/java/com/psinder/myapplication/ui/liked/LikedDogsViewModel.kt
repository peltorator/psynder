package com.psinder.myapplication.ui.liked

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.network.Psyna
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.SwipeApi
import com.psinder.myapplication.util.safeApiCall
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class LikedDogsViewModel @Inject constructor(
    private val api: SwipeApi
): ViewModel() {

    companion object {
        const val LOG_TAG = "LikedDogsViewModel"
    }

    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()


    init {
        loadLiked()
    }

    private fun loadLiked() {
        viewModelScope.launch {
            _viewState.emit(ViewState.Loading)
            val response = safeApiCall(Dispatchers.IO) {
                api.liked()
            }

            when (response) {
                is ResultWrapper.Success -> {
                    _viewState.emit(ViewState.Data(response.value))
                }
                is ResultWrapper.NetworkError -> {
                    Log.d(LOG_TAG, "net error")
                }
                is ResultWrapper.GenericError -> {
                    Log.d(LOG_TAG, response.code.toString() + response.error)
                }
            }
        }
    }

    sealed class ViewState {
        object Loading : ViewState()
        data class Data(val psynaList: List<Psyna>) : ViewState()
    }
}