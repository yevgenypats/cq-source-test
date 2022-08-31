package tables

import (
	"context"

	"github.com/cloudquery/plugin-sdk/schema"
)

func TestSomeTable() *schema.Table {
	return &schema.Table{
		Name:        "test_some_table",
		Description: "Test description",
		Resolver:    fetchBigqueryDatasets,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"project_id", "id"}},
		Columns: []schema.Column{
			{
				Name:        "column1",
				Description: "Test Column 1",
				Type:        schema.TypeString,
			},
			{
				Name:        "column2",
				Description: "Test Column 2",
				Type:        schema.TypeBigInt,
			},
		},
	}
}

func fetchBigqueryDatasets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	res <- map[string]interface{}{
		"column1": "test_project_id",
		"column2": "test_id",
	}
	return nil
}
