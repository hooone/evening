package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addStyleMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "style",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "card_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "field_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "type", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "property", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "value", Type: migrator.DB_NVarchar, Length: 1024, Nullable: false},
			{Name: "relation", Type: migrator.DB_Bit, Nullable: false},
			{Name: "must_number", Type: migrator.DB_Bit, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create style table", migrator.NewAddTableMigration(tbl))
}

func addStyleSetMigration(mg *migrator.Migrator) {
	var tbl = migrator.Table{
		Name: "style_set",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "type", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "property", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "value", Type: migrator.DB_NVarchar, Length: 1024, Nullable: false},
			{Name: "relation", Type: migrator.DB_Bit, Nullable: false},
			{Name: "must_number", Type: migrator.DB_Bit, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create style_set table", migrator.NewAddTableMigration(tbl))

	// const rawSQL = `
	// INSERT INTO style_set
	//  	(
	// 		[type],
	// 		[property],
	//  		[value],
	// 		[relation],
	//  	)
	//  	VALUES
	//  	('RECT','X2AXIS', '',true)
	//  	`

	// mg.AddMigration("save default style set in style_set table", migrator.NewRawSqlMigration(rawSQL))

}
