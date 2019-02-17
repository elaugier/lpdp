package db

import (
	"math"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetSectionQuery ...
func GetSectionQuery(params graphql.ResolveParams) (interface{}, error) {
	var sections []models.Section
	var count int
	page := params.Args["page"].(int)
	itemsPerPage := params.Args["itemsPerPage"].(int)
	Inst.c.Model(&models.Section{}).Count(&count)
	t := float64(count) / float64(itemsPerPage)
	o := math.RoundToEven(t)
	offset := int(o) * page
	Inst.c.Offset(offset).Limit(itemsPerPage).Find(&sections)
	return sections, nil
}
