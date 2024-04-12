package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

// CustomerLoader is a struct that represents an implementation of the Loader interface for customer json files
type CustomerLoader struct {
	FilePath   string
	Repository internal.RepositoryCustomer
}

// NewCustomerLoaderJSON is a method that creates a new customer loader json
func NewCustomerLoaderJSON(filePath string, repository internal.RepositoryCustomer) *CustomerLoader {
	return &CustomerLoader{filePath, repository}
}

// CustomerJSON ia struct that contains the information of a customer attributes in json format
type CustomerJSON struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition int    `json:"condition"`
}

// Load is a method that loads the costumers into bd
func (c *CustomerLoader) Load() (err error) {
	// Read file
	data, err := os.ReadFile(c.FilePath)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	var customers []internal.Customer
	if err = json.Unmarshal(data, &customers); err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	for _, customer := range customers {
		err = c.Repository.Save(&customer)
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
			return
		}
	}
	return
}
