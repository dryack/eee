package main

type player struct {
	Name     string  `json:"name"`
	Password string  `json:"password"` // salted Argon2 hash
	Empire   *Empire `json:"empire"`
}

// playerLogin is the point at which a user is greeted.
// If they haven't created an empire, they can quit, or they can register a new empire.
// Otherwise they can login with their Name and Password (which is matched against the hash we have stored.
//		After 3 attempts to login, the user is disconnected.
func playerLogin() {

}

// playerRegister is called if a user wishes to create a new empire, it returns a an instance of Empire
func playerRegister() {

}
