package com.psinder.myapplication.ui.registration

import android.util.Log
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.network.AuthorizationApi
import com.psinder.myapplication.network.RegisterData
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.util.safeApiCall
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class RegistrationViewModel @Inject constructor(
    private val api: AuthorizationApi
) : ViewModel() {
    fun signUp(email: String, password: String, accountKind: AccountKind) {
        viewModelScope.launch {
            val result = safeApiCall(Dispatchers.IO) {
                api.register(RegisterData(email, password, accountKind.identifier))
            }
            val message = when (result) {
                is ResultWrapper.Success -> "Вы успешно зарегистрировались!"
                is ResultWrapper.NetworkError -> "network error"
                is ResultWrapper.GenericError ->  result.error.toString()
            }
            Log.d("Registration",  message)
        }
    }
}