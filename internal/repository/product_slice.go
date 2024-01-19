package repository

import "github.com/usuario/repositorio/internal"

// NewProductSlice returns a new instance of ProductSlice
func NewProductSlice(products []internal.Product, lastID int) *ProductSlice {
	// default values
	defaultProducts := make([]internal.Product, 0)
	defaultLastID := 0
	if products != nil {
		defaultProducts = products
	}
	if lastID > 0 {
		defaultLastID = lastID
	}

	return &ProductSlice{
		products: defaultProducts,
		lastID:  defaultLastID,
	}
}

// ProductSlice is an implementation of ProductRepository
type ProductSlice struct {
	// products is a slice of Product
	products []internal.Product
	// lastID is the last used ID
	lastID int
}

// FindAll returns all products
func (p *ProductSlice) FindAll() (pr []internal.Product, err error) {
	// make a copy of the slice
	pr = make([]internal.Product, len(p.products))
	copy(pr, p.products)

	return
}

// FindByID returns a product by ID
func (p *ProductSlice) FindByID(id int) (pr internal.Product, err error) {
	// check if product exists
	var exists bool; var idx int
	for i, v := range p.products {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrProductRepositoryNotFound
		return
	}

	// set product
	pr = p.products[idx]

	return
}

// Save saves a product
func (p *ProductSlice) Save(pr *internal.Product) (err error) {
	// check constraints
	// - unique field
	for _, v := range p.products {
		if v.ProductCode == pr.ProductCode {
			err = internal.ErrProductRepositoryDuplicated
			return
		}
	}

	// increment lastID
	p.lastID++
	// set ID
	(*pr).ID = p.lastID
	// append to slice
	(*p).products = append((*p).products, *pr)

	return
}

// Update updates a product
func (p *ProductSlice) Update(pr *internal.Product) (err error) {
	// check if product exists
	var exists bool; var idx int
	for i, v := range p.products {
		if v.ID == pr.ID {
			exists = true
			idx = i
			break
		}

		// check constraints
		// - unique field
		if v.ProductCode == pr.ProductCode {
			err = internal.ErrProductRepositoryDuplicated
			return
		}
	}
	if !exists {
		err = internal.ErrProductRepositoryNotFound
		return
	}

	// update product
	(*p).products[idx] = *pr

	return
}

// Delete deletes a product
func (p *ProductSlice) Delete(id int) (err error) {
	// check if product exists
	var exists bool; var idx int
	for i, v := range p.products {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrProductRepositoryNotFound
		return
	}

	// delete product
	(*p).products = append((*p).products[:idx], (*p).products[idx+1:]...)

	return
}