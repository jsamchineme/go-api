package main

// Model manages all Entities within the system
// A model has table fields and needs to interact with the persistence layer
type Model struct {
	TableReader
}

// TableReader setups a table for managing a Model's data
type TableReader interface {
	getTableName() string
	getTableData() []ModelDTO
	setTableData([]ModelDTO)
}

// ModelDTO - the data transfer object for a model
type ModelDTO struct {
	IsDTO
}

// IsDTO marks an object as DTO
type IsDTO interface {
	getID() string
}

// NewModel - factory for generating models
func NewModel(t TableReader) Model {
	return Model{t}
}

// MakeDTO - factory to generate a ModelDTO
func MakeDTO(t IsDTO) ModelDTO {
	return ModelDTO{t}
}

// CreateRecord is used to insert a new row into a table
func (m Model) CreateRecord(d ModelDTO) (ModelDTO, error) {
	rows := m.getTableData()
	rows = append(rows, d)
	m.setTableData(rows)

	return d, nil
}
