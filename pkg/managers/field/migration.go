package field

import "github.com/hooone/evening/pkg/managers/sqlstore/migrator"

func AddFieldMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "field",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "card_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "seq", Type: migrator.DB_Int, Nullable: false},
			{Name: "is_visible", Type: migrator.DB_Bit, Nullable: false},
			{Name: "type", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "default", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "filter", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "create_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "update_at", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create field table", migrator.NewAddTableMigration(tbl))
}
