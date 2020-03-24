package main

// Model manages all Entities within the system
// A model has table fields and needs to interact with the persistence layer
type Model struct {
	TableReader
}

// TableReader setups a table for managing a Model's data
type TableReader interface {
	getTableName() string
	getTableData() interface{}
}

// ModelDTO - the data transfer object for a model
type ModelDTO interface {
	getID() string
}

// NewModel - Factory for generating models
func NewModel(t TableReader) Model {
	return Model{t}
}

// CreateRecord is used to insert a new row into a table
func (m Model) CreateRecord(d ModelDTO) (ModelDTO, error) {
	rows := m.getTableData().([]ModelDTO)
	rows = append(rows, d)

	return d, nil
}

func (m Model) setTableData(d interface{}) {
}
