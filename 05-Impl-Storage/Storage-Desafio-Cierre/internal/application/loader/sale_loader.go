package loader

import (
	"app/internal"
	"encoding/json"
	"fmt"
	"os"
)

// SaleLoader is a struct that represents an implementation of the loader interface for sale json files
type SaleLoader struct {
	FilePath   string
	Repository internal.RepositorySale
}

// NewSaletLoaderJSON is a method that creates a new sale loader json
func NewSaleLoaderJSON(filepath string, repository internal.RepositorySale) *SaleLoader {
	return &SaleLoader{filepath, repository}
}

// Load is a method that loads the sales into db
func (p *SaleLoader) Load() (err error) {
	// Read file
	data, err := os.ReadFile(p.FilePath)
	if err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	var sales []internal.Sale
	if err = json.Unmarshal(data, &sales); err != nil {
		err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
		return
	}

	for _, sale := range sales {
		err = p.Repository.Save(&sale)
		if err != nil {
			err = fmt.Errorf("%w: %s", ErrLoaderInternal, err)
			return
		}
	}
	return
}
