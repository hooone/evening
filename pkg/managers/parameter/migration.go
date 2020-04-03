package parameter

import "github.com/hooone/evening/pkg/managers/sqlstore/migrator"

func AddParameterMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "parameter",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "action_Id", Type: migrator.DB_BigInt},
			{Name: "field_id", Type: migrator.DB_BigInt},
			{Name: "is_visible", Type: migrator.DB_Bit, Nullable: false},
			{Name: "is_editable", Type: migrator.DB_Bit, Nullable: false},
			{Name: "default", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "create_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "update_at", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create parameter table", migrator.NewAddTableMigration(tbl))
}
