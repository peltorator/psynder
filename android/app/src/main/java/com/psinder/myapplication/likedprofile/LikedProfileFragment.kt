package com.psinder.myapplication.likedprofile

import android.os.Bundle
import android.view.View
import android.widget.Toast
import androidx.fragment.app.Fragment
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.findNavController
import androidx.navigation.navGraphViewModels
import by.kirich1409.viewbindingdelegate.viewBinding
import com.bumptech.glide.Glide
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentLikedProfileBinding
import com.psinder.myapplication.repository.AuthRepository
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch


class LikedProfileFragment : Fragment(R.layout.fragment_liked_profile) {
    val defaultDog = "https://sun9-49.userapi.com/impf/phAQReMA5qa6Z67a19uwvb39PKdu6kL-MuetrA/mTJQrWPdv9s.jpg?size=1080x1027&quality=96&sign=764698d9ba05155df1d408c068264c29&type=album"

    val viewModel: LikedProfileViewModel
            by navGraphViewModels(R.id.user_bottom_nav_graph)
    private val viewBinding by viewBinding(FragmentLikedProfileBinding::bind)

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewModel.psynaId = (arguments?.getInt("psynaId") ?: 0)
        Glide.with(viewBinding.avatarImage)
            .load((arguments?.getString("photo") ?: defaultDog))
            .circleCrop()
            .into(viewBinding.avatarImage)

        viewBinding.backButton.setOnClickListener {
            view.findNavController().navigate(R.id.action_likedProfileFragment2_to_likedFragment)
        }
        viewLifecycleOwner.lifecycleScope.launch {
            viewLifecycleOwner.lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                viewModel.token.emit(
                    AuthRepository.token
                )
                viewModel.viewState.collect {
                        viewState ->
                    renderViewState(viewState)
                }
            }
        }
    }

    private fun renderViewState(viewState: LikedProfileViewModel.ViewState) {
        when (viewState) {
            is LikedProfileViewModel.ViewState.Loading -> {
                // loading
            }
            is LikedProfileViewModel.ViewState.Data -> {
                viewBinding.cityTextView.text = viewState.shelter?.city
                viewBinding.addressTextView.text = viewState.shelter?.address
                viewBinding.phoneTextView.text = viewState.shelter?.phone
            }
        }
    }

}