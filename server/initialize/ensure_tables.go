package initialize

import (
	"context"
	adapter "github.com/casbin/gorm-adapter/v3"
	sysModel "wvp-go/server/model/system"
	"wvp-go/server/service/system"
	"gorm.io/gorm"
)

const initOrderEnsureTables = system.InitOrderExternal - 1

type ensureTables struct{}

// auto run
func init() {
	system.RegisterInit(initOrderEnsureTables, &ensureTables{})
}

func (e *ensureTables) InitializerName() string {
	return "ensure_tables_created"
}
func (e *ensureTables) InitializeData(ctx context.Context) (next context.Context, err error) {
	return ctx, nil
}

func (e *ensureTables) DataInserted(ctx context.Context) bool {
	return true
}

func (e *ensureTables) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	tables := []interface{}{
		sysModel.SysApi{},
		sysModel.SysUser{},
		sysModel.SysBaseMenu{},
		sysModel.SysAuthority{},
		sysModel.JwtBlacklist{},
		sysModel.SysDictionary{},
		sysModel.SysAutoCodeHistory{},
		sysModel.SysAIWorkflowSession{},
		sysModel.SysOperationRecord{},
		sysModel.SysDictionaryDetail{},
		sysModel.SysBaseMenuParameter{},
		sysModel.SysBaseMenuBtn{},
		sysModel.SysAuthorityBtn{},
		sysModel.SysAutoCodePackage{},
		sysModel.SysExportTemplate{},
		sysModel.Condition{},
		sysModel.JoinTemplate{},
		sysModel.SysParams{},
		sysModel.SysVersion{},
		sysModel.SysError{},
		sysModel.SysLoginLog{},
		sysModel.SysApiToken{},
		adapter.CasbinRule{},
		sysModel.Device{},
		sysModel.DeviceChannel{},
		sysModel.Alarm{},
		sysModel.DevicePosition{},
		sysModel.RecordPlan{},
		sysModel.RecordFile{},
		sysModel.Platform{},
		sysModel.PlatformChannel{},
	}
	for _, t := range tables {
		_ = db.AutoMigrate(&t)
		// 视图 authority_menu 会被当成表来创建，引发冲突错误（更新版本的gorm似乎不会）
		// 由于 AutoMigrate() 基本无需考虑错误，因此显式忽略
	}

	// 手动确保 GB28181 表拥有 GVA_MODEL 所需的字段
	// AutoMigrate 对已有表有时无法正确添加列
	ensureGB28181Columns(db)

	return ctx, nil
}

// ensureGB28181Columns 通过 raw SQL 确保 GB28181 相关表包含 GVA_MODEL 标准字段
func ensureGB28181Columns(db *gorm.DB) {
	type columnDef struct {
		table string
		col   string
		def   string
	}
	columns := []columnDef{
		{"wvp_device", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_device", "created_at", "DATETIME(3) NULL"},
		{"wvp_device", "updated_at", "DATETIME(3) NULL"},
		{"wvp_device", "deleted_at", "DATETIME(3) NULL"},
		{"wvp_device", "online", "TINYINT(1) DEFAULT 0"},

		{"wvp_device_channel", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_device_channel", "created_at", "DATETIME(3) NULL"},
		{"wvp_device_channel", "updated_at", "DATETIME(3) NULL"},
		{"wvp_device_channel", "deleted_at", "DATETIME(3) NULL"},

		{"wvp_alarm", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_alarm", "created_at", "DATETIME(3) NULL"},
		{"wvp_alarm", "updated_at", "DATETIME(3) NULL"},
		{"wvp_alarm", "deleted_at", "DATETIME(3) NULL"},

		{"wvp_device_position", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_device_position", "created_at", "DATETIME(3) NULL"},
		{"wvp_device_position", "updated_at", "DATETIME(3) NULL"},
		{"wvp_device_position", "deleted_at", "DATETIME(3) NULL"},

		{"wvp_record_plan", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_record_plan", "created_at", "DATETIME(3) NULL"},
		{"wvp_record_plan", "updated_at", "DATETIME(3) NULL"},
		{"wvp_record_plan", "deleted_at", "DATETIME(3) NULL"},

		{"wvp_record_file", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_record_file", "created_at", "DATETIME(3) NULL"},
		{"wvp_record_file", "updated_at", "DATETIME(3) NULL"},
		{"wvp_record_file", "deleted_at", "DATETIME(3) NULL"},

		{"wvp_platform", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_platform", "created_at", "DATETIME(3) NULL"},
		{"wvp_platform", "updated_at", "DATETIME(3) NULL"},
		{"wvp_platform", "deleted_at", "DATETIME(3) NULL"},

		{"wvp_platform_channel", "id", "INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY"},
		{"wvp_platform_channel", "created_at", "DATETIME(3) NULL"},
		{"wvp_platform_channel", "updated_at", "DATETIME(3) NULL"},
		{"wvp_platform_channel", "deleted_at", "DATETIME(3) NULL"},
	}

	for _, c := range columns {
		if !db.Migrator().HasTable(c.table) {
			continue
		}
		if !db.Migrator().HasColumn(c.table, c.col) {
			sql := "ALTER TABLE `" + c.table + "` ADD COLUMN `" + c.col + "` " + c.def
			db.Exec(sql)
		}
	}
}

func (e *ensureTables) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	tables := []interface{}{
		sysModel.SysApi{},
		sysModel.SysUser{},
		sysModel.SysBaseMenu{},
		sysModel.SysAuthority{},
		sysModel.JwtBlacklist{},
		sysModel.SysDictionary{},
		sysModel.SysAutoCodeHistory{},
		sysModel.SysAIWorkflowSession{},
		sysModel.SysOperationRecord{},
		sysModel.SysDictionaryDetail{},
		sysModel.SysBaseMenuParameter{},
		sysModel.SysBaseMenuBtn{},
		sysModel.SysAuthorityBtn{},
		sysModel.SysAutoCodePackage{},
		sysModel.SysExportTemplate{},
		sysModel.Condition{},
		sysModel.JoinTemplate{},

		adapter.CasbinRule{},
	}
	yes := true
	for _, t := range tables {
		yes = yes && db.Migrator().HasTable(t)
	}
	return yes
}
