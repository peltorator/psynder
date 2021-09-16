package com.psinder.myapplication

import android.content.Intent
import android.os.Build
import android.os.Bundle
import android.view.View
import android.view.WindowManager
import android.widget.EditText
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.ui.AppBarConfiguration
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import com.psinder.myapplication.databinding.ActivityLoginBinding
import com.psinder.myapplication.network.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch


class RegisterActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_register)
        changeStatusBarColor()
    }

    private fun changeStatusBarColor() {
        val window = window
        window.addFlags(WindowManager.LayoutParams.FLAG_DRAWS_SYSTEM_BAR_BACKGROUNDS)
        //            window.setStatusBarColor(Color.TRANSPARENT);
        window.statusBarColor = resources.getColor(R.color.register_bk_color)
    }

    fun onLoginClick(view: View?) {
        startActivity(Intent(this, LoginActivity::class.java))
        overridePendingTransition(R.anim.slide_in_left, android.R.anim.slide_out_right)
    }

    fun onRegisterClick(view: View?) {
        (view as CircularProgressButton).startAnimation()
        lifecycleScope.launch {
            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                val result = safeApiCall(Dispatchers.IO) {
                    provideApi().register(
                        RegisterData(
                            findViewById<EditText>(R.id.editTextEmail).text.toString(),
                            findViewById<EditText>(R.id.editTextPassword).text.toString()
                        )
                    )
                }
                val message = when(result) {
                    is ResultWrapper.Success -> "id: " + result.value.id + "\ntoken: " + result.value.token
                    is ResultWrapper.NetworkError -> "network error"
                    is ResultWrapper.GenericError -> "code:" + result.code // TODO: why error message is null???
                }
                Toast.makeText(this@RegisterActivity, message, Toast.LENGTH_SHORT).show()
                view.revertAnimation()
            }
        }
    }
}