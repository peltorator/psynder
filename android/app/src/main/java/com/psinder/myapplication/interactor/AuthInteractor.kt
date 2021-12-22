package com.psinder.myapplication.interactor

import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.entity.AuthToken
import com.psinder.myapplication.entity.toAccountKind
import com.psinder.myapplication.network.LoginResponse
import com.psinder.myapplication.network.RegistrationResponse
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.flow.Flow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class AuthInteractor @Inject constructor(
    private val authRepository: AuthRepository
) {

    suspend fun accountKindFlow(): Flow<AccountKind> = authRepository.accountKindFlow()

    suspend fun signIn(email: String, password: String): ResultWrapper<LoginResponse> {
        val result = authRepository.generateAuthToken(email, password)
        if (result is ResultWrapper.Success) {
            authRepository.saveAuthToken(AuthToken(result.value.token, result.value.kind.toAccountKind()))
        } else {
            val message = when (result) {
                is ResultWrapper.GenericError -> result.error.toString()
                else -> "network error"
            }
            throw Exception(message)
        }
        return result
    }

    suspend fun logout() {
        authRepository.saveAuthToken(null)
    }
}