package com.psinder.myapplication

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.fragment.NavHostFragment
import androidx.navigation.ui.setupWithNavController
import com.google.android.material.bottomnavigation.BottomNavigationView
import com.psinder.myapplication.databinding.ActivityMainBinding


import androidx.fragment.app.FragmentManager
import androidx.fragment.app.FragmentTransaction
import com.psinder.myapplication.ui.chat.ChatFragment
import com.psinder.myapplication.ui.swipe.SwipeFragment
import com.psinder.myapplication.ui.profile.ProfileFragment


class MainActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMainBinding
    lateinit var token: String

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)

        val intent = intent

        token = intent.getStringExtra("TOKEN").toString()

        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)
        val navView: BottomNavigationView = binding.navView

        val navHostFragment = supportFragmentManager.findFragmentById(R.id.nav_host_fragment_activity_main) as NavHostFragment
        val navController = navHostFragment.navController

        navView.setupWithNavController(navController)
        binding.navView.setOnItemSelectedListener{
            val fragmentManager = supportFragmentManager
            val transaction = fragmentManager.beginTransaction()
            when (it.itemId) {

                R.id.navigation_swipe -> {
                    Log.d("HelloWorld", "Swipe")
                    transaction.replace(R.id.nav_host_fragment_activity_main, SwipeFragment()).commit()
                    return@setOnItemSelectedListener true
                }
                R.id.navigation_chat -> {

                    Log.d("HelloWorld", "Chat")
                    transaction.replace(R.id.nav_host_fragment_activity_main, ChatFragment()).commit()
                    return@setOnItemSelectedListener true
                }
                R.id.navigation_profile -> {

                    Log.d("HelloWorld", "Profile")
                    transaction.replace(R.id.nav_host_fragment_activity_main, ProfileFragment()).commit()
                    return@setOnItemSelectedListener true
                }

            }
            false
        }
        val fragmentManager: FragmentManager = supportFragmentManager
        val transaction: FragmentTransaction = fragmentManager.beginTransaction()
        transaction.replace(R.id.nav_host_fragment_activity_main, SwipeFragment()).commit()
    }
}