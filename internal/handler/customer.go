// getCustomers responds with the list of all customers as JSON.
func getCustomers(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, customers)
}

// postCustomers adds a customer from JSON received in the request body.
func postCustomers(c *gin.Context) {
    var newCustomer customer

    // Call BindJSON to bind the received JSON to
    // newCustomer.
    if err := c.BindJSON(&newCustomer); err != nil {
        return
    }

    // Add the new customer to the slice.
    customers = append(customers, newCustomer)
    c.IndentedJSON(http.StatusCreated, newCustomer)
}

// getCustomerByID locates the customer whose ID value matches the id
// parameter sent by the client, then returns that customer as a response.
func getCustomerByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of customers, looking for
    // a customer whose ID value matches the parameter.
    for _, a := range customers {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}