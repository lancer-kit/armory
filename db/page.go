package db

import (
	"fmt"
	"net/url"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

const (
	DefaultPageSize uint64 = 20
	MaxPageSize     uint64 = 500

	OrderAscending  = "asc"
	OrderDescending = "desc"
)

var (
	ErrInvalidOrder = func(val string) error {
		return fmt.Errorf("order(%s): accept only %s|%s",
			val, OrderAscending, OrderDescending)
	}
	ErrTooBigPage = func(val uint64) error {
		return fmt.Errorf("pageSize(%d): shoud be less or equal %d", val, MaxPageSize)
	}
)

// PageQuery is the structure for building query with pagination.
type PageQuery struct {
	Order    string `json:"order" schema:"order"`
	Page     uint64 `json:"page" schema:"page"`
	PageSize uint64 `json:"pageSize" schema:"pageSize"`
	OrderBy  string `json:"orderBy" schema:"orderBy"`
}

// ParsePageQuery extracts `PageQuery` from the url Query Values.
func ParsePageQuery(values url.Values) (pq PageQuery, err error) {
	err = pq.FromRQuery(values)
	return
}

// FromRQuery extracts `PageQuery` from the url Query Values and validates.
func (pq *PageQuery) FromRQuery(query url.Values) error {
	urlValuesEncoder := schema.NewDecoder()
	urlValuesEncoder.IgnoreUnknownKeys(true)
	err := urlValuesEncoder.Decode(pq, query)
	if err != nil {
		return errors.Wrap(err, "failed to decode PageQuery from url.Values")
	}

	return pq.Validate()
}

// Validate checks is correct values and
// sets default values if `PageQuery` empty.
// WARN: the receiver MUST be a pointer so that the default values works
func (pq *PageQuery) Validate() error {
	switch strings.ToLower(pq.Order) {
	case "":
		pq.Order = OrderAscending
	case OrderAscending, OrderDescending:
		break
	default:
		return ErrInvalidOrder(pq.Order)
	}

	if pq.Page == 0 {
		pq.Page = 1
	}

	if pq.PageSize == 0 {
		pq.PageSize = DefaultPageSize
	}

	if pq.PageSize > MaxPageSize {
		return ErrTooBigPage(pq.PageSize)
	}

	return nil
}

// Offset calculates select offset.
func (pq *PageQuery) Offset() uint64 {
	return (pq.Page - 1) * pq.PageSize
}

// DEPRECATED
// use ApplyByOrderColumn instead
// Apply sets limit and ordering params to SelectBuilder.
func (pq *PageQuery) Apply(query sq.SelectBuilder, orderColumn string) sq.SelectBuilder {
	query = query.Limit(pq.PageSize).Offset(pq.Offset())
	if pq.Order != "" && orderColumn != "" {
		query = query.OrderBy(orderColumn + " " + pq.Order)
	}

	return query
}

// ApplyByOrderColumn sets limit and ordering params to SelectBuilder.
func (pq *PageQuery) ApplyByOrderColumn(query sq.SelectBuilder) sq.SelectBuilder {
	query = query.Limit(pq.PageSize).Offset(pq.Offset())
	if pq.Order != "" && pq.OrderBy != "" {
		query = query.OrderBy(pq.OrderBy + " " + pq.Order)
	}

	return query
}
