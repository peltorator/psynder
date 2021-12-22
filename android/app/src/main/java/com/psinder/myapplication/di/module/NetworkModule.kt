package com.psinder.myapplication.di.module

import com.psinder.myapplication.data.network.interceptor.ApplicationMCSAuthenticator
import com.psinder.myapplication.data.network.interceptor.AuthorizationInterceptor
import com.psinder.myapplication.network.AuthorizationApi
import com.psinder.myapplication.network.ShelterApi
import com.psinder.myapplication.network.SwipeApi
import com.psinder.myapplication.repository.AuthRepository
import com.squareup.moshi.Moshi
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import okhttp3.OkHttpClient
import okhttp3.logging.HttpLoggingInterceptor
import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory
import java.security.SecureRandom
import java.security.cert.X509Certificate
import javax.inject.Singleton
import javax.net.ssl.SSLContext
import javax.net.ssl.X509TrustManager

@Module
@InstallIn(SingletonComponent::class)
object NetworkModule {

    @Provides
    @Singleton
    fun provideOkhttpClient(authRepository: AuthRepository): OkHttpClient {
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
            .apply {
                sslSocketFactory(sslContext.socketFactory, trustAllCerts)
                hostnameVerifier { hostname, session -> true }
                addNetworkInterceptor(AuthorizationInterceptor(authRepository))
                authenticator(ApplicationMCSAuthenticator(authRepository))
                val logging = HttpLoggingInterceptor()
                logging.setLevel(HttpLoggingInterceptor.Level.BODY)
                addInterceptor(logging)
            }
            .build()
    }

    @Provides
    @Singleton
    fun provideAuthorizationApi(okHttpClient: OkHttpClient, moshi: Moshi): AuthorizationApi =
        Retrofit.Builder()
            .client(okHttpClient)
            .baseUrl("https://10.0.2.2:444/")
            .addConverterFactory(MoshiConverterFactory.create(moshi))
            .build()
            .create(AuthorizationApi::class.java)

    @Provides
    @Singleton
    fun provideSwipeApi(okHttpClient: OkHttpClient, moshi: Moshi): SwipeApi =
        Retrofit.Builder()
            .client(okHttpClient)
            .baseUrl("https://10.0.2.2:445/")
            .addConverterFactory(MoshiConverterFactory.create(moshi))
            .build()
            .create(SwipeApi::class.java)

    @Provides
    @Singleton
    fun provideShelterApi(okHttpClient: OkHttpClient, moshi: Moshi): ShelterApi =
        Retrofit.Builder()
            .client(okHttpClient)
            .baseUrl("https://10.0.2.2:443/")
            .addConverterFactory(MoshiConverterFactory.create(moshi))
            .build()
            .create(ShelterApi::class.java)
}