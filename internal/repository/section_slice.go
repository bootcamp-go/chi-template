package repository

import "github.com/usuario/repositorio/internal"

// NewSectionSlice returns a new instance of SectionSlice
func NewSectionSlice(sections []internal.Section, lastID int) *SectionSlice {
	// default values
	defaultSections := make([]internal.Section, 0)
	defaultLastID := 0
	if sections != nil {
		defaultSections = sections
	}
	if lastID > 0 {
		defaultLastID = lastID
	}

	return &SectionSlice{
		sections: defaultSections,
		lastID:  defaultLastID,
	}
}

// SectionSlice is an implementation of SectionRepository
type SectionSlice struct {
	// sections is a slice of Section
	sections []internal.Section
	// lastID is the last used ID
	lastID int
}

// FindAll returns all sections
func (s *SectionSlice) FindAll() (st []internal.Section, err error) {
	// make a copy of the slice
	st = make([]internal.Section, len(s.sections))
	copy(st, s.sections)

	return
}

// FindByID returns a section by ID
func (s *SectionSlice) FindByID(id int) (st internal.Section, err error) {
	// check if section exists
	var exists bool; var idx int
	for i, v := range s.sections {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrSectionRepositoryNotFound
		return
	}

	// set section
	st = s.sections[idx]

	return
}

// Save saves a section
func (s *SectionSlice) Save(st *internal.Section) (err error) {
	// check constraints
	// - unique field
	for _, v := range s.sections {
		if v.SectionNumber == st.SectionNumber {
			err = internal.ErrSectionRepositoryDuplicated
			return
		}
	}

	// increment lastID
	s.lastID++
	// set ID
	(*st).ID = s.lastID
	// append to slice
	(*s).sections = append((*s).sections, *st)

	return
}

// Update updates a section
func (s *SectionSlice) Update(st *internal.Section) (err error) {
	// check if section exists
	var exists bool; var idx int
	for i, v := range s.sections {
		if v.ID == st.ID {
			exists = true
			idx = i
			break
		}

		// check constraints
		// - unique field
		if v.SectionNumber == st.SectionNumber {
			err = internal.ErrSectionRepositoryDuplicated
			return
		}
	}
	if !exists {
		err = internal.ErrSectionRepositoryNotFound
		return
	}

	// update section
	(*s).sections[idx] = *st

	return
}

// Delete deletes a section
func (s *SectionSlice) Delete(id int) (err error) {
	// check if section exists
	var exists bool; var idx int
	for i, v := range s.sections {
		if v.ID == id {
			exists = true
			idx = i
			break
		}
	}
	if !exists {
		err = internal.ErrSectionRepositoryNotFound
		return
	}

	// delete section
	(*s).sections = append((*s).sections[:idx], (*s).sections[idx+1:]...)

	return
}