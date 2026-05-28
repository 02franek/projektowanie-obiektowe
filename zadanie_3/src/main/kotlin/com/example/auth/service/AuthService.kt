package com.example.auth.service

import org.springframework.stereotype.Service

@Service
class AuthService {
    init {
        println("[EAGER] AuthService is being created.")
    }

    fun authenticate(username: String, password: String): Boolean {
        if ((username == "admin" && password == "password") ||
            (username == "slawosz uznanski" && password == "poznanski")) {
            return true;
        }
        return false;
    }
}

