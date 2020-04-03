package user

import "github.com/hooone/evening/pkg/managers/sqlstore/migrator"

func AddUserMigration(mg *migrator.Migrator) {
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

	adminUserSql := "INSERT INTO `user` (`version`, `login`, `email`, `name`, `password`, `salt`, `rands`, `company`, `org_id`, `is_admin`, `email_verified`, `theme`, `created`, `updated`) VALUES ('1','admin', 'admin@localhost', 'admin', 'b649b8e44204e86004d027e0d57540c130d125520ab14534d132c08edb0d7bfc5c6c6dd3204c4cd0b31432aa8f43972636cc', 'grmKEzb4SU', 'EgIzpGxfud', '1', '1', '1', '1', '1','2020-03-25 17:44:24', '2020-03-25 17:44:27')"
	mg.AddMigration("insert admin user", migrator.NewRawSqlMigration(adminUserSql))

}

func AddUserTokenMigrations(mg *migrator.Migrator) {
	userTokenV1 := migrator.Table{
		Name: "user_token",
		Columns: []*migrator.Column{
			{Name: "id", Type: migrator.DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "user_id", Type: migrator.DB_BigInt, Nullable: false},
			{Name: "auth_token", Type: migrator.DB_NVarchar, Length: 100, Nullable: false},
			{Name: "prev_auth_token", Type: migrator.DB_NVarchar, Length: 100, Nullable: false},
			{Name: "user_agent", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "client_ip", Type: migrator.DB_NVarchar, Length: 255, Nullable: false},
			{Name: "auth_token_seen", Type: migrator.DB_Bool, Nullable: false},
			{Name: "seen_at", Type: migrator.DB_Int, Nullable: true},
			{Name: "rotated_at", Type: migrator.DB_Int, Nullable: false},
			{Name: "created_at", Type: migrator.DB_Int, Nullable: false},
			{Name: "updated_at", Type: migrator.DB_Int, Nullable: false},
		},
		Indices: []*migrator.Index{
			{Cols: []string{"auth_token"}, Type: migrator.UniqueIndex},
			{Cols: []string{"prev_auth_token"}, Type: migrator.UniqueIndex},
		},
	}

	mg.AddMigration("create user token table", migrator.NewAddTableMigration(userTokenV1))
	mg.AddMigration("add unique index user_token.token", migrator.NewAddIndexMigration(userTokenV1, userTokenV1.Indices[0]))
	mg.AddMigration("add unique index user_token.prev_token", migrator.NewAddIndexMigration(userTokenV1, userTokenV1.Indices[1]))
}
