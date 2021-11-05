package com.psinder.myapplication.ui.doglist

import android.os.Bundle
import android.view.View
import androidx.core.view.isVisible
import androidx.fragment.app.Fragment
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.ViewModelProvider
import by.kirich1409.viewbindingdelegate.viewBinding
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.findNavController
import androidx.navigation.navGraphViewModels
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentDogListBinding
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch

class DogListFragment : Fragment(R.layout.fragment_dog_list) {
    val viewModel: DogListViewModel
            by navGraphViewModels(R.id.shelter_nav_graph)
    private val viewBinding by viewBinding(FragmentDogListBinding::bind)

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        setupRecyclerView()
        viewLifecycleOwner.lifecycleScope.launch {
            viewLifecycleOwner.lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                viewModel.viewState.collect {
                        viewState ->
                    renderViewState(viewState)
                }
            }
        }
        viewBinding.floatingActionButton.setOnClickListener {
            view.findNavController().navigate(R.id.action_dogListFragment_to_editDogFragment2)
        }
    }

    private fun setupRecyclerView(): DogAdapter {
        val recyclerView = viewBinding.dogsRecyclerView
        val adapter = DogAdapter()
        recyclerView.adapter = adapter
        return adapter
    }

    private fun renderViewState(viewState: DogListViewModel.ViewState) {
        when (viewState) {
            is DogListViewModel.ViewState.Loading -> {
                viewBinding.dogsRecyclerView.isVisible = false
                viewBinding.progressBar.isVisible = true
            }
            is DogListViewModel.ViewState.Data -> {
                viewBinding.dogsRecyclerView.isVisible = true
                (viewBinding.dogsRecyclerView.adapter as DogAdapter).apply {
                    psynasList = viewState.psynaList
                    notifyDataSetChanged()
                }
                viewBinding.progressBar.isVisible = false
            }
        }
    }
}