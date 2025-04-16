package interfaces

import (
    "context"
    "go-gin/internal/domain"    
)

// CustomerRepository defines the contract for customer data operations
type CustomerRepository interface {
    // FindAll retrieves all customers
    FindAll(ctx context.Context) ([]domain.Customer, error)
    
    // FindByID retrieves a customer by their ID
    FindByID(ctx context.Context, id string) (*domain.Customer, error)
    
    // Create stores a new customer
    Create(ctx context.Context, customer *domain.Customer) error
    
    // Update modifies an existing customer
    Update(ctx context.Context, customer *domain.Customer) error
    
    // Delete removes a customer
    Delete(ctx context.Context, id string) error
    
    // FindByEmail finds a customer by email address
    FindByEmail(ctx context.Context, email string) (*domain.Customer, error)
}