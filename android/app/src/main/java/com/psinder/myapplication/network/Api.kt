package com.psinder.myapplication.network

import com.psinder.myapplication.entity.AccountKind
import com.squareup.moshi.Json
import com.squareup.moshi.JsonClass
import retrofit2.http.*

interface Api {
    @POST("login")
    suspend fun login(@Body loginData: LoginData): LoginResponse

    @POST("signup")
    suspend fun register(@Body registerData: RegisterData)

    @POST("get-all-info")
    suspend fun getInfo(@Header("Authorization") bearerToken: String): InfoResponse

    @GET("browse-psynas?limit=50&offset=0")
    suspend fun loadpsynas(
        @Header("Authorization") bearerToken: String
//                           @Query("limit") limit: String,
//                           @Query("offset") offset: String,
    ): List<Psyna>

    @POST("like-psyna")
    suspend fun like(@Header("Authorization") bearerToken: String, @Body likeData: LikeRequest)

    @GET("liked-psynas?limit=100&offset=0")
    suspend fun liked(@Header("Authorization") bearerToken: String): List<Psyna>

    @POST("browse-my-psynas?limit=100&offset=0")
    suspend fun browseShleterPsynas(@Header("Authorization") bearerToken: String): List<Psyna>

    @POST("psyna-info")
    suspend fun getShelterInfo(@Header("Authorization") bearerToken: String,
    @Body psynasRequest: LikeRequest): Shelter
}

//@JsonClass(generateAdapter = true)
//data class LoadPsynasResponse(
//    @Json(name = "psynas") val psynas: List<Psyna>
//)

@JsonClass(generateAdapter = true)
data class LikeRequest(
    @Json(name = "psynaId") val psynaId: Int
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
    @Json(name = "breed") val breed: String,
    @Json(name = "description") val description: String,
    @Json(name = "photoLink") val photoLink: String
)

@JsonClass(generateAdapter = true)
data class LoginResponse(
    @Json(name = "token") val token: String,
    @Json(name = "kind") val kind: String
)

@JsonClass(generateAdapter = true)
data class InfoResponse(
    @Json(name = "users") val users: Int,
    @Json(name = "shelters") val shelters: Int,
    @Json(name = "psynas") val dogs: Int
)

@JsonClass(generateAdapter = true)
data class ErrorResponse(
    @Json(name = "errorDisplayText") val errorDisplayText: String,
    @Json(name = "errorDescription") val errorDescription: String,
    @Json(name = "errorDebugInfo") val errorDebugInfo: String
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
    @Json(name = "kind") val kind: String,
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

@JsonClass(generateAdapter = true)
data class Shelter(
    @Json(name = "account_id") val accountId: Int,
    @Json(name = "city") val city: String,
    @Json(name = "address") val address: String,
    @Json(name = "phone") val phone: String,
)


sealed class ResultWrapper<out T> {
    data class Success<out T>(val value: T) : ResultWrapper<T>()
    data class GenericError(val code: Int? = null, val error: ErrorResponse? = null) :
        ResultWrapper<Nothing>()

    object NetworkError : ResultWrapper<Nothing>()
}