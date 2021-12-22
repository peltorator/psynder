package com.psinder.myapplication.ui.login

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.interactor.AuthInteractor
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.repository.AuthRepository
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.CoroutineExceptionHandler
import kotlinx.coroutines.launch
import javax.inject.Inject

@HiltViewModel
class LoginViewModel @Inject constructor(
    private val authInteractor: AuthInteractor
) : ViewModel() {
    fun signIn(
        email: String,
        password: String,
        coroutineExceptionHandler: CoroutineExceptionHandler
    ) {
        viewModelScope.launch(coroutineExceptionHandler) {
            authInteractor.signIn(email, password)
        }
    }
}

