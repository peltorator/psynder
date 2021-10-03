package com.psinder.myapplication

import android.os.Bundle
import androidx.fragment.app.Fragment
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.navigation.findNavController
import com.psinder.myapplication.databinding.FragmentProfileBinding

class ProfileFragment : Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        val binding = DataBindingUtil.inflate<FragmentProfileBinding>(inflater,
            R.layout.fragment_profile,container,false)
        binding.button.setOnClickListener {
            it.findNavController().navigate(R.id.action_profileFragment_to_swipeFragment)
        }
        return binding.root
    }
}