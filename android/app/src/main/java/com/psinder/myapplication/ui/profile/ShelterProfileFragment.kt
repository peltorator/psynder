package com.psinder.myapplication.ui.profile

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.fragment.app.Fragment
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentShelterProfileBinding
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class ShelterProfileFragment : Fragment() {

    override fun onCreateView(
        inflater: LayoutInflater, container: ViewGroup?,
        savedInstanceState: Bundle?
    ): View? {
        // Inflate the layout for this fragment
        val binding = DataBindingUtil.inflate<FragmentShelterProfileBinding>(inflater,
            R.layout.fragment_shelter_profile,container,false)
//        binding.button.setOnClickListener {
//            it.findNavController().navigate(R.id)
//        }
        return binding.root
    }
}