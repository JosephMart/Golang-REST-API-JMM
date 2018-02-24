package main

// TODO: make these all in scope while using a .env file

// BaseURL is base of the api of the app
var BaseURL = "/api"

// Set up a global string for our secret
var mySigningKey = []byte("secret1")

// Init books var as a slice Book struct
var books []Book

// Auth0 constants
var SECRET = []byte("lol")
var AUDIENCE = []string{"bookstore"}
var AUTH0PATH = "https://martinsen.auth0.com/"
