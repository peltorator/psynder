package com.psinder.myapplication.di.module

import com.psinder.myapplication.data.json.AuthTokenAdapter
import com.psinder.myapplication.entity.AuthToken
import com.squareup.moshi.Moshi
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object SerializationModule {

    @Provides
    @Singleton
    fun provideMoshi(): Moshi = Moshi.Builder()
        .add(AuthToken::class.java, AuthTokenAdapter().nullSafe())
        .build()

}