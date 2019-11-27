package mysql

import (
	"database/sql"
	"fmt"

	cloudsqlproxy "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	driver "github.com/go-sql-driver/mysql"
	"github.com/proproto/cloudsqldef/adapter"
)

type MysqlDatabase struct {
	config adapter.Config
	db     *sql.DB
}

func NewCloudSQL(cfg *driver.Config) (adapter.Database, error) {
	db, err := cloudsqlproxy.DialCfg(cfg)
	if err != nil {
		return nil, err
	}

	return &MysqlDatabase{db: db}, nil
}

func NewDatabase(config adapter.Config) (adapter.Database, error) {
	db, err := sql.Open("mysql", mysqlBuildDSN(config))
	if err != nil {
		return nil, err
	}

	return &MysqlDatabase{
		db:     db,
		config: config,
	}, nil
}

func (d *MysqlDatabase) TableNames() ([]string, error) {
	rows, err := d.db.Query("show full tables where Table_Type != 'VIEW'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := []string{}
	for rows.Next() {
		var table string
		var tableType string
		if err := rows.Scan(&table, &tableType); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (d *MysqlDatabase) DumpTableDDL(table string) (string, error) {
	var ddl string
	sql := fmt.Sprintf("show create table `%s`;", table) // TODO: escape table name

	err := d.db.QueryRow(sql).Scan(&table, &ddl)
	if err != nil {
		return "", err
	}

	return ddl, nil
}

func (d *MysqlDatabase) DB() *sql.DB {
	return d.db
}

func (d *MysqlDatabase) Close() error {
	return d.db.Close()
}

func mysqlBuildDSN(config adapter.Config) string {
	c := driver.NewConfig()
	c.User = config.User
	c.Passwd = config.Password
	c.DBName = config.DbName
	if config.Socket == "" {
		c.Net = "tcp"
		c.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	} else {
		c.Net = "unix"
		c.Addr = config.Socket
	}
	return c.FormatDSN()
}
