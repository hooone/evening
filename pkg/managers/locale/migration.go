package locale

import "github.com/hooone/evening/pkg/managers/sqlstore/migrator"

func AddLocaleMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "locale",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "v_id", Type: migrator.DB_BigInt},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "type", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "language", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "text", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "create_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "update_at", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create locale table", migrator.NewAddTableMigration(tbl))
}
