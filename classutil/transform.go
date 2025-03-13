package classutil

import "github.com/weaviate/weaviate/entities/models"

func PropertyToNestedProperty(in *models.Property) *models.NestedProperty {
	if in == nil {
		return nil
	}

	var out models.NestedProperty
	out.DataType = in.DataType
	out.Description = in.Description
	out.IndexFilterable = in.IndexFilterable
	out.IndexRangeFilters = in.IndexRangeFilters
	out.IndexSearchable = in.IndexSearchable
	out.Name = in.Name
	out.NestedProperties = in.NestedProperties
	out.Tokenization = in.Tokenization

	return &out
}
