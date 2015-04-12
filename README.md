# Shortify [![Build Status](https://travis-ci.org/pseudomuto/shortify-go.svg?branch=master)](https://travis-ci.org/pseudomuto/shortify-go)

A simple URL shortener application written in Go.

## Getting Started

```
$ git clone https://github.com/pseudomuto/shortify-go.git
$ script/build
```

Now you can take the compiled `shortify` executable and put it where ever you want.

## Using Shortify

Shortify is very simple. It has two endpoints:

* `GET /{token}` - will redirect to the full URL for the specified token
* `POST /redirects` - creates a new redirect (URLs are unique in this app)

Creating redirects requires a valid user account (see below). To create a redirect, you'll need to set the
auth header on the request and supply a JSON object with a `url` property. Here's an example using
`CURL`:

```
$ curl --user <username>:<password> \
> -H "Content-Type: application/json" \
> -d '{ "url": "http://pseudomuto.com/" }' \
> http://localhost:8080/redirects
```

The result (if successful) will be a JSON object similar to this one:

```json
{
  "id": 2,
  "token": "2Bk",
  "url": "http://pseudomuto.com/",
  "createdAt": "2015-04-12T22:41:36.544495662Z"
}
```

After running the command above, you can browse to `http://localhost:8080/2Bk` and you will be
redirected to `http://pseudomuto.com/`.

**NOTE**: _tokens are case-sensitive (i.e. `2Bk` != `2bk`)_

Hopefully errors don't occur, but when they do, the response will have the correct HTTP status code
and will have the following response body:

```json
{
  "code": <HTTP STATUS CODE>,
  "text": <ERROR MESSAGE>
}
```

## Running Shortify

Shortify is a standalone application. You can run it as a service, or just in the foreground if you
like.

`./shortify`

### Configuring the Database

__*The database will be created when the app is run. Be sure to follow these instructions before
running the app*__

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

#### A Quick Optimization

You can speed up the application by adding _unique_ indexes for the following columns after creating
the database.

* `"redirects"."token"`
* `"users"."name"`

In a future version, this will happen automatically.

### Managing Users

Creating a new `Redirect` record requires authentication. This app is configured to use basic
authentication (read: should be served over SSL).

Passwords are randomly generated and then hashed for storage in the database.

* List all users - `./shortify users list`
* Create a new user - `./shortify users create [username]`
* Generate a new password for a user - `./shortify users resetpw [username]`

_See `./shortify help` for options._

### Logging

Logs are very simple at this point. They include (tab seperated) the following information:

* Timestamp
* Request method
* Request URL
* The handler that handled the request
* How long it took to execute

Logs are written to `STDOUT` and no log file is kept.
