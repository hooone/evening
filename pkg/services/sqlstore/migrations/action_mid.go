package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addViewActionMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "view_action",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "card_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "type", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "seq", Type: migrator.DB_Int, Nullable: false},
			{Name: "double_check", Type: migrator.DB_Bit, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create view_action table", migrator.NewAddTableMigration(tbl))
}
