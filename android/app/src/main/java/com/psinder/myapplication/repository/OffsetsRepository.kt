package com.psinder.myapplication.repository

import com.psinder.myapplication.data.persistent.LocalKeyValueStorage
import com.psinder.myapplication.di.AppCoroutineScope
import com.psinder.myapplication.di.IoCoroutineDispatcher
import com.psinder.myapplication.entity.AuthToken
import kotlinx.coroutines.*
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import javax.inject.Inject
import javax.inject.Singleton

@Singleton
class OffsetsRepository @Inject constructor(
    private val localKeyValueStorage: LocalKeyValueStorage,
    @AppCoroutineScope externalCoroutineScope: CoroutineScope,
    @IoCoroutineDispatcher private val ioDispatcher: CoroutineDispatcher
) {
    private val swipeOffsetFlow: Deferred<MutableStateFlow<Int?>> =
        externalCoroutineScope.async(context = ioDispatcher, start = CoroutineStart.LAZY) {
            MutableStateFlow(localKeyValueStorage.swipeOffset)
        }

    suspend fun getSwipeOffsetFlow(): StateFlow<Int?> {
        return swipeOffsetFlow.await().asStateFlow()
    }

    suspend fun saveSwipeOffset(swipeOffset: Int?) {
        withContext(ioDispatcher) {
            localKeyValueStorage.swipeOffset = swipeOffset
        }
        swipeOffsetFlow.await().emit(swipeOffset)
    }

    suspend fun incrementSwipeOffset() {
        val offset = localKeyValueStorage.swipeOffset?.plus(1)
        withContext(ioDispatcher) {
            localKeyValueStorage.swipeOffset = offset
        }
        swipeOffsetFlow.await().emit(offset)
    }
}

