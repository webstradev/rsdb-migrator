# rsdb-migrator
This repository holds the code that was used when migrating the rsdb-beta version from mongodb to mysql.

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
