package com.psinder.myapplication

import android.content.Intent
import android.os.Bundle
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
import com.psinder.myapplication.network.LoginData
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.provideApi
import com.psinder.myapplication.network.safeApiCall
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
            onLoginClick(binding)
        }

        return binding.root
    }



    fun onLoginClick(binding: FragmentLoginBinding) {
        binding.cirLoginButton.startAnimation()
        lifecycleScope.launch {
            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                val result = safeApiCall(Dispatchers.IO) {
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
                Toast.makeText(this@LoginFragment.context, message, Toast.LENGTH_SHORT).show()
                binding.cirLoginButton.revertAnimation()
            }
        }
    }
}