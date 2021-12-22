package com.psinder.myapplication.ui.registration

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.navigation.findNavController
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentRegistrationBinding
import com.psinder.myapplication.entity.AccountKind
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class RegistrationFragment : Fragment(R.layout.fragment_registration) {
    private val viewBinding by viewBinding(FragmentRegistrationBinding::bind)
    private val viewModel: RegistrationViewModel by viewModels()

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        viewBinding.navigateButton.setOnClickListener {
            it.findNavController().navigate(R.id.action_registrationFragment_to_loginFragment)
        }

        viewBinding.cirRegisterButton.setOnClickListener {
            onRegisterClick(viewBinding)
        }
    }

    private fun onRegisterClick(binding: FragmentRegistrationBinding) {
        binding.cirRegisterButton.startAnimation()
        viewModel.signUp(
            binding.editTextEmail.text.toString(),
            binding.editTextPassword.text.toString(),
            when (binding.userTypeButton.checkedRadioButtonId) {
                R.id.lookingForDog -> AccountKind.PERSON
                R.id.lookingForOwner -> AccountKind.SHELTER
                else -> AccountKind.UNDEFINED
            }
        )

        binding.cirRegisterButton.revertAnimation()
    }
}