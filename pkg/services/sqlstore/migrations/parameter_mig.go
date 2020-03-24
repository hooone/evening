package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addParameterMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "parameter",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "action_Id", Type: migrator.DB_BigInt},
			{Name: "field_id", Type: migrator.DB_BigInt},
			{Name: "is_visible", Type: migrator.DB_Bit, Nullable: false},
			{Name: "is_editable", Type: migrator.DB_Bit, Nullable: false},
			{Name: "default", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create parameter table", migrator.NewAddTableMigration(tbl))
}
