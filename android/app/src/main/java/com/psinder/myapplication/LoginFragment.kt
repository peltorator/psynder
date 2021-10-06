package com.psinder.myapplication

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.EditText
import android.widget.Toast
import androidx.databinding.DataBindingUtil
import androidx.fragment.app.Fragment
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.findNavController
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import com.psinder.myapplication.databinding.FragmentLoginBinding
import com.psinder.myapplication.network.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

// TODO: Hide keyboard at login form (after clocking login?)
class LoginFragment: Fragment() {
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val binding = DataBindingUtil.inflate<FragmentLoginBinding>(inflater,
            R.layout.fragment_login,container,false)
        binding.navigateButton.setOnClickListener {
            it.findNavController().navigate(R.id.action_loginFragment_to_registrationFragment)
        }

        binding.navRegister.setOnClickListener {
            it.findNavController().navigate(R.id.action_loginFragment_to_registrationFragment)
        }

        binding.cirLoginButton.setOnClickListener {
            (it as CircularProgressButton).startAnimation()
            lifecycleScope.launch {
                lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                    val result = safeApiCall(Dispatchers.IO) {
                        Log.d("Login", binding.editTextEmail.text.toString())
                        Log.d("Login", binding.editTextPassword.text.toString())
                        provideApi().login(
                            LoginData(
                                binding.editTextEmail.text.toString(),
                                binding.editTextPassword.text.toString()
                            )
                        )
                    }
                    val message = when(result) {
                        is ResultWrapper.Success -> "token: " + result.value.token
                        is ResultWrapper.NetworkError -> "network error"
                        is ResultWrapper.GenericError -> "code:" + result.code // TODO: why error message is null???
                    }
//                    Toast.makeText(this@LoginFragment.context, message, Toast.LENGTH_SHORT).show()

//                    if (result is ResultWrapper.Success) {
//                        val psynas = safeApiCall(Dispatchers.IO) {
//                            provideApi().loadpsynas(
//                                bearerToken = "Bearer ${result.value.token}",
//                                psynasData = LoadPsynasRequest(count=10)
//                            )
//                        }
//
//                        val messagePsynas = when(psynas) {
//                            is ResultWrapper.Success -> "psynas: " + psynas.value.psynas
//                            is ResultWrapper.NetworkError -> "network error"
//                            is ResultWrapper.GenericError -> "code:" + psynas.error // TODO: why error message is null???
//                        }
//                        Toast.makeText(this@LoginFragment.context, messagePsynas, Toast.LENGTH_SHORT).show()
//
//
//                    }

                    // TODO: Please fix this Ivan Pavlov 30.09
                    if (result is ResultWrapper.Success) {
                        Log.d("Activity", "Start another activity")
                        val intent = Intent(context, MainActivity::class.java)
                        intent.putExtra("TOKEN", result.value.token)
                        startActivity(intent)
//                        it.findNavController().navigate(R.id.action_loginFragment_to_profileFragment2)
                    }

                    (it as CircularProgressButton).revertAnimation()
                }
            }
        }

        return binding.root
    }

}