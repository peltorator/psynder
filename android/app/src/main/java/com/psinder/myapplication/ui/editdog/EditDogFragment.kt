package com.psinder.myapplication.ui.editdog

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.findNavController
import androidx.navigation.navGraphViewModels
import com.psinder.myapplication.R
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.databinding.FragmentEditDogBinding
import com.psinder.myapplication.network.Psyna
import com.psinder.myapplication.ui.doglist.DogListViewModel


class EditDogFragment : Fragment(R.layout.fragment_edit_dog) {
    private val viewBinding by viewBinding(FragmentEditDogBinding::bind)
    val viewModel: DogListViewModel
            by navGraphViewModels(R.id.shelter_nav_graph)

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        viewBinding.saveButton.setOnClickListener {
            viewModel.addPsyna(
                readPsyna()
            )
            view.findNavController().navigate(R.id.action_editDogFragment2_to_dogListFragment)
        }
    }

    fun readPsyna(): Psyna {
        val psynaName = viewBinding.dogNameEditText.editText?.text.toString() ?: "Ivan"
        val psynaDescription = viewBinding.description.editText?.text.toString() ?: "Olejnik"
        val psynaURL = "https://sun9-29.userapi.com/impg/a_X3O2c0SMxVUL3Aa9MDzEadE3ef0A3arZnfxA/aJ_vfjZ3DmA.jpg?size=512x512&quality=96&sign=2a2b521e2b6cf2bd6fdd7984bccc44a6&type=album"
        return Psyna(239, psynaName, psynaDescription, psynaURL)
    }
}