package com.psinder.myapplication.ui.doglist

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.network.*
import com.psinder.myapplication.util.safeApiCall
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class DogListViewModel @Inject constructor(
    private val api: ShelterApi
) : ViewModel() {

    companion object {
        val LOG_TAG = "DogListViewModel"
    }

    private val _viewState = MutableStateFlow<ViewState>(ViewState.Loading)
    val viewState: Flow<ViewState> get() = _viewState.asStateFlow()


    init {
        viewModelScope.launch {
            _viewState.emit(ViewState.Loading)
            Log.d(LOG_TAG, "Start loading users")
            val psynas = loadPsynas()
            Log.d(LOG_TAG, "End loading users")
            _viewState.emit(ViewState.Data(psynas))
        }
    }

    private fun mockPsynas(): List<Psyna> {
        return listOf(
            Psyna(
                1,
                "Биба",
                "",
                "Описание1",
                "https://sun9-10.userapi.com/c830408/v830408596/1e3417/lWKS4Fju4T0.jpg"
            ),
            Psyna(
                2,
                "Боба",
                "",
                "Описание2",
                "https://www.meme-arsenal.com/memes/c1b8a99053c58dbb02aec00361bb2db1.jpg"
            ),
            Psyna(
                3,
                "Иван",
                "",
                "Описание2",
                "https://thypix.com/wp-content/uploads/lustige-tiere-24.jpg"
            ),
            Psyna(
                4, "Кобан",
                "",
                "Описание2",
                "https://funik.ru/wp-content/uploads/2018/11/9b2d50675bd5ad956231-700x525.jpg"
            ),
            Psyna(
                5, "Буба",
                "",
                "Описание2",
                "https://www.fresher.ru/manager_content/images/sobaki-kotorye-prosto-ne-mogut/big/4.jpg"
            ),
            Psyna(
                6, "Добби",
                "",
                "Описание2",
                "https://i.pinimg.com/236x/cf/77/53/cf7753e2bb8398d25868b23975908bf8.jpg"
            )
        )
    }

    private suspend fun loadPsynas(): List<Psyna> {
        val psynas = safeApiCall(Dispatchers.IO) {
            api.loadpsynas()
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