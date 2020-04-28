package main

import (
	"fmt"
	"strings"

	"github.com/lancer-kit/armory/db"
)

func main() {

	table := db.NewTable("test", "t", "*")

	query := table.QBuilder.
		Where("name", "john").
		OrderBy("ds ASC").
		Limit(20).
		Offset(12)

	rawSQL, _, err := query.ToSql()

	fmt.Println(err)
	fmt.Println(rawSQL)
	fmt.Println(strings.Replace(rawSQL, table.Columns, "count(*) as count", 1))
}
