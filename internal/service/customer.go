package service

import (
    "context"
    "fmt"
    "your-project/internal/domain"
    "your-project/internal/repository/interfaces"
)

// CustomerService interface defines the business operations for customers
type CustomerService interface {
    GetAllCustomers(ctx context.Context) ([]domain.Customer, error)
    GetCustomerByID(ctx context.Context, id string) (*domain.Customer, error)
    CreateCustomer(ctx context.Context, customer *domain.Customer) error
    UpdateCustomer(ctx context.Context, customer *domain.Customer) error
    DeleteCustomer(ctx context.Context, id string) error
}

// customerService implements CustomerService interface
type customerService struct {
    repo interfaces.CustomerRepository
}

// NewCustomerService creates a new instance of CustomerService
func NewCustomerService(repo interfaces.CustomerRepository) CustomerService {
    return &customerService{
        repo: repo,
    }
}

// GetAllCustomers retrieves all customers with business logic applied
func (s *customerService) GetAllCustomers(ctx context.Context) ([]domain.Customer, error) {
    // Get customers from repository
    customers, err := s.repo.FindAll(ctx)
    if err != nil {
        return nil, fmt.Errorf("error fetching customers: %w", err)
    }

    // Here you could add business logic like:
    // - Filtering based on business rules
    // - Enriching customer data
    // - Applying transformations

    return customers, nil
}

// GetCustomerByID retrieves a specific customer
func (s *customerService) GetCustomerByID(ctx context.Context, id string) (*domain.Customer, error) {
    if id == "" {
        return nil, fmt.Errorf("customer ID cannot be empty")
    }

    customer, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return nil, fmt.Errorf("error fetching customer: %w", err)
    }

    return customer, nil
}

// CreateCustomer handles customer creation with business validation
func (s *customerService) CreateCustomer(ctx context.Context, customer *domain.Customer) error {
    // Validate customer data
    if err := s.validateCustomer(customer); err != nil {
        return fmt.Errorf("invalid customer data: %w", err)
    }

    // Check for existing customer with same email
    existing, err := s.repo.FindByEmail(ctx, customer.Email)
    if err == nil && existing != nil {
        return fmt.Errorf("customer with email %s already exists", customer.Email)
    }

    // Generate unique ID if not provided
    if customer.ID == "" {
        customer.ID = generateUniqueID() // You would implement this
    }

    // Additional business logic before creation
    // - Normalize phone number
    // - Format email to lowercase
    // - Add timestamps
    customer.Email = strings.ToLower(customer.Email)
    customer.Phone = normalizePhoneNumber(customer.Phone)

    // Create customer in repository
    if err := s.repo.Create(ctx, customer); err != nil {
        return fmt.Errorf("error creating customer: %w", err)
    }

    return nil
}

// UpdateCustomer handles customer updates with validation
func (s *customerService) UpdateCustomer(ctx context.Context, customer *domain.Customer) error {
    if customer.ID == "" {
        return fmt.Errorf("customer ID cannot be empty")
    }

    // Validate customer data
    if err := s.validateCustomer(customer); err != nil {
        return fmt.Errorf("invalid customer data: %w", err)
    }

    // Check if customer exists
    existing, err := s.repo.FindByID(ctx, customer.ID)
    if err != nil {
        return fmt.Errorf("customer not found: %w", err)
    }

    // Check email uniqueness if email is being changed
    if existing.Email != customer.Email {
        emailExists, _ := s.repo.FindByEmail(ctx, customer.Email)
        if emailExists != nil {
            return fmt.Errorf("email %s is already in use", customer.Email)
        }
    }

    // Apply business logic before update
    customer.Email = strings.ToLower(customer.Email)
    customer.Phone = normalizePhoneNumber(customer.Phone)

    if err := s.repo.Update(ctx, customer); err != nil {
        return fmt.Errorf("error updating customer: %w", err)
    }

    return nil
}

// DeleteCustomer handles customer deletion with business rules
func (s *customerService) DeleteCustomer(ctx context.Context, id string) error {
    if id == "" {
        return fmt.Errorf("customer ID cannot be empty")
    }

    // Check if customer exists
    _, err := s.repo.FindByID(ctx, id)
    if err != nil {
        return fmt.Errorf("customer not found: %w", err)
    }

    // Here you could add business logic like:
    // - Check if customer can be deleted
    // - Archive instead of delete
    // - Handle related data

    if err := s.repo.Delete(ctx, id); err != nil {
        return fmt.Errorf("error deleting customer: %w", err)
    }

    return nil
}

// validateCustomer performs business validation on customer data
func (s *customerService) validateCustomer(customer *domain.Customer) error {
    if customer == nil {
        return fmt.Errorf("customer cannot be nil")
    }

    if customer.Name == "" {
        return fmt.Errorf("customer name is required")
    }

    if !isValidEmail(customer.Email) {
        return fmt.Errorf("invalid email format")
    }

    if !isValidPhone(customer.Phone) {
        return fmt.Errorf("invalid phone number format")
    }

    return nil
}

// Helper functions (you would implement these)
func isValidEmail(email string) bool {
    // Implement email validation logic
    return true
}

func isValidPhone(phone float64) bool {
    // Implement phone validation logic
    return true
}

func normalizePhoneNumber(phone float64) float64 {
    // Implement phone normalization logic
    return phone
}

func generateUniqueID() string {
    // Implement unique ID generation
    return ""
}