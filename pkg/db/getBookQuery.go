package db

import (
	"math"

	"github.com/elaugier/lpdp/pkg/models"
	"github.com/graphql-go/graphql"
)

//GetBookQuery ...
func GetBookQuery(params graphql.ResolveParams) (interface{}, error) {
	var books []models.Book
	var count int
	page := params.Args["page"].(int)
	logger.Printf("page requested : %d", page)
	itemsPerPage := params.Args["itemsPerPage"].(int)
	logger.Printf("number of items per page requested : %d", itemsPerPage)
	Inst.c.Model(&models.Book{}).Count(&count)
	pages := math.RoundToEven(float64(count) / float64(itemsPerPage))
	logger.Printf("computed total pages : %d", int(pages))
	offset := (page - 1) * itemsPerPage
	Inst.c.Offset(offset).Limit(itemsPerPage).Find(&books)
	return books, nil
}
