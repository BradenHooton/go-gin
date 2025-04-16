package main

import (
    "github.com/gin-gonic/gin"
    "go-gin/internal/domain"
    "go-gin/internal/repository/memory"
    "go-gin/internal/handler"
)

// customers slice to seed record customer data.
var customers = []domain.Customer{
    {ID: "1", Name: "John Doe", Email: "john.doe@example.com", Phone: 1234567890},
    {ID: "2", Name: "Jane Smith", Email: "jane.smith@example.com", Phone: 9876543210},
    {ID: "3", Name: "Alice Johnson", Email: "alice.johnson@example.com", Phone: 5555555555},
}

func main() {
    // Initialize repository
    repo := memory.NewCustomerRepository()
    
    // Initialize handler
    customerHandler := handler.NewCustomerHandler(repo)
    
    // Setup router
    router := gin.Default()
    
    // Define routes using the handler methods
    router.GET("/customers", customerHandler.GetCustomers)
    router.GET("/customers/:id", customerHandler.GetCustomerByID)
    router.POST("/customers", customerHandler.CreateCustomer)

    router.Run("localhost:8080")
}
