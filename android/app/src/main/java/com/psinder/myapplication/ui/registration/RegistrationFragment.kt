package com.psinder.myapplication.ui.registration

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.fragment.app.Fragment
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.findNavController
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentRegistrationBinding
import com.psinder.myapplication.entity.AccountKind
import com.psinder.myapplication.network.RegisterData
import com.psinder.myapplication.network.ResultWrapper
import com.psinder.myapplication.network.provideApi
import com.psinder.myapplication.network.safeApiCall
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch

class RegistrationFragment: Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View {
        val binding = DataBindingUtil.inflate<FragmentRegistrationBinding>(inflater,
            R.layout.fragment_registration,container,false)
        binding.navigateButton.setOnClickListener {
            it.findNavController().navigate(R.id.action_registrationFragment_to_loginFragment)
        }
        changeStatusBarColor()

        binding.cirRegisterButton.setOnClickListener {
            onRegisterClick(binding)
        }

        return binding.root
    }
    private fun changeStatusBarColor() {
        // TODO: changeStatusBar Color for registration layout
    }

    fun onRegisterClick(binding: FragmentRegistrationBinding) {
        binding.cirRegisterButton.startAnimation()
        lifecycleScope.launch {
            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                val result = safeApiCall(Dispatchers.IO) {
                    provideApi().register(
                        RegisterData(
                            binding.editTextEmail.text.toString(),
                            binding.editTextPassword.text.toString(),
                            when(binding.userTypeButton.checkedRadioButtonId) {
                                R.id.lookingForDog -> AccountKind.PERSON
                                R.id.lookingForOwner -> AccountKind.SHELTER
                                else -> AccountKind.UNDEFINED
                            }.identifier
                        )
                    )
                }
                val message = when(result) {
                    is ResultWrapper.Success -> "id: " + result.value.id + "\ntoken: " + result.value.token
                    is ResultWrapper.NetworkError -> "network error"
                    is ResultWrapper.GenericError -> "code:" + result.code // TODO: why error message is null???
                }
//                Toast.makeText(this@RegistrationFragment.context, message, Toast.LENGTH_SHORT).show()
                binding.cirRegisterButton.revertAnimation()
            }
        }
    }
}