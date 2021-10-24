package com.psinder.myapplication.ui.swipe

import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.databinding.DataBindingUtil
import androidx.recyclerview.widget.RecyclerView
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.CardViewProfileBinding
import com.psinder.myapplication.entity.Profile


class ProfilesAdapter : RecyclerView.Adapter<ProfilesAdapter.ProfileViewHolder>() {

    private var profiles: List<Profile>? = null

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int) = ProfileViewHolder(
        DataBindingUtil.inflate(
            LayoutInflater.from(parent.context),
            R.layout.card_view_profile,
            parent,
            false
        )
    )

    override fun getItemCount() = profiles?.size ?: 0

    override fun onBindViewHolder(holder: ProfileViewHolder, position: Int) {
        profiles?.let {
            holder.binding.profile = it[position]
            holder.binding.executePendingBindings()
        }
    }

    fun setProfiles(profiles: List<Profile>) {
        this.profiles = profiles
    }

    fun getProfiles(): List<Profile> {
        return this.profiles ?: emptyList()
    }

    inner class ProfileViewHolder(val binding: CardViewProfileBinding) :
        RecyclerView.ViewHolder(binding.root)

}