package shortify

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
	"os"
)

const dbDriverEnvVar = "SHORTIFY_DB_DRIVER"
const dbDataSourceEnvVar = "SHORTIFY_DB_DATASOURCE"

var testDbConnectionInfo = dbConnectionInfo{"sqlite3", "/tmp/redirects_db_test.bin"}
var prodDbConnectionInfo = dbConnectionInfo{os.Getenv(dbDriverEnvVar), os.Getenv(dbDataSourceEnvVar)}
var db = database{prodDbConnectionInfo, false}

type dbConnectionInfo struct {
	driver     string
	dataSource string
}

func (self dbConnectionInfo) Dialect() gorp.Dialect {
	switch self.driver {
	case "mysql":
		return gorp.MySQLDialect{"InnoDB", "UTF8"}
	case "postgres":
		return gorp.PostgresDialect{}
	default:
		return gorp.SqliteDialect{}
	}
}

type database struct {
	connectionInfo dbConnectionInfo
	inited         bool
}

func (self database) reset() error {
	return withConnection(func(dbMap *gorp.DbMap) error {
		return dbMap.TruncateTables()
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
	dbMap := &gorp.DbMap{Db: sqlDb, Dialect: db.connectionInfo.Dialect()}
	dbMap.AddTableWithName(Redirect{}, "redirects").SetKeys(true, "Id")
	dbMap.AddTableWithName(User{}, "users").SetKeys(true, "Id")
	return dbMap
}

func withConnection(routine func(*gorp.DbMap) error) error {
	sqlDb, err := sql.Open(db.connectionInfo.driver, db.connectionInfo.dataSource)
	if err != nil {
		return err
	}

	defer sqlDb.Close()

	dbMap := mapForDatabase(sqlDb)
	if !db.inited {
		dbMap.CreateTablesIfNotExists()
		db.inited = true
	}

	return routine(dbMap)
}

func useTestingDatabase() {
	db = database{testDbConnectionInfo, false}
}
