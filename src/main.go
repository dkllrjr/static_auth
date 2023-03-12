package main

import (
    "os"
    "static_auth/database"
    "static_auth/routing"
)

//  ──────────────────────────────────────────────────────────────────────────

func main() {
	
	// database
	database.InitDB("website/" + os.Getenv("SA_DATABASE"))
	defer database.DB.Close()

	// router
	//router := routing.InitRouter()
	//router.Run(os.Getenv("SA_URL"))

	// router
	router := routing.InitRouter()
    router.Logger.Fatal(router.Start(os.Getenv("SA_URL")))
}

// ──────────────────────────────────────────────────────────────────────────


