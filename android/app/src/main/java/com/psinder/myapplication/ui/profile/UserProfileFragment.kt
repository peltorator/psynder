package com.psinder.myapplication.ui.profile

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentUserProfileBinding

class UserProfileFragment : Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        val binding = DataBindingUtil.inflate<FragmentUserProfileBinding>(inflater,
            R.layout.fragment_user_profile,container,false)
//        binding.button.setOnClickListener {
//            it.findNavController().navigate(R.id)
//        }
        return binding.root
    }
}