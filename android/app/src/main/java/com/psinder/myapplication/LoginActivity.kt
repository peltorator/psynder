package com.psinder.myapplication

import android.content.Intent
import android.os.Build
import android.os.Bundle
import android.view.View
import com.google.android.material.snackbar.Snackbar
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.navigateUp
import androidx.navigation.ui.setupActionBarWithNavController
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import com.psinder.myapplication.databinding.ActivityLoginBinding
import android.os.AsyncTask
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import android.os.Handler
import android.widget.EditText
import android.widget.Toast
import com.psinder.myapplication.network.LoginData
import com.psinder.myapplication.network.provideApi


class LoginActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        //for changing status bar icon colors
        //window.decorView.systemUiVisibility = View.SYSTEM_UI_FLAG_LIGHT_STATUS_BAR
        setContentView(R.layout.activity_login)

        val btn = findViewById<CircularProgressButton>(R.id.cirLoginButton)

//        lifecycleScope.launch {
//            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
//                withContext(Dispatchers.IO) {
//                    Thread.sleep(3000)
//                    btn.revertAnimation();
//                }
//            }
//        }


    }

    fun onRegisterClick(View: View?) {
        startActivity(Intent(this, RegisterActivity::class.java))
        overridePendingTransition(R.anim.slide_in_right, R.anim.stay)
    }

    fun onLoginClick(view: View?) {
        (view as CircularProgressButton).startAnimation()
        lifecycleScope.launch {
            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {

                val nepapka = withContext(Dispatchers.IO) {
                    //Thread.sleep(3000)
                    provideApi().login(
                    LoginData(
                        "eve.holt@reqres.in",//findViewById<EditText>(R.id.editTextEmail).text.toString(),
                        "cityslicka" // findViewById<EditText>(R.id.editTextPassword).text.toString()
                    )
                    )
                }
                println("KEK" + nepapka.token)
                view.revertAnimation()
            }
        }
    }
}