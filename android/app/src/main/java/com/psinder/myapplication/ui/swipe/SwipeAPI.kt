package com.psinder.myapplication.ui.swipe

import com.psinder.myapplication.entity.Profile
import retrofit2.Call
import retrofit2.Retrofit
import retrofit2.converter.gson.GsonConverterFactory
import retrofit2.http.GET

interface SwipeAPI {

    @GET("profiles")
    fun getProfiles(): Call<List<Profile>>

    companion object {
        operator fun invoke(): SwipeAPI {
            return Retrofit.Builder()
                .baseUrl("https://api.simplifiedcoding.in/course-apis/tinder/")
                .addConverterFactory(GsonConverterFactory.create())
                .build()
                .create(SwipeAPI::class.java)
        }
    }
}