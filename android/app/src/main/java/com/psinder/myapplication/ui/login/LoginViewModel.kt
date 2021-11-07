package com.psinder.myapplication.ui.login

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.CoroutineExceptionHandler
import kotlinx.coroutines.launch


class LoginViewModel : ViewModel() {
    fun signIn(email: String, password: String, coroutineExceptionHandler: CoroutineExceptionHandler) {
        viewModelScope.launch(coroutineExceptionHandler) {
            AuthRepository.signIn(email, password)
        }
    }
}

