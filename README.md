# Shortify

A simple URL shortener application written in Go.

## Getting Started

```
$ git clone https://github.com/pseudomuto/shortify-go.git
$ script/build
```

Now you can take the compiled `shortify` executable and put it where ever you want.

## Running Shortify

Shortify is a standalone application. You can run it as a service, or just in the foreground if you
like.

### Configuring the Database

Shortify uses environment variables to handle database connections. Currently, there is support for
MySQL, PostgreSQL and Sqlite3. Very happy to receive PRs for others if you need them.

If you're having connectivity issues, please consult the corresponding driver's source for
connection string details.

* MySQL
    * `SHORTIFY_DB_DRIVER=mysql`
    * `SHORTIFY_DB_DATASOURCE=tcp:localhost:3306*mydb/myuser/mypassword`
* PostgreSQL
    * `SHORTIFY_DB_DRIVER=postgres`
    * `SHORTIFY_DB_DATASOURCE="user=<username> password=<password> dbname=<dbName> sslmode=disable"`
* Sqlite3
    * `SHORTIFY_DB_DRIVER=sqlite3`
    * `SHORTIFY_DB_DATASOURCE=/path/to/db/file.bin`

You can also just supply these on the command line if you prefer:

`SHORTIFY_DB_DRIVER=sqlite3 SHORTIFY_DB_DATASOURCE=local_db.bin ./shortify`

### Logging

Logs are very simple at this point. They include (tab seperated) the following information:

* Timestamp
* Request method
* Request URL
* The handler that handled the request
* How long it took to execute

Logs are written to `STDOUT` and no log file is kept.
