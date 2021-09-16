package com.psinder.myapplication

import android.content.Intent
import android.os.Bundle
import android.view.View
import android.widget.EditText
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import com.psinder.myapplication.network.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch


class LoginActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        //for changing status bar icon colors
        //window.decorView.systemUiVisibility = View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR
        setContentView(R.layout.activity_login)
    }

    fun onRegisterClick(View: View?) {
        startActivity(Intent(this, RegisterActivity::class.java))
        overridePendingTransition(R.anim.slide_in_right, R.anim.stay)
    }

    fun onLoginClick(view: View?) {
        (view as CircularProgressButton).startAnimation()
        lifecycleScope.launch {
            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                val result = safeApiCall(Dispatchers.IO) {
                    provideApi().login(
                        LoginData(
                            findViewById<EditText>(R.id.editTextEmail).text.toString(),
                            findViewById<EditText>(R.id.editTextPassword).text.toString()
                        )
                    )
                }
                val message = when(result) {
                    is ResultWrapper.Success -> "token: " + result.value.token
                    is ResultWrapper.NetworkError -> "network error"
                    is ResultWrapper.GenericError -> "code:" + result.code // TODO: why error message is null???
                }
                Toast.makeText(this@LoginActivity, message, Toast.LENGTH_SHORT).show()
                    view.revertAnimation()
            }
        }
    }
}