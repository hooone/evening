package migrations

import "github.com/hooone/evening/pkg/services/sqlstore/migrator"

//AddMigrations initialization database
func AddMigrations(mg *migrator.Migrator) {
	addMigrationLogMigrations(mg)
	addLocaleMigration(mg)
	addFolderMigration(mg)
	addPageMigration(mg)
	addCardMigration(mg)
	addFieldMigration(mg)
	addViewActionMigration(mg)
	addParameterMigration(mg)
	addTestDataMigration(mg)
	addStyleMigration(mg)
	addStyleSetMigration(mg)
	addUserTokenMigrations(mg)
	addUserMigration(mg)
}

func addMigrationLogMigrations(mg *migrator.Migrator) {
	migrationLogV1 := migrator.Table{
		Name: "migration_log",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "migration_id", Type: migrator.DB_NVarchar, Length: 255},
			{Name: "sql", Type: migrator.DB_Text},
			{Name: "success", Type: migrator.DB_Bool},
			{Name: "error", Type: migrator.DB_Text},
			{Name: "timestamp", Type: migrator.DB_DateTime},
		},
	}

	mg.AddMigration("create migration_log table", migrator.NewAddTableMigration(migrationLogV1))
}
