package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addLocaleMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "locale",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "v_id", Type: migrator.DB_BigInt},
			{Name: "type", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "language", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "text", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create locale table", migrator.NewAddTableMigration(tbl))
}
