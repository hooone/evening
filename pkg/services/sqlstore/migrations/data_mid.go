package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addTestDataMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "test_data",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true},
			{Name: "card_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "value", Type: migrator.DB_NVarchar, Length: 16000, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create card test data", migrator.NewAddTableMigration(tbl))
}
