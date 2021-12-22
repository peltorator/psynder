package com.psinder.myapplication.ui.login

import android.os.Bundle
import android.view.View
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.navigation.findNavController
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentLoginBinding
import com.psinder.myapplication.network.*
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.CoroutineExceptionHandler

// TODO: Hide keyboard at login form (after clocking login?)
@AndroidEntryPoint
class LoginFragment : Fragment(R.layout.fragment_login) {
    private val viewBinding by viewBinding(FragmentLoginBinding::bind)
    private val viewModel: LoginViewModel by viewModels()

    private val coroutineExceptionHanlder = CoroutineExceptionHandler { _, throwable ->
        Toast.makeText(this@LoginFragment.context, throwable.message, Toast.LENGTH_SHORT).show()
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewBinding.navigateButton.setOnClickListener {
            it.findNavController().navigate(R.id.action_loginFragment_to_registrationFragment)
        }

        viewBinding.navRegister.setOnClickListener {
            it.findNavController().navigate(R.id.action_loginFragment_to_registrationFragment)
        }

        viewBinding.cirLoginButton.setOnClickListener {
            (it as CircularProgressButton).startAnimation()
            viewModel.signIn(
                viewBinding.editTextEmail.text?.toString() ?: "",
                viewBinding.editTextPassword.text?.toString() ?: "",
                coroutineExceptionHanlder
            )
            (it as CircularProgressButton).revertAnimation()
        }
    }
}