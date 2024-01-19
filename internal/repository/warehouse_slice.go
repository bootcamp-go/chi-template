package repository

import "github.com/usuario/repositorio/internal"

// NewWarehouseSlice returns a new instance of WarehouseSlice
func NewWarehouseSlice(warehouses []internal.Warehouse, lastID int) *WarehouseSlice {
	// default values
	defaultWarehouses := make([]internal.Warehouse, 0)
	defaultLastID := 0
	if warehouses != nil {
		defaultWarehouses = warehouses
	}
	if lastID > 0 {
		defaultLastID = lastID
	}

	return &WarehouseSlice{
		warehouses: defaultWarehouses,
		lastID:  defaultLastID,
	}
}

// WarehouseSlice is an implementation of WarehouseRepository
type WarehouseSlice struct {
	// warehouses is a slice of Warehouse
	warehouses []internal.Warehouse
	// lastID is the last used ID
	lastID int
}

// FindAll returns all warehouses
func (w *WarehouseSlice) FindAll() (wh []internal.Warehouse, err error) {
	// make a copy of the slice
	wh = make([]internal.Warehouse, len(w.warehouses))
	copy(wh, w.warehouses)

	return
}

// FindByID returns a warehouse by ID
func (w *WarehouseSlice) FindByID(id int) (wh internal.Warehouse, err error) {
	// check if warehouse exists
	var exists bool; var idx int
	for i, v := range w.warehouses {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrWarehouseRepositoryNotFound
		return
	}

	// set warehouse
	wh = w.warehouses[idx]

	return
}

// Save saves a warehouse
func (w *WarehouseSlice) Save(wh *internal.Warehouse) (err error) {
	// check constraints
	// - unique field
	for _, v := range w.warehouses {
		if v.WarehouseCode == wh.WarehouseCode {
			err = internal.ErrWarehouseRepositoryDuplicated
			return
		}
	}

	// increment lastID
	w.lastID++
	// set ID
	(*wh).ID = w.lastID
	// append to slice
	(*w).warehouses = append((*w).warehouses, *wh)

	return
}

// Update updates a warehouse
func (w *WarehouseSlice) Update(wh *internal.Warehouse) (err error) {
	// check if warehouse exists
	var exists bool; var idx int
	for i, v := range w.warehouses {
		if v.ID == wh.ID {
			exists = true
			idx = i
			break
		}

		// check constraints
		// - unique field
		if v.WarehouseCode == wh.WarehouseCode {
			err = internal.ErrWarehouseRepositoryDuplicated
			return
		}
	}
	if !exists {
		err = internal.ErrWarehouseRepositoryNotFound
		return
	}

	// update warehouse
	(*w).warehouses[idx] = *wh

	return
}

// Delete deletes a warehouse
func (w *WarehouseSlice) Delete(id int) (err error) {
	// check if warehouse exists
	var exists bool; var idx int
	for i, v := range w.warehouses {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrWarehouseRepositoryNotFound
		return
	}

	// delete warehouse
	(*w).warehouses = append((*w).warehouses[:idx], (*w).warehouses[idx+1:]...)

	return
}