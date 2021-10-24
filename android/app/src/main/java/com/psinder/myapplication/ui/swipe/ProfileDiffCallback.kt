package com.psinder.myapplication.ui.swipe


import androidx.recyclerview.widget.DiffUtil
import com.psinder.myapplication.entity.Profile

class ProfileDiffCallback(
    private val old: List<Profile>,
    private val new: List<Profile>
) : DiffUtil.Callback() {

    override fun getOldListSize(): Int {
        return old.size
    }

    override fun getNewListSize(): Int {
        return new.size
    }

    override fun areItemsTheSame(oldPosition: Int, newPosition: Int): Boolean {
        return old[oldPosition].id == new[newPosition].id
    }

    override fun areContentsTheSame(oldPosition: Int, newPosition: Int): Boolean {
        return old[oldPosition] == new[newPosition]
    }

}
