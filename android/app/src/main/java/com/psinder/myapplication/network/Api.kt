package com.psinder.myapplication.network

import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass
import retrofit2.http.Body
import retrofit2.http.GET
import retrofit2.http.POST

interface Api {
    @POST("login")
    suspend fun login(@Body loginData: LoginData): Token
//
//    @POST("register")
//    suspend fun register(): RegistrationResponse
}

@JsonClass(generateAdapter = true)
data class Token(
    @Json(name="token") val token: String
)

@JsonClass(generateAdapter = true)
data class LoginData(
    @Json(name="email")
    val email: String,
    @Json(name="password")
    val password: String
)

@JsonClass(generateAdapter = true)
data class RegistrationResponse(
    val id: String,
    val token: String
)


@JsonClass(generateAdapter = true)
data class User(
    @Json(name = "avatar") val avatarUrl: String, // For example: "https://mydomain.com/user_1_avatar.jpg"
    @Json(name = "first_name") val userName: String,
    @Json(name = "email") val groupName: String
)