package main

import (
    "encoding/json"
    "net/http"
    "strconv"
)

// Product represents an item in the store.
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

// filterProductsByPrice is a pure function that filters products under a certain price.
func filterProductsByPrice(products []Product, maxPrice float64) []Product {
    var filteredProducts []Product
    for _, product := range products {
        if product.Price <= maxPrice {
            filteredProducts = append(filteredProducts, product)
        }
    }
    // Since the output is solely determined by the input without any side effects, this function is pure.
    return filteredProducts
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
    // Define a sample list of products.
    products := []Product{
        {ID: 1, Name: "Laptop", Price: 1000.00},
        {ID: 2, Name: "Mouse", Price: 25.00},
        {ID: 3, Name: "Keyboard", Price: 45.00},
    }

    // Parse maxPrice query parameter.
    maxPriceParam := r.URL.Query().Get("maxPrice")
    maxPrice, err := strconv.ParseFloat(maxPriceParam, 64)
    if err != nil {
        http.Error(w, "Invalid maxPrice parameter", http.StatusBadRequest)
        return
    }

    // Use the pure function to filter products.
    filteredProducts := filterProductsByPrice(products, maxPrice)

    // Marshal the filtered products to JSON for the response.
    jsonResponse, err := json.Marshal(filteredProducts)
    if err != nil {
        http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
        return
    }

    // Set the Content-Type header and write the JSON response.
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonResponse)
}

func main() {
    http.HandleFunc("/products", productsHandler)
    http.ListenAndServe(":8080", nil)
}
