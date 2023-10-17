package main

import (
    "gtp/routes"
    "log"
    "net/http"
)

func main() {
    log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/products/", routes.HandleAllProducts)
    http.HandleFunc("/product/", routes.HandleProduct)
	http.HandleFunc("/reviews/", routes.HandleProductReviews)
    http.HandleFunc("/images/", routes.HandleProductImages)
    http.ListenAndServe("localhost:8080", nil)
}