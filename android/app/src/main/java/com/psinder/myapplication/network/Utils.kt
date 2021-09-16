package com.psinder.myapplication.network

import com.squareup.moshi.Moshi
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory

private fun provideOkHttpClient(): OkHttpClient {
    return OkHttpClient.Builder().build()
}

private fun provideMoshi(): Moshi {
    return Moshi.Builder().build()
}

fun provideApi(): Api {
    return Retrofit.Builder()
        .client(provideOkHttpClient())
        .baseUrl("https://reqres.in/api/")
        .addConverterFactory(MoshiConverterFactory.create(provideMoshi()))
        .build()
        .create(Api::class.java)
}