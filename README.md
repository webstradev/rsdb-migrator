# rsdb-migrator
[![Run Tests](https://github.com/webstradev/rsdb-migrator/actions/workflows/test.yml/badge.svg)](https://github.com/webstradev/rsdb-migrator/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/webstradev/rsdb-migrator)](https://goreportcard.com/report/github.com/webstradev/rsdb-migrator)
[![CodeQL](https://github.com/webstradev/rsdb-migrator/actions/workflows/codeql.yml/badge.svg)](https://github.com/webstradev/gin-migrator/actions/workflows/codeql.yml)
This repository holds the code that was used when migrating the rsdb-beta version from mongodb to mysql.

*This code is not to be distributed or used in any way without explicit permission of the Erik Westra*

## Migration steps
1. Create a mysql database and initialize the schemas by running the [rsdb-backend](https://github.com/webstradev/rsdb-backend) (this will automatically create all the required tables)
2. Get a JSON dump of the mongodb database collections (articles, platforms, projects, contacts) using [mongoexport](https://www.mongodb.com/docs/database-tools/mongoexport/)
3. For the articles and projects tables there were four or five entries where `tags` was specified as a string instead of an array of strings, I manually changed these in the json file as there were only a few cases of them.
4. Pull this repository to your local machine.
5. Create an .env environment file with a database connection string to connect to your mysql database following the format below.
  `DB_CONNECTION_STRING="*USER*:*PASSWORD*@*PROTOCOL*(*HOST*:*PORT*)/*DATABASENAME*?parseTime=true"`
6. Place the json files in the repo using the following file structure:
```
- .env
- mongodump/
  -articles.json
  -contacts.json
  -platforms.json
  -projects.json
- main.go
```
7. Run the main file `go run main.go` to populate the database with all the migrated data from mongodb

## Adding Edge cases
The LoadedData Type which holds all the data after it has been unmarshalled from json contains a function `HandleEdgeCases()` located in `./importer/edgecases.go`. This function can be used to add any manipulations to any of the data arrays after being marshalled from JSON. I have used this to handle any weird edge cases in the mongodb database.

 ---

## Running unit tests
Unit tests can be run using the `make unit-test` command or any of the normal testing commands for go. The advantage of using the make command is that it will install a small package called tparse which will parse the test coverage and success results into a nice table to display in the command line:
![image](https://user-images.githubusercontent.com/82543732/204091447-2296e5b4-6b3d-4802-9f56-cec1d90c9396.png)

