package com.psinder.myapplication.ui.home

import android.os.Bundle
import android.view.View
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentHomeBinding
import com.psinder.myapplication.ui.profile.UserProfileViewModel
import kotlinx.coroutines.CoroutineExceptionHandler

class HomeFragment: Fragment(R.layout.fragment_home) {
    private val viewBinding by viewBinding(FragmentHomeBinding::bind)
    private val viewModel: UserProfileViewModel by viewModels()
    private val coroutineExceptionHanlder = CoroutineExceptionHandler { _, throwable ->
        Toast.makeText(this@HomeFragment.context, throwable.message, Toast.LENGTH_SHORT).show()
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