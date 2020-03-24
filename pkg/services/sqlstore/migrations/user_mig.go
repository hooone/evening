package migrations

import (
	"github.com/hooone/evening/pkg/services/sqlstore/migrator"
)

func addUserMigration(mg *migrator.Migrator) {
	userV1 := migrator.Table{
		Name: "user",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "version", Type: migrator.DB_Int, Nullable: false},
			{Name: "login", Type: migrator.DB_NVarchar, Length: 190, Nullable: false},
			{Name: "email", Type: migrator.DB_NVarchar, Length: 190, Nullable: false},
			{Name: "name", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "password", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "salt", Type: migrator.DB_NVarchar, Length: 50, Nullable: true},
			{Name: "rands", Type: migrator.DB_NVarchar, Length: 50, Nullable: true},
			{Name: "company", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "org_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "is_admin", Type: migrator.DB_Bool, Nullable: false},
			{Name: "email_verified", Type: migrator.DB_Bool, Nullable: true},
			{Name: "theme", Type: migrator.DB_NVarchar, Length: 255, Nullable: true},
			{Name: "created", Type: migrator.DB_DateTime, Nullable: false},
			{Name: "updated", Type: migrator.DB_DateTime, Nullable: false},
		},
		Indices: []*migrator.Index{
			{Cols: []string{"login"}, Type: migrator.UniqueIndex},
			{Cols: []string{"email"}, Type: migrator.UniqueIndex},
		},
	}

	// create table
	mg.AddMigration("create user table", migrator.NewAddTableMigration(userV1))
	// add indices
	mg.AddMigration("add unique index user.login", migrator.NewAddIndexMigration(userV1, userV1.Indices[0]))
}
