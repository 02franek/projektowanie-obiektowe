package com.example.auth.controller

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

data class User(
    val id: Int,
    val username: String,
    val role: String,
    val salary: Int
)

@RestController
@RequestMapping("/api")
class ApiController {
    @GetMapping("/users")
    fun getUsers(): List<User> {
        return listOf(
            User(1, "slawosz uznanski", "kosmonauta", 2137),
	        User(2, "slawosz poznanski", "astrolog", 43422),
	        User(3, "slawomir poznansko-uznanski", "muzyk", 22222),
	        User(4, "admin", "informatyk", 2000)
        )
    }
}