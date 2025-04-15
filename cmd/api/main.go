package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/YourGithub/go-gin/internal/domain"
)

// customers slice to seed record customer data.
var customers = []customer{
    {ID: "1", Name: "John Doe", Email: "john.doe@example.com", Phone: 1234567890},
    {ID: "2", Name: "Jane Smith", Email: "jane.smith@example.com", Phone: 9876543210},
    {ID: "3", Name: "Alice Johnson", Email: "alice.johnson@example.com", Phone: 5555555555},
}

func main() {
    router := gin.Default()
    router.GET("/customers", getCustomers)
    router.GET("/customers/:id", getCustomerByID)
    router.POST("/customers", postCustomers)

    router.Run("localhost:8080")
}
