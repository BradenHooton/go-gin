// internal/repository/memory/customer_repository.go
package memory

import (
    "go-gin/internal/domain"
	"fmt"
)

type CustomerRepository struct {
    customers []domain.Customer
}

func NewCustomerRepository() *CustomerRepository {
    // Initialize with sample data
    return &CustomerRepository{
        customers: []domain.Customer{
            {ID: "1", Name: "John Doe", Email: "john@example.com", Phone: 1234567890},
            {ID: "2", Name: "Jane Smith", Email: "jane@example.com", Phone: 9876543210},
			{ID: "3", Name: "Alice Johnson", Email: "alice@example.com", Phone: 5555555555},
			{ID: "4", Name: "Bob Brown", Email: "bob@example.com", Phone: 1111111111},
			{ID: "5", Name: "Charlie Davis", Email: "charlie@example.com", Phone: 2222222222},
			{ID: "6", Name: "Diana Evans", Email: "diana@example.com", Phone: 3333333333},
			{ID: "7", Name: "Ethan Foster", Email: "ethan@example.com", Phone: 4444444444},
			{ID: "8", Name: "Fiona Green", Email: "fiona@example.com", Phone: 5555555555},
        },
    }
}

func (r *CustomerRepository) GetAll() []domain.Customer {
    return r.customers
}

func (r *CustomerRepository) FindByID(id string) (*domain.Customer, error) {
    for _, c := range r.customers {
        if c.ID == id {
            return &c, nil
        }
    }
    return nil, fmt.Errorf("customer not found")
}

func (r *CustomerRepository) Create(customer domain.Customer) error {
    r.customers = append(r.customers, customer)
    return nil
}