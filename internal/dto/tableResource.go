package dto

type TableResource GenericResource

func NewTableResource() *TableResource {
	return &TableResource{
		"kind":              "Table",
		"apiVersion":        "meta.k8s.io/v1",
		"columnDefinitions": []GenericResource{},
		"rows":              []GenericResource{},
	}
}

func (tr *TableResource) AddColumnDefinition(colName string, colType string, colFormat string, colPriority int) {
	columnDefinitions, ok := (*tr)["columnDefinitions"].([]GenericResource)
	if !ok {
		columnDefinitions = []GenericResource{}
	}
	columnDefinitions = append(columnDefinitions, GenericResource{
		"name":     colName,
		"type":     colType,
		"format":   colFormat,
		"priority": colPriority,
	})

	(*tr)["columnDefinitions"] = columnDefinitions
}

func (tr *TableResource) AddRow(cells []string, object *GenericResource) {
	rows, ok := (*tr)["rows"].([]GenericResource)
	if !ok {
		rows = []GenericResource{}
	}

	rows = append(rows, GenericResource{
		"cells":  cells,
		"object": object,
	})

	(*tr)["rows"] = rows
}
