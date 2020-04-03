package style

import "github.com/hooone/evening/pkg/managers/sqlstore/migrator"

func AddStyleMigration(mg *migrator.Migrator) {
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
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "create_at", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "update_at", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{},
	}

	mg.AddMigration("create style table", migrator.NewAddTableMigration(tbl))
}

func AddStyleSetMigration(mg *migrator.Migrator) {
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

	stylesetInsert1 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('RECT', 'XAXIS', '', true, false)"
	mg.AddMigration("insert rect chart style XAXIS", migrator.NewRawSqlMigration(stylesetInsert1))
	stylesetInsert2 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('RECT', 'Y1AXIS', 'LINE', true, true)"
	mg.AddMigration("insert rect chart style Y1AXIS", migrator.NewRawSqlMigration(stylesetInsert2))
	stylesetInsert3 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('RECT', 'Y2AXIS', 'LINE', true, true)"
	mg.AddMigration("insert rect chart style Y2AXIS", migrator.NewRawSqlMigration(stylesetInsert3))
	stylesetInsert4 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('RECT', 'Y1COLOR', '#fad248', false, false)"
	mg.AddMigration("insert rect chart style Y1COLOR", migrator.NewRawSqlMigration(stylesetInsert4))
	stylesetInsert5 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('RECT', 'Y2COLOR', 'red', false, false)"
	mg.AddMigration("insert rect chart style Y2COLOR", migrator.NewRawSqlMigration(stylesetInsert5))

	stylesetInsert6 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('POINT', 'XAXIS', '', true, true)"
	mg.AddMigration("insert point chart style XAXIS", migrator.NewRawSqlMigration(stylesetInsert6))
	stylesetInsert7 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('POINT', 'Y1AXIS', '#fad248', true, true)"
	mg.AddMigration("insert point chart style Y1AXIS", migrator.NewRawSqlMigration(stylesetInsert7))
	stylesetInsert8 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('POINT', 'Y2AXIS', 'red', true, true)"
	mg.AddMigration("insert point chart style Y2AXIS", migrator.NewRawSqlMigration(stylesetInsert8))

	stylesetInsert9 := "INSERT INTO `style_set` (`type`, `property`, `value`, `relation`, `must_number`) " +
		"VALUES ('STAT', 'FIELD', '', true, true)"
	mg.AddMigration("insert STAT Card Filed", migrator.NewRawSqlMigration(stylesetInsert9))

}
