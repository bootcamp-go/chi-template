package repository

import "github.com/usuario/repositorio/internal"

// NewEmployeeSlice returns a new instance of EmployeeSlice
func NewEmployeeSlice(employees []internal.Employee, lastID int) *EmployeeSlice {
	// default values
	defaultEmployees := make([]internal.Employee, 0)
	defaultLastID := 0
	if employees != nil {
		defaultEmployees = employees
	}
	if lastID > 0 {
		defaultLastID = lastID
	}

	return &EmployeeSlice{
		employees: defaultEmployees,
		lastID:  defaultLastID,
	}
}

// EmployeeSlice is an implementation of EmployeeRepository
type EmployeeSlice struct {
	// employees is a slice of Employee
	employees []internal.Employee
	// lastID is the last used ID
	lastID int
}

// FindAll returns all employees
func (e *EmployeeSlice) FindAll() (em []internal.Employee, err error) {
	// make a copy of the slice
	em = make([]internal.Employee, len(e.employees))
	copy(em, e.employees)

	return
}

// FindByID returns a employee by ID
func (e *EmployeeSlice) FindByID(id int) (em internal.Employee, err error) {
	// check if employee exists
	var exists bool; var idx int
	for i, v := range e.employees {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrEmployeeRepositoryNotFound
		return
	}

	// set employee
	em = e.employees[idx]

	return
}

// Save saves a employee
func (e *EmployeeSlice) Save(em *internal.Employee) (err error) {
	// check constraints
	// - unique field
	for _, v := range e.employees {
		if v.CardNumberID == em.CardNumberID {
			err = internal.ErrEmployeeRepositoryDuplicated
			return
		}
	}
	
	// increment lastID
	e.lastID++
	// set ID
	(*em).ID = e.lastID
	// append to slice
	(*e).employees = append((*e).employees, *em)

	return
}

// Update updates a employee
func (e *EmployeeSlice) Update(em *internal.Employee) (err error) {
	// check if employee exists
	var exists bool; var idx int
	for i, v := range e.employees {
		if v.ID == em.ID {
			exists = true
			idx = i
			break
		}

		// check constraints
		// - unique field
		if v.CardNumberID == em.CardNumberID {
			err = internal.ErrEmployeeRepositoryDuplicated
			return
		}
	}
	if !exists {
		err = internal.ErrEmployeeRepositoryNotFound
		return
	}

	// update employee
	(*e).employees[idx] = *em

	return
}

// Delete deletes a employee
func (e *EmployeeSlice) Delete(id int) (err error) {
	// check if employee exists
	var exists bool; var idx int
	for i, v := range e.employees {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrEmployeeRepositoryNotFound
		return
	}

	// delete employee
	(*e).employees = append((*e).employees[:idx], (*e).employees[idx+1:]...)

	return
}