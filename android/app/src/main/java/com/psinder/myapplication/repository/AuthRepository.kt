package com.psinder.myapplication.repository

import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.entity.toAccountKind
import com.psinder.myapplication.network.LoginData
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.provideApi
import com.psinder.myapplication.network.safeApiCall
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow

data class AuthState(val isAuthorized: Boolean, val kind: AccountKind)

object AuthRepository {

    private val _authStateFlow = MutableStateFlow(AuthState(false, AccountKind.UNDEFINED))
    private var _token = ""

    val authStateFlow = _authStateFlow.asStateFlow()
    val token
        get() = _token

    suspend fun signIn(email: String, password: String) {
        val result = safeApiCall(Dispatchers.IO) {
            provideApi().login(LoginData(email, password))
        }
        if (result is ResultWrapper.Success) {
            _token = result.value.token
            val kind = result.value.kind.toAccountKind()
            if (kind == AccountKind.UNDEFINED) {
                throw Exception("Undefined kind!")
            }
            _authStateFlow.emit(AuthState(true, kind))
        } else {
            val message = when (result) {
                is ResultWrapper.GenericError -> result.error.toString()
                else -> "network error"
            }
            throw Exception(message)
        }
    }

    suspend fun logout() {
        _authStateFlow.emit(AuthState(false, AccountKind.UNDEFINED))
    }
}
