package com.psinder.myapplication.ui.liked

import android.content.Context
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import android.widget.Toast
import androidx.core.os.bundleOf
import androidx.navigation.findNavController
import androidx.recyclerview.widget.RecyclerView
import com.bumptech.glide.Glide
import com.psinder.myapplication.R
import com.psinder.myapplication.network.Psyna
import com.psinder.myapplication.ui.doglist.DogAdapter

// TODO: move to entity or API


class LikedDogAdapter : RecyclerView.Adapter<LikedDogAdapter.ViewHolder>() {

    var psynasList: List<Psyna> = emptyList()

    class ViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {

        val avatarImageView = itemView.findViewById<ImageView>(R.id.avatarImageView)

        val dogNameTextView = itemView.findViewById<TextView>(R.id.dogNameTextView)

        val dogDescriptionTextView = itemView.findViewById<TextView>(R.id.dogDescription)
    }

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): ViewHolder {
        val itemView = LayoutInflater.from(parent.context).inflate(
            R.layout.liked_dog, parent, false
        )

        return ViewHolder(itemView)
    }

    override fun onBindViewHolder(holder: ViewHolder, position: Int) {
        Glide.with(holder.avatarImageView)
            .load(psynasList[position].photoLink)
            .circleCrop()
            .into(holder.avatarImageView)

        holder.itemView.setOnClickListener {
//            Toast.makeText(holder.itemView.context, "Hello, " + psynasList[position].name, Toast.LENGTH_LONG).show()
            //we can then create an intent here and start a new activity
            holder.itemView.findNavController().navigate(R.id.action_likedFragment_to_likedProfileFragment2,
            bundleOf(
                "psynaId" to psynasList[position].id,
                "photo" to psynasList[position].photoLink
            ))
            //with our data

        }
        holder.dogNameTextView.text = psynasList[position].name
        holder.dogDescriptionTextView.text = psynasList[position].description
    }

    override fun getItemCount(): Int = psynasList.size
}