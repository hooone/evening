package migration

import (
	"fmt"
	"log"

	"github.com/hooone/evening/pkg/bus"
	"github.com/hooone/evening/pkg/managers/action"
	"github.com/hooone/evening/pkg/managers/card"
	"github.com/hooone/evening/pkg/managers/data"
	"github.com/hooone/evening/pkg/managers/field"
	"github.com/hooone/evening/pkg/managers/folder"
	"github.com/hooone/evening/pkg/managers/locale"
	"github.com/hooone/evening/pkg/managers/page"
	"github.com/hooone/evening/pkg/managers/parameter"
	"github.com/hooone/evening/pkg/managers/sqlstore"
	"github.com/hooone/evening/pkg/managers/sqlstore/migrator"
	"github.com/hooone/evening/pkg/managers/style"
	"github.com/hooone/evening/pkg/managers/user"
	"github.com/hooone/evening/pkg/registry"
	"github.com/hooone/evening/pkg/setting"
)

func init() {
	registry.RegisterService(&MigrationService{})
}

type MigrationService struct {
	Cfg      *setting.Cfg       `inject:""`
	Bus      bus.Bus            `inject:""`
	SQLStore *sqlstore.SQLStore `inject:""`
	log      log.Logger
}

func (ss *MigrationService) Init() error {
	//用migrator初始化数据库
	migrator := migrator.NewMigrator(ss.SQLStore.Engine)
	addMigrations(migrator)
	if err := migrator.Start(); err != nil {
		return fmt.Errorf("Migration failed err: %v", err)
	}
	return nil
}
func addMigrations(mg *migrator.Migrator) {
	addMigrationLogMigrations(mg)
	locale.AddLocaleMigration(mg)
	folder.AddFolderMigration(mg)
	page.AddPageMigration(mg)
	card.AddCardMigration(mg)
	field.AddFieldMigration(mg)
	action.AddViewActionMigration(mg)
	parameter.AddParameterMigration(mg)
	data.AddTestDataMigration(mg)
	style.AddStyleMigration(mg)
	style.AddStyleSetMigration(mg)
	user.AddUserMigration(mg)
	user.AddUserTokenMigrations(mg)
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
