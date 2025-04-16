
### Run the web server
```bash
go run cmd/api/main.go
```

### Run go command to manage dependencies
```bash
go mod tidy
```

### Get all customers
```bash
curl -X GET http://localhost:8080/customers
```

### Remove go.sum
```bash
rm go.sum
```

### Get customer by id
```bash
curl -X GET http://localhost:8080/customers/1
```

### Create customer
```bash
curl -X POST http://localhost:8080/customers \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john.doe@example.com",
    "phone": 1234567890
  }'
```

### Run test
```bash
./test/test_api.sh
```

### PowerShell commands
```powershell
# Get All Customers
Invoke-RestMethod -Method GET -Uri "http://localhost:8080/customers" -ContentType "application/json"

# Get Customer by ID
Invoke-RestMethod -Method GET -Uri "http://localhost:8080/customers/1" -ContentType "application/json"

# Create New Customer
$body = @{
    name = "John Doe"
    email = "john.doe@example.com"
    phone = 1234567890
} | ConvertTo-Json

Invoke-RestMethod -Method POST -Uri "http://localhost:8080/customers" -ContentType "application/json" -Body $body
```
