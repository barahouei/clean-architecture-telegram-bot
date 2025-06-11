package models

type Database uint

const (
	_ Database = iota
	PostgreSQL
	MySQL
	MongoDB
)

func (d Database) String() string {
	databases := []string{"", "postgres", "mysql", "mongodb"}

	return databases[d]
}
