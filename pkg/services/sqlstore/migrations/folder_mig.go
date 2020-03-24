package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addFolderMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "folder",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "is_folder", Type: migrator.DB_Bit, Length: 1, Nullable: false},
			{Name: "seq", Type: migrator.DB_Int, Nullable: false},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create folder table", migrator.NewAddTableMigration(tbl))
}
