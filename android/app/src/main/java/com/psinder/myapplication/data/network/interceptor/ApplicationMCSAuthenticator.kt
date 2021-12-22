package com.psinder.myapplication.data.network.interceptor

import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.runBlocking
import okhttp3.Authenticator
import okhttp3.Request
import okhttp3.Response
import okhttp3.Route

class ApplicationMCSAuthenticator(
    private val authRepository: AuthRepository
) : Authenticator {

    override fun authenticate(route: Route?, response: Response): Request? {
        if (1 < responseCount(response)) {
            return null
        }
        val token = runBlocking {
            authRepository.getAuthTokenFlow().first()?.token ?: return@runBlocking ""
        }
        return response
            .request
            .newBuilder()
            .header("Authorization", "Bearer $token")
            .build()
    }

    private fun responseCount(response: Response): Int {
        var resp = response
        var result = 1
        while (resp.priorResponse?.also { resp = it } != null) {
            result++
        }
        return result
    }
}