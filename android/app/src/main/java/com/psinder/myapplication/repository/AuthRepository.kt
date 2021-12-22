package com.psinder.myapplication.repository

import com.psinder.myapplication.data.persistent.LocalKeyValueStorage
import com.psinder.myapplication.di.AppCoroutineScope
import com.psinder.myapplication.di.IoCoroutineDispatcher
import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.entity.AuthToken
import com.psinder.myapplication.network.*
import com.psinder.myapplication.util.safeApiCall
import dagger.Lazy
import kotlinx.coroutines.*
import kotlinx.coroutines.flow.*
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class AuthRepository @Inject constructor(
    private val apiLazy: Lazy<AuthorizationApi>,
    private val localKeyValueStorage: LocalKeyValueStorage,
    @AppCoroutineScope externalCoroutineScope: CoroutineScope,
    @IoCoroutineDispatcher private val ioDispatcher: CoroutineDispatcher
) {

    private val api by lazy { apiLazy.get() }

    private val authTokenFlow: Deferred<MutableStateFlow<AuthToken?>> =
        externalCoroutineScope.async(context = ioDispatcher, start = CoroutineStart.LAZY) {
            MutableStateFlow(localKeyValueStorage.authToken)
        }

    suspend fun getAuthTokenFlow(): StateFlow<AuthToken?> {
        return authTokenFlow.await().asStateFlow()
    }

    suspend fun saveAuthToken(authToken: AuthToken?) {
        withContext(ioDispatcher) {
            localKeyValueStorage.authToken = authToken
        }
        authTokenFlow.await().emit(authToken)
    }

    suspend fun accountKindFlow(): Flow<AccountKind> {
        return authTokenFlow
            .await()
            .asStateFlow()
            .map { it?.kind ?: AccountKind.UNDEFINED }
    }

    suspend fun generateAuthToken(email: String, password: String): ResultWrapper<LoginResponse> {
        return safeApiCall(Dispatchers.IO) {
            api.login(LoginData(email, password))
        }
    }
}
