package com.psinder.myapplication.data.persistent

import android.content.SharedPreferences
import androidx.core.content.edit
import com.psinder.myapplication.entity.AuthToken
import com.squareup.moshi.JsonAdapter
import com.squareup.moshi.JsonDataException
import com.squareup.moshi.Moshi
import kotlin.properties.ReadWriteProperty
import kotlin.reflect.KProperty

class LocalKeyValueStorage(
    private val pref: SharedPreferences,
    moshi: Moshi
) {

    var authToken: AuthToken? by JsonDelegate(
        AUTH_TOKEN,
        pref,
        moshi.adapter(AuthToken::class.java)
    )

    var swipeOffset: Int? by IntDelegate(
        SWIPE_OFFSET,
        pref
    )

    private class JsonDelegate<T>(
        private val key: String,
        private val pref: SharedPreferences,
        private val adapter: JsonAdapter<T>
    ) : ReadWriteProperty<LocalKeyValueStorage, T?> {

        override fun setValue(thisRef: LocalKeyValueStorage, property: KProperty<*>, value: T?) {
            pref.edit(commit = true) {
                if (value == null) remove(key)
                else putString(key, adapter.toJson(value))
            }
        }

        override fun getValue(thisRef: LocalKeyValueStorage, property: KProperty<*>): T? {
            return try {
                pref.getString(key, null)?.let { adapter.fromJson(it) }
            } catch (e: JsonDataException) {
                setValue(thisRef, property, null)
                null
            }
        }
    }

    private class IntDelegate(
        private val key: String,
        private val pref: SharedPreferences
    ) : ReadWriteProperty<LocalKeyValueStorage, Int?> {

        override fun setValue(thisRef: LocalKeyValueStorage, property: KProperty<*>, value: Int?) {
            pref.edit(commit = true) {
                if (value == null) remove(key)
                else putString(key, value.toString())
            }
        }

        override fun getValue(thisRef: LocalKeyValueStorage, property: KProperty<*>): Int? {
            return try {
                pref.getString(key, null)?.toInt()
            } catch (e: JsonDataException) {
                setValue(thisRef, property, null)
                null
            }
        }
    }

    companion object {
        private const val AUTH_TOKEN = "auth_token"
        private const val SWIPE_OFFSET = "swipe_offset"
    }
}