package com.psinder.myapplication.ui.editdog

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.navigation.findNavController
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentEditDogBinding
import com.psinder.myapplication.network.Psyna
import dagger.hilt.android.AndroidEntryPoint

@AndroidEntryPoint
class EditDogFragment : Fragment(R.layout.fragment_edit_dog) {
    private val viewBinding by viewBinding(FragmentEditDogBinding::bind)
    val viewModel: EditDogViewModel by viewModels()

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewBinding.saveButton.setOnClickListener {

            viewModel.addPsyna(readPsyna())
            view.findNavController().navigate(R.id.action_editDogFragment2_to_dogListFragment)
        }
    }

    fun readPsyna(): Psyna {
        val psynaName = viewBinding.dogNameEditText.editText?.text?.toString() ?: "Loma"
        val psynaDescription = viewBinding.description.editText?.text?.toString() ?: ""
        val psynaURL = "https://images.dog.ceo/breeds/terrier-lakeland/n02095570_4560.jpg"
        return Psyna(239, psynaName, "", psynaDescription, psynaURL)
    }
}