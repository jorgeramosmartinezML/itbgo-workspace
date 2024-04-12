package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

// ProductLoader is a struct that represents an implementation of the loader interface for product json files
type ProductLoader struct {
	FilePath   string
	Repository internal.RepositoryProduct
}

// NewProductLoaderJSON is a method that creates a new product loader json
func NewProductLoaderJSON(filePath string, repository internal.RepositoryProduct) *ProductLoader {
	return &ProductLoader{filePath, repository}
}

// Load is a method that loads the products into db
func (p *ProductLoader) Load() (err error) {
	// Read file
	data, err := os.ReadFile(p.FilePath)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	var products []internal.Product
	if err = json.Unmarshal(data, &products); err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	for _, product := range products {
		err = p.Repository.Save(&product)
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
			return
		}
	}
	return
}
