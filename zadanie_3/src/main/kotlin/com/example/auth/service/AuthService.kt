package com.example.auth.service

import org.springframework.stereotype.Service

@Service
class AuthService {
    init {
        println("[EAGER] AuthService is being created.")
    }

    fun authenticate(username: String): Boolean {
        return true;
    }
}

