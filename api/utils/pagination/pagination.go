package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/samigmuseyibli/gin-gonic-blog/api/models"
)

func GeneratePaginationRequest(context *gin.Context) *models.Pagination {
	// default limit, page & sort parameter
	limit := 100
	page := 0
	sort := "created_at desc"

	var searchs []models.Search

	query := context.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}

		/*
			// check if query parameter key contains dot
			if strings.Contains(key, ".") {
				// split query parameter key by dot
				searchKeys := strings.Split(key, ".")

				// create search object
				search := models.Search{Column: searchKeys[0], Action: searchKeys[1], Query: queryValue}

				// add search object to searchs array
				searchs = append(searchs, search)
			}
		*/
	}

	return &models.Pagination{Limit: limit, Page: page, Sort: sort, Searchs: searchs}
}
