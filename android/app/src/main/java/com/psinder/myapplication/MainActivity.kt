package com.psinder.myapplication

import android.os.Bundle
import android.util.Log
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.findNavController
import androidx.navigation.fragment.NavHostFragment
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.NavigationUI
import androidx.navigation.ui.setupActionBarWithNavController
import androidx.navigation.ui.setupWithNavController
import com.google.android.material.bottomnavigation.BottomNavigationView
import com.psinder.myapplication.databinding.ActivityMainBinding


import androidx.annotation.NonNull
import androidx.fragment.app.FragmentManager
import androidx.fragment.app.FragmentTransaction
import com.psinder.myapplication.swipe.SwipeFragment


class MainActivity : AppCompatActivity() {

    private lateinit var binding: ActivityMainBinding

//    private val mOnNavigationItemSelectedListener =
//        BottomNavigationView.OnNavigationItemSelectedListener { item ->
//            val fragmentManager: FragmentManager = supportFragmentManager
//            val transaction: FragmentTransaction = fragmentManager.beginTransaction()
//            when (item.itemId) {
//                R.id -> {
//                    transaction.replace(R.id.container, TaskFragment()).commit()
//                    return@OnNavigationItemSelectedListener true
//                }
//                R.id.navigation_calendar -> {
//                    transaction.replace(R.id.container, CalendarFragment()).commit()
//                    return@OnNavigationItemSelectedListener true
//                }
//                R.id.navigation_app -> {
//                    transaction.replace(R.id.container, AppFragment()).commit()
//                    return@OnNavigationItemSelectedListener true
//                }
//            }
//            false
//        }

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
//
        binding = ActivityMainBinding.inflate(layoutInflater)
        setContentView(binding.root)
//
        val navView: BottomNavigationView = binding.navView
////

        val navHostFragment = supportFragmentManager.findFragmentById(R.id.nav_host_fragment_activity_main) as NavHostFragment
        val navController = navHostFragment.navController
//        val navController = findNavController(R.id.nav_host_fragment_activity_main)
        // Passing each menu ID as a set of Ids because each
        // menu should be considered as top level destinations.
//        val appBarConfiguration = AppBarConfiguration(
//            setOf(
//                R.id.navigation_home, R.id.navigation_dashboard, R.id.navigation_notifications
//            )
//        )
//        setupActionBarWithNavController(navController, appBarConfiguration)
         navView.setupWithNavController(navController)
//        NavigationUI.setupActionBarWithNavController(this, navController)
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