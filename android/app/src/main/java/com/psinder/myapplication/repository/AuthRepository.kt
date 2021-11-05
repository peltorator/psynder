package com.psinder.myapplication.repository

import com.psinder.myapplication.network.LoginData
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.provideApi
import com.psinder.myapplication.network.safeApiCall
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow


object AuthRepository {

    private val _isAuthorizedFlow = MutableStateFlow(false)
    private var _token = ""

    val isAuthorizedFlow = _isAuthorizedFlow.asStateFlow()
    val token
        get() = _token

    suspend fun signIn(email: String, password: String) {
        val result = safeApiCall(Dispatchers.IO) {
            provideApi().login(LoginData(email, password))
        }
        if (result is ResultWrapper.Success) {
            _token = result.value.token
            _isAuthorizedFlow.emit(true)
        } else {
            val message = when (result) {
                is ResultWrapper.GenericError -> result.error.toString()
                else -> "network error"
            }
            throw Exception(message)
        }
    }

    suspend fun logout() {
        _isAuthorizedFlow.emit(false)
    }
}
