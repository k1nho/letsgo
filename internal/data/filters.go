package data

import "github.com/k1nho/letsgo/internal/validator"

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

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0, "page", "must be greater than zero")
	v.Check(f.Page <= 10_000_000, "page", "must not exceed 10 million")
	v.Check(f.PageSize > 0, "page_size", "must be greater than zero")
	v.Check(f.PageSize <= 100, "page_size", "must not exceed 100")

	v.Check(validator.In(f.Sort, f.SortSafeList...), "sort", "invalid sort value")
}
