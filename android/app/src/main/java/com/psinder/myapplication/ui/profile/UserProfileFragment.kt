package com.psinder.myapplication.ui.profile

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.Toast
import androidx.databinding.DataBindingUtil
import androidx.fragment.app.viewModels
import androidx.navigation.findNavController
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentLoginBinding
import com.psinder.myapplication.databinding.FragmentUserProfileBinding
import com.psinder.myapplication.ui.login.LoginViewModel
import kotlinx.coroutines.CoroutineExceptionHandler

class UserProfileFragment : Fragment(R.layout.fragment_user_profile) {

    private val viewBinding by viewBinding(FragmentUserProfileBinding::bind)
    private val viewModel: UserProfileViewModel by viewModels()

    private val coroutineExceptionHanlder = CoroutineExceptionHandler { _, throwable ->
        Toast.makeText(this@UserProfileFragment.context, throwable.message, Toast.LENGTH_SHORT).show()
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)

        viewBinding.cirLogoutButton.setOnClickListener {
            (it as CircularProgressButton).startAnimation()
            viewModel.signOut(coroutineExceptionHanlder)
            (it as CircularProgressButton).revertAnimation()
        }
    }
}