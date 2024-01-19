package repository

import "github.com/usuario/repositorio/internal"

// NewBuyerSlice returns a new instance of BuyerSlice
func NewBuyerSlice(buyers []internal.Buyer, lastID int) *BuyerSlice {
	// default values
	defaultBuyers := make([]internal.Buyer, 0)
	defaultLastID := 0
	if buyers != nil {
		defaultBuyers = buyers
	}
	if lastID > 0 {
		defaultLastID = lastID
	}

	return &BuyerSlice{
		buyers: defaultBuyers,
		lastID:  defaultLastID,
	}
}

// BuyerSlice is an implementation of BuyerRepository
type BuyerSlice struct {
	// buyers is a slice of Buyer
	buyers []internal.Buyer
	// lastID is the last used ID
	lastID int
}

// FindAll returns all buyers
func (b *BuyerSlice) FindAll() (by []internal.Buyer, err error) {
	// make a copy of the slice
	by = make([]internal.Buyer, len(b.buyers))
	copy(by, b.buyers)

	return
}

// FindByID returns a buyer by ID
func (b *BuyerSlice) FindByID(id int) (by internal.Buyer, err error) {
	// check if buyer exists
	var exists bool; var idx int
	for i, v := range b.buyers {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrBuyerRepositoryNotFound
		return
	}

	// set buyer
	by = b.buyers[idx]

	return
}

// Save saves a buyer
func (b *BuyerSlice) Save(by *internal.Buyer) (err error) {
	// check constraints
	// - unique field
	for _, v := range b.buyers {
		if v.CardNumberID == by.CardNumberID {
			err = internal.ErrBuyerRepositoryDuplicated
			return
		}
	}
	
	// increment lastID
	b.lastID++
	// set ID
	(*by).ID = b.lastID
	// append to slice
	(*b).buyers = append((*b).buyers, *by)

	return
}

// Update updates a buyer
func (b *BuyerSlice) Update(by *internal.Buyer) (err error) {
	// check if buyer exists
	var exists bool; var idx int
	for i, v := range b.buyers {
		if v.ID == by.ID {
			exists = true
			idx = i
			break
		}

		// check constraints
		// - unique field
		if v.CardNumberID == by.CardNumberID {
			err = internal.ErrBuyerRepositoryDuplicated
			return
		}
	}
	if !exists {
		err = internal.ErrBuyerRepositoryNotFound
		return
	}

	// update buyer
	(*b).buyers[idx] = *by

	return
}

// Delete deletes a buyer
func (b *BuyerSlice) Delete(id int) (err error) {
	// check if buyer exists
	var exists bool; var idx int
	for i, v := range b.buyers {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrBuyerRepositoryNotFound
		return
	}

	// delete buyer
	(*b).buyers = append((*b).buyers[:idx], (*b).buyers[idx+1:]...)

	return
}