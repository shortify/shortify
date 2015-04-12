package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/gorp.v1"
)

type dbConnectionDetails struct {
	driver     string
	dataSource string
	dialect    gorp.Dialect
}

var prodDb = dbConnectionDetails{"sqlite3", "/tmp/redirects_db.bin", gorp.SqliteDialect{}}
var testDb = dbConnectionDetails{"sqlite3", "/tmp/redirects_db_test.bin", gorp.SqliteDialect{}}
var currentDb = prodDb

func mapForDatabase(db *sql.DB) *gorp.DbMap {
	dbMap := &gorp.DbMap{Db: db, Dialect: currentDb.dialect}
	dbMap.AddTableWithName(Redirect{}, "redirects").SetKeys(true, "Id")
	return dbMap
}

func openDb() (*sql.DB, error) {
	return sql.Open(currentDb.driver, currentDb.dataSource)
}

func SetCurrentDb(testing bool) {
	currentDb = prodDb
	if testing {
		currentDb = testDb
	}
}

func InitializeDb() error {
	db, err := openDb()
	if err != nil {
		return err
	}

	defer db.Close()
	dbMap := mapForDatabase(db)
	return dbMap.CreateTablesIfNotExists()
}

func TruncateDb() error {
	db, err := openDb()
	if err != nil {
		return err
	}

	defer db.Close()
	dbMap := mapForDatabase(db)
	return dbMap.TruncateTables()
}

func DbInsert(entities ...interface{}) error {
	db, err := openDb()
	if err != nil {
		return err
	}

	defer db.Close()
	dbMap := mapForDatabase(db)
	return dbMap.Insert(entities...)
}

func DbUpdate(entities ...interface{}) (int64, error) {
	db, err := openDb()
	if err != nil {
		return 0, err
	}

	defer db.Close()
	dbMap := mapForDatabase(db)
	return dbMap.Update(entities...)
}

func DbSelectOne(holder interface{}, query string, args ...interface{}) error {
	db, err := openDb()
	if err != nil {
		return err
	}

	defer db.Close()
	dbMap := mapForDatabase(db)
	return dbMap.SelectOne(holder, query, args...)
}
