package com.psinder.myapplication.ui.profile

import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentUserProfileBinding

class UserProfileFragment : Fragment(R.layout.fragment_user_profile) {
    private val viewBinding by viewBinding(FragmentUserProfileBinding::bind)
    private val viewModel: UserProfileViewModel by viewModels()
}