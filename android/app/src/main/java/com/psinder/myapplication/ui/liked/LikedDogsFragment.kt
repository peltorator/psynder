package com.psinder.myapplication.ui.liked

import android.annotation.SuppressLint
import android.os.Bundle
import android.text.style.ClickableSpan
import android.util.Log
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.view.animation.LinearInterpolator
import androidx.core.view.isVisible
import androidx.databinding.DataBindingUtil
import androidx.databinding.ViewDataBinding
import androidx.fragment.app.Fragment
import androidx.fragment.app.viewModels
import androidx.lifecycle.Lifecycle
import by.kirich1409.viewbindingdelegate.viewBinding
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.findNavController
import androidx.navigation.navGraphViewModels
import androidx.recyclerview.widget.GridLayoutManager
import androidx.recyclerview.widget.RecyclerView
import com.facebook.drawee.backends.pipeline.Fresco
import com.facebook.drawee.gestures.GestureDetector
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.FragmentDogListBinding
import com.psinder.myapplication.databinding.FragmentLikedBinding
import com.psinder.myapplication.databinding.FragmentSwipeBinding
import com.psinder.myapplication.databinding.LikedDogBinding
import com.psinder.myapplication.repository.AuthRepository
import com.psinder.myapplication.ui.doglist.DogAdapter
import com.psinder.myapplication.ui.swipe.ProfilesAdapter
import com.psinder.myapplication.ui.swipe.SwipeViewModel
import com.yuyakaido.android.cardstackview.CardStackLayoutManager
import com.yuyakaido.android.cardstackview.CardStackView
import com.yuyakaido.android.cardstackview.SwipeableMethod
import kotlinx.android.synthetic.main.fragment_liked.*
import kotlinx.coroutines.flow.collect
import kotlinx.coroutines.launch

class LikedDogsFragment : Fragment(R.layout.fragment_liked) {
    val viewModel: LikedDogsViewModel
            by navGraphViewModels(R.id.user_bottom_nav_graph)
    private val viewBinding by viewBinding(FragmentLikedBinding::bind)

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        setupRecyclerView()
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
                    notifyDataSetChanged()
                }
            }
        }
    }

}