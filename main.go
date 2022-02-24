package main

import "URL_Shortener/router"

func main() {
	r := router.NewRouter()

	r.Run(":8080")
}
