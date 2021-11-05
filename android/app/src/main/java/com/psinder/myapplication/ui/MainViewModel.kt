package com.psinder.myapplication.ui

import androidx.lifecycle.ViewModel
import com.psinder.myapplication.repository.AuthRepository
import com.psinder.myapplication.repository.AuthState
import kotlinx.coroutines.flow.Flow

class MainViewModel : ViewModel() {
    val authStateFlow: Flow<AuthState> = AuthRepository.authStateFlow
}