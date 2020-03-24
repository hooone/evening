package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addPageMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "page",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "folder_id", Type: migrator.DB_BigInt},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "seq", Type: migrator.DB_Int, Length: 4, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create page table", migrator.NewAddTableMigration(tbl))
}
