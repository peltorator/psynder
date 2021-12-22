package com.psinder.myapplication.ui.editdog

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.network.AddPsynaRequest
import com.psinder.myapplication.network.Psyna
import com.psinder.myapplication.network.ShelterApi
import com.psinder.myapplication.util.safeApiCall
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class EditDogViewModel @Inject constructor(
    private val api: ShelterApi
) : ViewModel() {

    fun addPsyna(dog: Psyna) {
        viewModelScope.launch {
            safeApiCall(Dispatchers.IO) {
                api.addPsyna(
                    AddPsynaRequest(
                        name = dog.name,
                        breed = dog.breed ?: "loma",
                        description = dog.description,
                        photoLink = dog.photoLink
                    )
                )
            }
        }
    }
}