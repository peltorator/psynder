package com.psinder.myapplication.ui.profile

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.CoroutineExceptionHandler
import kotlinx.coroutines.launch

class UserProfileViewModel : ViewModel() {
    fun signOut(coroutineExceptionHandler: CoroutineExceptionHandler) {
        viewModelScope.launch(coroutineExceptionHandler) {
            AuthRepository.logout()
        }
    }
}