package repository

import "github.com/usuario/repositorio/internal"

// NewSellerSlice returns a new instance of SellerSlice
func NewSellerSlice(sellers []internal.Seller, lastID int) *SellerSlice {
	// default values
	defaultSellers := make([]internal.Seller, 0)
	defaultLastID := 0
	if sellers != nil {
		defaultSellers = sellers
	}
	if lastID > 0 {
		defaultLastID = lastID
	}

	return &SellerSlice{
		sellers: defaultSellers,
		lastID:  defaultLastID,
	}
}

// SellerSlice is an implementation of SellerRepository
type SellerSlice struct {
	// sellers is a slice of Seller
	sellers []internal.Seller
	// lastID is the last used ID
	lastID int
}

// FindAll returns all sellers
func (s *SellerSlice) FindAll() (sl []internal.Seller, err error) {
	// make a copy of the slice
	sl = make([]internal.Seller, len(s.sellers))
	copy(sl, s.sellers)

	return
}

// FindByID returns a seller by ID
func (s *SellerSlice) FindByID(id int) (sl internal.Seller, err error) {
	// check if seller exists
	var exists bool; var idx int
	for i, v := range s.sellers {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrSellerRepositoryNotFound
		return
	}

	// set seller
	sl = s.sellers[idx]

	return
}

// Save saves a seller
func (s *SellerSlice) Save(sl *internal.Seller) (err error) {
	// check constraints
	// - unique field
	for _, v := range s.sellers {
		if v.CID == sl.CID {
			err = internal.ErrSellerRepositoryDuplicated
			return
		}
	}

	// increment lastID
	s.lastID++
	// set ID
	(*sl).ID = s.lastID
	// append to slice
	(*s).sellers = append((*s).sellers, *sl)

	return
}

// Update updates a seller
func (s *SellerSlice) Update(sl *internal.Seller) (err error) {
	// check if seller exists
	var exists bool; var idx int
	for i, v := range s.sellers {
		if v.ID == sl.ID {
			exists = true
			idx = i
			break
		}

		// check constraints
		// - unique field
		if v.CID == sl.CID {
			err = internal.ErrSellerRepositoryDuplicated
			return
		}
	}
	if !exists {
		err = internal.ErrSellerRepositoryNotFound
		return
	}

	// update seller
	(*s).sellers[idx] = *sl

	return
}

// Delete deletes a seller
func (s *SellerSlice) Delete(id int) (err error) {
	// check if seller exists
	var exists bool; var idx int
	for i, v := range s.sellers {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrSellerRepositoryNotFound
		return
	}

	// delete seller
	(*s).sellers = append((*s).sellers[:idx], (*s).sellers[idx+1:]...)

	return
}