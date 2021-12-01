package com.psinder.myapplication.util

import android.content.res.Resources
import android.text.SpannableStringBuilder
import android.text.Spanned
import android.widget.TextView
import androidx.annotation.StringRes
import androidx.core.text.buildSpannedString
import androidx.core.text.inSpans

fun Resources.getSpannedString(@StringRes resId: Int, vararg formatArgs: CharSequence): Spanned {
    var lastArgIndex = 0
    val spannableStringBuilder = SpannableStringBuilder(getString(resId, *formatArgs))
    for (arg in formatArgs) {
        val argString = arg.toString()
        lastArgIndex = spannableStringBuilder.indexOf(argString, lastArgIndex)
        if (lastArgIndex != -1) {
            spannableStringBuilder.replace(lastArgIndex, lastArgIndex + argString.length, arg)
            lastArgIndex += argString.length
        }
    }
    return spannableStringBuilder
}

fun TextView.setAmount(template: Int, amount: Int) {
    text = resources.getSpannedString(
        template,
        buildSpannedString {
            inSpans {
                append(amount.toString())
            }
        }
    )
}
