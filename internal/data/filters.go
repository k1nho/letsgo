package data

import (
	"strings"

	"github.com/k1nho/letsgo/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

/* Validator Contraints
   -- Page: must be between 1 and 10,000,000
   -- PageSize: must be between 1 and 100
   -- Sort: must be accept a valid sort paramater (id, title, year, runtime) and its descending variants (-)
*/

// ValidateFilters: Validates that page, page size and sort are under the constraints defined
func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must not exceed 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must not exceed 100")

	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}

// SortColumn: Checks that the sort string given is included in the safelist, if it is then it returns the sort string without the prefix, otherwise it panics
func (f Filters) SortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}
	panic("unsafe sort parameter: " + f.Sort)
}

// sortDirection: Returns DESC if the sort string contains '-', otherwise it returns ASC
func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}
