package card

import "github.com/hooone/evening/pkg/managers/sqlstore/migrator"

func AddCardMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "card",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "page_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "style", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "seq", Type: migrator.DB_Int, Nullable: false},
			{Name: "pos", Type: migrator.DB_Int, Nullable: false},
			{Name: "width", Type: migrator.DB_Int, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "create_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "update_at", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create card table", migrator.NewAddTableMigration(tbl))
}
