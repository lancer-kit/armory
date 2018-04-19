package db

import (
	sq "github.com/Masterminds/squirrel"
)

type Table struct {
	Name  string
	Alias string

	DB       sq.BaseRunner
	QBuilder sq.SelectBuilder
	Page     *PageQuery
}

func (t Table) AliasedName() string {
	return t.Name + " " + t.Alias
}

func (t *Table) SetPage(pq *PageQuery) {
	t.Page = pq
}

func (t *Table) ApplyPage(orderColumn string) {
	if t.Page != nil {
		t.QBuilder = t.Page.Apply(t.QBuilder, orderColumn)
		return
	}

	t.QBuilder = t.QBuilder.OrderBy(orderColumn)
}

func (t *Table) WithCount() {
	t.QBuilder = t.QBuilder.Column("count(*) OVER() AS row_count")
}
