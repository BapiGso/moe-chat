package database

import (
	_ "embed"
	"github.com/jmoiron/sqlx"
	"log/slog"
	_ "modernc.org/sqlite"
)

//go:embed moechat.sql
var sqlTable string

var DB = func() *sqlx.DB {
	//if _, err := os.Stat("moechat.db"); os.IsNotExist(err) {
	//	exePath, _ := os.Executable()
	//	_ = os.Chdir(filepath.Dir(exePath))
	//}

	db := sqlx.MustConnect("sqlite", "moechat.db")
	//读取sql文件创建表
	if _, err := db.Exec(sqlTable); err != nil {
		slog.Error("创建表结构失败")
	}
	return db
}()
