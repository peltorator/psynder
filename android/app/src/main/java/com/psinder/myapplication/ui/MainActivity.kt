package com.psinder.myapplication.ui


import android.os.Bundle
import androidx.activity.viewModels
import androidx.appcompat.app.AppCompatActivity
import androidx.lifecycle.Lifecycle
import androidx.lifecycle.lifecycleScope
import androidx.lifecycle.repeatOnLifecycle
import androidx.navigation.findNavController
import by.kirich1409.viewbindingdelegate.viewBinding
import com.psinder.myapplication.R
import com.psinder.myapplication.databinding.ActivityMainBinding
import com.psinder.myapplication.entity.AccountKind
import dagger.hilt.android.AndroidEntryPoint
import kotlinx.coroutines.launch
import kotlinx.coroutines.flow.collect

@AndroidEntryPoint
class MainActivity : AppCompatActivity(R.layout.activity_main) {

    private val viewBinding by viewBinding(ActivityMainBinding::bind)

    private val viewModel: MainViewModel by viewModels()

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        subscribeToAuthorizationStatus()
    }

    private fun subscribeToAuthorizationStatus() {
        lifecycleScope.launch {
            lifecycle.repeatOnLifecycle(Lifecycle.State.STARTED) {
                viewModel.accountKindFlow().collect {
                    showSuitableNavigationFlow(it)
                }
            }
        }
    }

    // This method have to be idempotent. Do not override restored backstack.
    private fun showSuitableNavigationFlow(accountKind: AccountKind) {
        val navController = findNavController(R.id.mainActivityNavigationHost)
        when (accountKind) {
            AccountKind.PERSON -> {
                if (navController.backQueue.any { it.destination.id == R.id.user_nav_graph }) {
                    return
                }
                navController.navigate(R.id.action_userNavGraph)
            }
            AccountKind.SHELTER -> {
                if (navController.backQueue.any { it.destination.id == R.id.user_nav_graph }) {
                    return
                }
                navController.navigate(R.id.action_shelterNavGraph)
            }
            AccountKind.UNDEFINED -> {
                if (navController.backQueue.any { it.destination.id == R.id.guest_nav_graph }) {
                    return
                }
                navController.navigate(R.id.action_guestNavGraph)
            }
        }
    }
}