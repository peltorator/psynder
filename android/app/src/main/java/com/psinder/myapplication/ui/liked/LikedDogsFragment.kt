package com.psinder.myapplication.ui.liked

import android.os.Bundle
import android.util.Log
import android.view.View
import androidx.core.view.isVisible
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.recyclerview.widget.GridLayoutManager
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentLikedBinding
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch

@AndroidEntryPoint
class LikedDogsFragment : Fragment(R.layout.fragment_liked) {
    val viewModel: LikedDogsViewModel by viewModels()
    private val viewBinding by viewBinding(FragmentLikedBinding::bind)

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        setupRecyclerView()
        viewLifecycleOwner.lifecycleScope.launch {
            viewLifecycleOwner.lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                viewModel.viewState.collect { viewState -> renderViewState(viewState) }
            }
        }
    }

    override fun onResume() {
        super.onResume()
        Log.d("liked", "resumed")

    }

    private fun setupRecyclerView(): LikedDogAdapter {
        val recyclerView = viewBinding.likedDogsRecyclerView
        recyclerView.layoutManager = GridLayoutManager(context,2)
        val adapter = LikedDogAdapter()
        recyclerView.adapter = adapter

        return adapter
    }

    private fun renderViewState(viewState: LikedDogsViewModel.ViewState) {
        when (viewState) {
            is LikedDogsViewModel.ViewState.Loading -> {
                viewBinding.likedDogsRecyclerView.isVisible = false
            }
            is LikedDogsViewModel.ViewState.Data -> {
                viewBinding.likedDogsRecyclerView.isVisible = true
                (viewBinding.likedDogsRecyclerView.adapter as LikedDogAdapter).apply {
                    psynasList = viewState.psynaList
                    Log.d("liked", viewState.psynaList.toString())
                    notifyDataSetChanged()
                }
            }
        }
    }

}