package com.example.auth.service

import org.springframework.stereotype.Service
import org.springframework.context.annotation.Lazy

fun interface AuthService {
    fun authenticate(username: String, password: String): Boolean
}

@Service("eagerAuthService")
class EagerAuthService : AuthService {
    init {
        println("[EAGER] EagerAuthService is being created.")
    }

    override fun authenticate(username: String, password: String): Boolean {
        println()
        println("Eager singleton instance was used")
        println()
        if ((username == "admin" && password == "password") ||
            (username == "slawosz uznanski" && password == "poznanski")) {
            return true;
        }
        return false;
    }
}

@Service("lazyAuthService")
@Lazy
class LazyAuthService : AuthService {
    init {
        println("[LAZY] LazyAuthService is being created.")
    }

    override fun authenticate(username: String, password: String): Boolean {
        println()
        println("Lazy singleton instance was used")
        println()
        if ((username == "admin" && password == "password") ||
            (username == "slawosz uznanski" && password == "poznanski")) {
            return true;
        }
        return false;
    }
}
