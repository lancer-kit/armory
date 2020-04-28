package db

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
)

// Table is the basis for implementing
// Querier for some model or table.
type Table struct {
	Name    string
	Alias   string
	Columns string

	DB        sq.BaseRunner
	QBuilder  sq.SelectBuilder
	GQBuilder sq.SelectBuilder
	IQBuilder sq.InsertBuilder
	UQBuilder sq.UpdateBuilder
	DQBuilder sq.DeleteBuilder
	Page      *PageQuery
}

func NewTable(table, alias, columns string) Table {
	return Table{
		Name:    table,
		Alias:   alias,
		Columns: columns,

		QBuilder:  sq.Select(columns).From(table),
		GQBuilder: sq.Select(columns).From(table),
		IQBuilder: sq.Insert(table),
		UQBuilder: sq.Update(table),
		DQBuilder: sq.Delete(table),
	}
}

// AliasedName returns table name with the alias postfix.
func (t Table) AliasedName() string {
	return t.Name + " " + t.Alias
}

// SetPage is a setter for Page field.
func (t *Table) SetPage(pq *PageQuery) {
	t.Page = pq
}

// ApplyPage adds limit/offset and/or order to the queryBuilder.
func (t *Table) ApplyPage(orderColumn string) {
	if t.Page != nil {
		t.QBuilder = t.Page.Apply(t.QBuilder, orderColumn)
		return
	}

	t.QBuilder = t.QBuilder.OrderBy(orderColumn)
}

func (t *Table) CountQuery() (string, []interface{}, error) {
	rawSQL, args, err := t.QBuilder.RemoveLimit().ToSql()
	if err != nil {
		return "", nil, err
	}

	countQuery := strings.Replace(rawSQL, t.Columns, "count(1) as count", 1)
	return countQuery, args, nil
}

func (t *Table) GetCount(sqlConn *SQLConn) (int64, error) {
	rawSQL, args, err := t.QBuilder.RemoveLimit().ToSql()
	if err != nil {
		return 0, err
	}

	countQuery := strings.Replace(rawSQL, t.Columns, "count(1) as count", 1)

	dest := new(Count)
	err = sqlConn.GetRaw(dest, countQuery, args...)
	return dest.Count, err
}

func (t *Table) SelectWithCount(sqlConn *SQLConn, dest interface{}, orderColumn string, query *PageQuery) (int64, error) {
	count, err := t.GetCount(sqlConn)
	if err != nil {
		return 0, errors.Wrap(err, "can not GET count")
	}

	if query != nil {
		t.Page = query
	}

	t.ApplyPage(orderColumn)
	err = sqlConn.Select(t.QBuilder, dest)
	return count, err
}

// DEPRECATED
// Do not use this method for counting,
// because it very slow and does count for each row,
// WithCount adds a column with the total number of records.
// ATTENTION! The model must have a destination for this `row_count` column.
func (t *Table) WithCount() {
	t.QBuilder = t.QBuilder.Column("count(*) OVER() AS row_count")
}

type Count struct {
	Count int64 `db:"count"`
}
