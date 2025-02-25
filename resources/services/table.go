package services

import (
	"context"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

type SampleStruct struct {
	NullableField    *string
	NonNullableField string
}

func SampleTable() *schema.Table {
	return &schema.Table{
		Name:      "cloudquery-struct-nullable-fields_sample_table",
		Resolver:  fetchSampleTable,
		Transform: transformers.TransformWithStruct(SampleStruct{}),
	}
}

func fetchSampleTable(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	nullable := "nullable"

	res <- SampleStruct{
		NullableField:    &nullable,
		NonNullableField: "non-nullable",
	}

	return nil
}
