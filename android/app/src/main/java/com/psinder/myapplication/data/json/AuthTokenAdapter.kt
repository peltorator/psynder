package com.psinder.myapplication.data.json

import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.entity.AuthToken
import com.psinder.myapplication.entity.toAccountKind
import com.squareup.moshi.JsonAdapter
import com.squareup.moshi.JsonDataException
import com.squareup.moshi.JsonReader
import com.squareup.moshi.JsonWriter

class AuthTokenAdapter : JsonAdapter<AuthToken>() {

    companion object {

        private const val FIELD_NAME_TOKEN = "token"
        private const val FIELD_NAME_KIND = "kind"

        private val FIELD_NAMES = JsonReader.Options.of(
            FIELD_NAME_TOKEN,  // 0
            FIELD_NAME_KIND  // 1
        )
    }

    override fun fromJson(reader: JsonReader): AuthToken {

        var token: String? = null
        var kind: AccountKind?  = null

        reader.beginObject()
        while (reader.hasNext()) {
            when(reader.selectName(FIELD_NAMES)) {
                0 -> {
                    token = reader.nextString()
                }
                1 -> {
                    kind = reader.nextString().toAccountKind()
                }
                else -> {
                    reader.skipName()
                    reader.skipValue()
                }
            }
        }
        reader.endObject()

        return AuthToken(
            token ?: throw JsonDataException("Required property '$FIELD_NAME_TOKEN' missing at ${reader.path}"),
            kind ?: throw JsonDataException("Required property '$FIELD_NAME_KIND' missing at ${reader.path}")
        )
    }

    override fun toJson(writer: JsonWriter, authTokens: AuthToken?) {
        authTokens?.let {
            writer.beginObject()
            writer.name(FIELD_NAME_TOKEN)
            writer.value(it.token)
            writer.name(FIELD_NAME_KIND)
            writer.value(it.kind.identifier)
            writer.endObject()
        }
    }
}