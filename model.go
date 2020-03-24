package main

import (
	"fmt"
)

// Model manages all Entities within the system
// A model has table fields and needs to interact with the persistence layer
type Model interface {
	TableReader
}

// TableReader setups a table for managing a Model's data
type TableReader interface {
	getTableName() string
}

// ModelDTO - the data transfer object for a model
type ModelDTO interface {
	getID() string
}

// CreateRecord is used to insert a new row into a table
func CreateRecord(model Model, data ModelDTO) (ModelDTO, error) {
	fmt.Println(model.getTableName())

	model.getTableData()
	return nil, nil
}
