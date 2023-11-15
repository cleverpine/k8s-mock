package dto

type TableResource struct {
	GenericResource

	ColumnDefinitions []ColumnDefinition `json:"columnDefinitions"`
	Rows              []RowDefinition    `json:"rows"`
}

func NewTableResource() *TableResource {
	return &TableResource{
		GenericResource: GenericResource{
			Kind:       "Table",
			APIVersion: "meta.k8s.io/v1",
		},
		ColumnDefinitions: []ColumnDefinition{},
		Rows:              []RowDefinition{},
	}
}

func (tr *TableResource) AddColumnDefinitions(cd ...ColumnDefinition) {
	for v := range cd {
		tr.ColumnDefinitions = append(tr.ColumnDefinitions, cd[v])
	}
}

func (tr *TableResource) AddRows(r ...RowDefinition) {
	for v := range r {
		tr.Rows = append(tr.Rows, r[v])
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
