package com.psinder.myapplication.network

import com.squareup.moshi.Moshi
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.withContext
import okhttp3.OkHttpClient
import retrofit2.HttpException
import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory
import java.io.IOException
import java.security.SecureRandom
import java.security.cert.X509Certificate
import javax.net.ssl.SSLContext
import javax.net.ssl.X509TrustManager

// works only with emulator, to run on local device:
// https://stackoverflow.com/questions/4779963/how-can-i-access-my-localhost-from-my-android-device
private const val BASE_URL = "https://10.0.2.2:8080/"

private fun provideOkHttpClient(): OkHttpClient {
    // TODO: this is not safe
    val trustAllCerts = object : X509TrustManager {
        override fun checkClientTrusted(p0: Array<out X509Certificate>?, p1: String?) {}

        override fun checkServerTrusted(p0: Array<out X509Certificate>?, p1: String?) {}

        override fun getAcceptedIssuers(): Array<X509Certificate> {
            return emptyArray()
        }
    }
    val sslContext: SSLContext = SSLContext.getInstance("SSL").apply {
        init(null, arrayOf(trustAllCerts), SecureRandom())
    }
    return OkHttpClient.Builder()
        .sslSocketFactory(sslContext.socketFactory, trustAllCerts)
        .hostnameVerifier { hostname, session -> true }
        .build()
}

private fun provideMoshi(): Moshi {
    return Moshi.Builder().build()
}

fun provideApi(): Api {
    return Retrofit.Builder()
        .client(provideOkHttpClient())
        .baseUrl(BASE_URL)
        .addConverterFactory(MoshiConverterFactory.create(provideMoshi()))
        .build()
        .create(Api::class.java)
}

suspend fun <T> safeApiCall(dispatcher: CoroutineDispatcher, apiCall: suspend () -> T): ResultWrapper<T> {
    return withContext(dispatcher) {
        try {
            ResultWrapper.Success(apiCall.invoke())
        } catch (throwable: Throwable) {
            when (throwable) {
                is IOException -> ResultWrapper.NetworkError
                is HttpException -> {
                    val code = throwable.code()
                    val errorResponse = convertErrorBody(throwable)
                    ResultWrapper.GenericError(code, errorResponse)
                }
                else -> {
                    ResultWrapper.GenericError(null, null)
                }
            }
        }
    }
}

private fun convertErrorBody(throwable: HttpException): ErrorResponse? {
    return try {
        throwable.response()?.errorBody()?.source()?.let {
            val moshiAdapter = Moshi.Builder().build().adapter(ErrorResponse::class.java)
            moshiAdapter.fromJson(it)
        }
    } catch (exception: Exception) {
        null
    }
}