// internal/handler/customer.go
package handler

import (
    "github.com/gin-gonic/gin"
    "go-gin/internal/domain"
    "go-gin/internal/repository/memory"
)

type CustomerHandler struct {
    repo *memory.CustomerRepository  // Direct repository dependency
}

func NewCustomerHandler(repo *memory.CustomerRepository) *CustomerHandler {
    return &CustomerHandler{
        repo: repo,
    }
}

// Handler methods using repository directly
func (h *CustomerHandler) GetCustomers(c *gin.Context) {
    customers := h.repo.GetAll()
    c.JSON(200, customers)
}

func (h *CustomerHandler) GetCustomerByID(c *gin.Context) {
    id := c.Param("id")
    customer, err := h.repo.FindByID(id)
    if err != nil {
        c.JSON(404, gin.H{"error": "Customer not found"})
        return
    }
    c.JSON(200, customer)
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) {
    var customer domain.Customer
    if err := c.BindJSON(&customer); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    
    err := h.repo.Create(customer)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(201, customer)
}