package domain

// customer represents data about a record customer.
type customer struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Email string  `json:"email"`
    Phone  float64 `json:"phone"`
}

