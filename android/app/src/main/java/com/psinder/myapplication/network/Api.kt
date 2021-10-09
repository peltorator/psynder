package com.psinder.myapplication.network

import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass
import retrofit2.http.Body
import retrofit2.http.Header
import retrofit2.http.POST

interface Api {
    @POST("login")
    suspend fun login(@Body loginData: LoginData): Token

    @POST("signup")
    suspend fun register(@Body registerData: RegisterData): RegistrationResponse

    @POST("loadpsynas")
    suspend fun loadpsynas(@Header("Authorization") bearerToken: String,
                           @Body psynasData: LoadPsynasRequest): LoadPsynasResponse
}

@JsonClass(generateAdapter = true)
data class LoadPsynasResponse(
    @Json(name = "psynas") val psynas: List<Psyna>
)

@JsonClass(generateAdapter = true)
data class LoadPsynasRequest(
    @Json(name = "count") val count: Int
)

//type Psyna struct {
//    Id PsynaId `json:"id"`
//    Name string `json:"name"`
//    Description string `json:"description"`
//    PhotoLink string `json:"photo_link"`
//}

@JsonClass(generateAdapter = true)
data class Psyna(
    @Json(name = "id") val id: Int,
    @Json(name = "name") val name: String,
    @Json(name = "description") val description: String,
    @Json(name = "photo_link") val photoLink: String
)

@JsonClass(generateAdapter = true)
data class Token(
    @Json(name = "token") val token: String
)

data class ErrorResponse(
    @Json(name = "error") val error: String
)

@JsonClass(generateAdapter = true)
data class LoginData(
    @Json(name = "email") val email: String,
    @Json(name = "password") val password: String
)

@JsonClass(generateAdapter = true)
data class RegisterData(
    // @Json(name = "name") val name: String, TODO: add when api is ready
    @Json(name = "email") val email: String,
    @Json(name = "password") val password: String,
    // @Json(name = "mobile number") val mobileNumber: String
)

@JsonClass(generateAdapter = true)
data class RegistrationResponse(
    @Json(name = "id") val id: String,
    @Json(name = "token") val token: String
)


@JsonClass(generateAdapter = true)
data class User(
    @Json(name = "avatar") val avatarUrl: String, // For example: "https://mydomain.com/user_1_avatar.jpg"
    @Json(name = "first_name") val userName: String,
    @Json(name = "email") val groupName: String
)

sealed class ResultWrapper<out T> {
    data class Success<out T>(val value: T) : ResultWrapper<T>()
    data class GenericError(val code: Int? = null, val error: ErrorResponse? = null) :
        ResultWrapper<Nothing>()

    object NetworkError : ResultWrapper<Nothing>()
}