package shortify

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

var shortifyDb database

type database struct {
	connectionInfo dbConnectionInfo
	inited         bool
}

type dbConnectionInfo struct {
	provider   string
	dataSource string
}

func (self dbConnectionInfo) Dialect() gorp.Dialect {
	switch self.provider {
	case "mysql":
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "postgres":
		return gorp.PostgresDialect{}
	default:
		return gorp.SqliteDialect{}
	}
}

func newDatabase(provider string, dataSource string) database {
	connectionInfo := dbConnectionInfo{provider, dataSource}
	return database{connectionInfo, false}
}

func (self database) reset() error {
	return withConnection(func(dbMap *gorp.DbMap) error {
		dbMap.DropTables()
		return dbMap.CreateTablesIfNotExists()
	})
}

func (self database) insert(entities ...interface{}) error {
	return withConnection(func(dbMap *gorp.DbMap) error {
		return dbMap.Insert(entities...)
	})
}

func (self database) update(entities ...interface{}) (int64, error) {
	affectedRecordCount := int64(0)
	err := withConnection(func(dbMap *gorp.DbMap) error {
		records, err := dbMap.Update(entities...)
		affectedRecordCount = records
		return err
	})

	return affectedRecordCount, err
}

func (self database) selectAll(holder interface{}, query string, args ...interface{}) ([]interface{}, error) {
	var records []interface{}
	err := withConnection(func(dbMap *gorp.DbMap) error {
		rows, err := dbMap.Select(holder, query, args...)
		records = rows
		return err
	})

	return records, err
}

func (self database) selectOne(holder interface{}, query string, args ...interface{}) error {
	return withConnection(func(dbMap *gorp.DbMap) error {
		return dbMap.SelectOne(holder, query, args...)
	})
}

func mapForDatabase(sqlDb *sql.DB) *gorp.DbMap {
	dbMap := &gorp.DbMap{Db: sqlDb, Dialect: shortifyDb.connectionInfo.Dialect()}
	dbMap.AddTableWithName(Redirect{}, "redirects").SetKeys(true, "Id").ColMap("token").SetUnique(true)
	dbMap.AddTableWithName(User{}, "users").SetKeys(true, "Id").ColMap("name").SetUnique(true)
	return dbMap
}

func withConnection(routine func(*gorp.DbMap) error) error {
	sqlDb, err := sql.Open(shortifyDb.connectionInfo.provider, shortifyDb.connectionInfo.dataSource)
	if err != nil {
		return err
	}

	defer sqlDb.Close()

	dbMap := mapForDatabase(sqlDb)
	if !shortifyDb.inited {
		dbMap.CreateTablesIfNotExists()
		shortifyDb.inited = true
	}

	return routine(dbMap)
}
