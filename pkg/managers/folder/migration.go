package folder

import (
	"github.com/hooone/evening/pkg/managers/sqlstore/migrator"
)

func AddFolderMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "folder",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "is_folder", Type: migrator.DB_Bit, Length: 1, Nullable: false},
			{Name: "seq", Type: migrator.DB_Int, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "create_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "update_at", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create folder table", migrator.NewAddTableMigration(tbl))
}
