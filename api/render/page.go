package render

import (
	"math"
	"net/http"
)

// Page is a standard response structure to render paginated list.
type Page struct {
	// Page number of current page.
	Page uint64 `json:"page"`
	// PageSize is a number of records per page.
	PageSize uint64 `json:"pageSize"`
	// Order is ordering direction: asc or desc.
	Order string `json:"order"`
	// Total is total count of pages.
	Total int64 `json:"total"`
	// Total is total count of rows.
	TotalRows int64 `json:"total_rows"`
	// Records is an array of rows.
	Records interface{} `json:"records"`
}

// Render writes page with http.StatusOK.
func (page *Page) Render(w http.ResponseWriter) {
	WriteJSON(w, http.StatusOK, page)
}

// SetTotal fills total count properly.
func (page *Page) SetTotal(rowCount, pageSize uint64) {
	page.Total = int64(math.Ceil(float64(rowCount) / float64(pageSize)))
	page.TotalRows = int64(rowCount)
}

// DEPRECATED
// BaseRow is a helper struct.
type BaseRow struct {
	RowCount int64 `db:"row_count" json:"-"`
}
