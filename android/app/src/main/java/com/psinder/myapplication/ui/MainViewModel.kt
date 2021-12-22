package com.psinder.myapplication.ui

import androidx.lifecycle.ViewModel
import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.interactor.AuthInteractor
import dagger.hilt.android.lifecycle.HiltViewModel
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject

@HiltViewModel
class MainViewModel  @Inject constructor(
    private val authInteractor: AuthInteractor
): ViewModel() {

    suspend fun accountKindFlow(): Flow<AccountKind> = authInteractor.accountKindFlow()
}