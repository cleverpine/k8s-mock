package dto

type TableResource Resource

func NewTableResource() TableResource {
	return TableResource{
		"kind":              "Table",
		"columnDefinitions": []ColumnDefinition{},
		"rows":              []RowDefinition{},
	}
}

func (tr *TableResource) AddColumnDefinitions(newCds ...ColumnDefinition) {
	if cds, ok := (*tr)["columnDefinitions"]; ok {
		if cdsTyped, ok := cds.([]ColumnDefinition); ok {
			(*tr)["columnDefinitions"] = append(cdsTyped, newCds...)
		}
	}
}

func (tr *TableResource) AddRows(newRows ...RowDefinition) {
	if rows, ok := (*tr)["rows"]; ok {
		if rowsTyped, ok := rows.([]RowDefinition); ok {
			(*tr)["rows"] = append(rowsTyped, newRows...)
		}
	}
}

type ColumnDefinition struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Format   string `json:"format"`
	Priority int    `json:"priority"`
}

type RowDefinition struct {
	Cells  []string `json:"cells"`
	Object any      `json:"object"`
}
