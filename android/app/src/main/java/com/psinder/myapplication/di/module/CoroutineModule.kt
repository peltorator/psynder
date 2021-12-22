package com.psinder.myapplication.di.module

import com.psinder.myapplication.di.*
import dagger.Module
import dagger.Provides
import dagger.hilt.InstallIn
import dagger.hilt.components.SingletonComponent
import kotlinx.coroutines.CoroutineDispatcher
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import javax.inject.Singleton

@Module
@InstallIn(SingletonComponent::class)
object CoroutineModule {

    @DefaultCoroutineDispatcher
    @Provides
    fun providesDefaultCoroutineDispatcher(): CoroutineDispatcher = Dispatchers.Default

    @IoCoroutineDispatcher
    @Provides
    fun providesIoCoroutineDispatcher(): CoroutineDispatcher = Dispatchers.IO

    @MainCoroutineDispatcher
    @Provides
    fun providesMainCoroutineDispatcher(): CoroutineDispatcher = Dispatchers.Main

    @MainImmediateCoroutineDispatcher
    @Provides
    fun providesMainImmediateCoroutineDispatcher(): CoroutineDispatcher = Dispatchers.Main.immediate

    @Provides
    @Singleton
    @AppCoroutineScope
    fun provideAppCoroutineScope(@DefaultCoroutineDispatcher defaultDispatcher: CoroutineDispatcher): CoroutineScope =
        CoroutineScope(SupervisorJob() + defaultDispatcher)
}