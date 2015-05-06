# Shortify 

[![Build Status](https://travis-ci.org/shortify/shortify.svg?branch=master)](https://travis-ci.org/shortify/shortify)
[![GoDoc](https://godoc.org/github.com/shortify/shortify/app?status.svg)](https://godoc.org/github.com/shortify/shortify/app)

A simple URL shortener application written in Go.

## Getting Started

Head over to the [latest release] and download the `shortify` executable.

[latest release]: https://github.com/shortify/shortify/releases/latest

### Building From Source

```bash
git clone https://github.com/shortify/shortify.git
script/build
```

### Configuration

In order to run shortify, you'll need to create a config file named `shortify.gcfg`. This file must be placed right 
beside the executable.

Here's an example of a typical config file:

```ini
[database]
provider = mysql
dataSource = tcp://localhost:3306*mydb/myuser/mypassword

[settings]
; alphanumeric without abiguous characters
alphabet = 23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
port = 80
```

Any of these settings can be supplied as an env var by using `$ENV_VAR_NAME` as the value. For example, if you have the
following environment variables set:

```bash
export SHORTIFY_DATASOURCE=tcp://localhost:3306*mydb/myuser/mypassword
export SHORTIFY_PORT=8080
```

Then you can use those values in the configuration like this:

```ini
[database]
provider = mysql
dataSource = $SHORTIFY_DATASOURCE

[settings]
; alphanumeric without abiguous characters
alphabet = 23456789abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
port = $SHORTIFY_PORT
```

### Running the App

With the config file in place, simply run `./shortify`. If the database does not exist yet, shortify will create it for
you.

Currently, there is support for MySQL, PostgreSQL and Sqlite3. Very happy to receive PRs for others if you need them.

For examples of how to configure each of these, check out the _examples/_ directory.

## Using Shortify

Shortify is very simple. It has two endpoints:

* `GET /{token}` - will redirect to the full URL for the specified token
* `POST /redirects` - creates a new redirect (URLs are unique in this app)

Creating redirects requires a valid user account (see below). To create a redirect, you'll need to set the
auth header on the request and supply a JSON object with a `url` property. Here's an example using
`CURL`:

```bash
curl --user <username>:<password> \
  -H "Content-Type: application/json" \
  -d '{ "url": "http://pseudomuto.com/" }' \
  http://localhost:8080/redirects
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
and will have the following response body (code and text change depending on the error):

```json
{
  "code": 404,
  "text": "Redirect was not found"
}
```

## Managing Users

Creating a new `Redirect` record requires authentication. This app is configured to use basic
authentication (read: should be served over SSL).

Passwords are randomly generated and then hashed for storage in the database.

* List all users - `./shortify users list`
* Create a new user - `./shortify users create [username]`
* Generate a new password for a user - `./shortify users resetpw [username]`

_See `./shortify help` for options._

## Logging

Logs are very simple at this point. They include (tab seperated) the following information:

* Timestamp
* Request method
* Request URL
* The handler that handled the request
* How long it took to execute

Logs are written to `STDOUT` and no log file is kept.
