package routes

func Start() {
	router := setRouter()

	// Start listening and serving requests
	router.Run("localhost:5000")
}
