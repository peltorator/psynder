package com.psinder.myapplication.ui.home

import android.os.Bundle
import android.view.View
import android.widget.Toast
import androidx.core.view.isVisible
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import br.com.simplepass.loading_button_lib.customViews.CircularProgressButton
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentHomeBinding
import com.psinder.myapplication.util.setAmount
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.CoroutineExceptionHandler
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch

@AndroidEntryPoint
class HomeFragment : Fragment(R.layout.fragment_home) {
    private val viewBinding by viewBinding(FragmentHomeBinding::bind)
    private val viewModel: HomeViewModel by viewModels()
    private val coroutineExceptionHanlder = CoroutineExceptionHandler { _, throwable ->
        Toast.makeText(this@HomeFragment.context, throwable.message, Toast.LENGTH_SHORT).show()
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        subscribeToViewState()
        viewBinding.cirLogoutButton.setOnClickListener {
            (it as CircularProgressButton).startAnimation()
            viewModel.signOut(coroutineExceptionHanlder)
            (it as CircularProgressButton).revertAnimation()
        }
    }

    private fun subscribeToViewState() {
        viewLifecycleOwner.lifecycleScope.launch {
            viewLifecycleOwner.repeatOnLifecycle(Lifecycle.State.STARTED) {
                viewModel.viewState.collect { viewState -> renderViewState(viewState) }
            }
        }
    }

    private fun renderViewState(viewState: HomeViewModel.ViewState) {
        when (viewState) {
            is HomeViewModel.ViewState.Loading -> {
                viewBinding.infoAllTextViews.isVisible = false
                viewBinding.progressBar.isVisible = true
            }
            is HomeViewModel.ViewState.Data -> {
                viewBinding.infoUsersTextView.setAmount(
                    R.string.home_info_users_template,
                    viewState.info.users
                )
                viewBinding.infoSheltersTextView.setAmount(
                    R.string.home_info_shelters_template,
                    viewState.info.shelters
                )
                viewBinding.infoDogsTextView.setAmount(
                    R.string.home_info_dogs_template,
                    viewState.info.dogs
                )
                viewBinding.infoAllTextViews.isVisible = true
                viewBinding.progressBar.isVisible = false
            }
        }
    }
}