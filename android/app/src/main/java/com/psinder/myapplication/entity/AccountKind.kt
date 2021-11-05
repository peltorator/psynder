package com.psinder.myapplication.entity

enum class AccountKind(val identifier: String) {
    PERSON("person"),
    SHELTER("shelter"),
    UNDEFINED("undefined");
}

fun String.toAccountKind(): AccountKind =
    AccountKind.values().firstOrNull { it.identifier == this } ?: AccountKind.UNDEFINED