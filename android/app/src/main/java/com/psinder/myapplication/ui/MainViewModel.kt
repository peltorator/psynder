package com.psinder.myapplication.ui

import androidx.lifecycle.ViewModel
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.flow.Flow

class MainViewModel : ViewModel() {
    val isAuthorizedFlow: Flow<Boolean> = AuthRepository.isAuthorizedFlow
}