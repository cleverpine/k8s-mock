package dto

type TableResource struct {
	Resource

	ColumnDefinitions []ColumnDefinition `json:"columnDefinitions"`
	Rows              []RowDefinition    `json:"rows"`
}

func NewTableResource() *TableResource {
	return &TableResource{
		Resource: Resource{
			"kind":       "Table",
			"apiVersion": "meta.k8s.io/v1",
		},
		ColumnDefinitions: []ColumnDefinition{},
		Rows:              []RowDefinition{},
	}
}

func (tr *TableResource) AddColumnDefinitions(cd ...ColumnDefinition) {
	tr.ColumnDefinitions = append(tr.ColumnDefinitions, cd...)
}

func (tr *TableResource) AddRows(r ...RowDefinition) {
	tr.Rows = append(tr.Rows, r...)
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
