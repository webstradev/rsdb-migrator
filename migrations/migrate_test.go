package migrations

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func TestMigrate(t *testing.T) {
	// Create mock database
	mockDb, mockSql, err := sqlmock.New()
	if err != nil {
		t.Errorf("error creating database mock")
	}

	// Close this database at the end of the test
	defer mockDb.Close()

	// Create sqlx instance which an be passed to functions
	mockedDatabase := sqlx.NewDb(mockDb, "sqlmock")

	// Expect PrepareTables queries
	mockSql.ExpectExec("ALTER TABLE `articles`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `contacts`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `platforms`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `projects`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `contacts`").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigratePlatforms queries
	mockSql.ExpectExec("INSERT INTO platforms").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigratePlatformsToCategories queries
	mockSql.ExpectQuery("SELECT id, category FROM categories").WillReturnRows(sqlmock.NewRows([]string{"id", "category"}).AddRow(1, "Category1"))
	mockSql.ExpectQuery("SELECT id FROM platforms").WithArgs("1a1").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectExec("INSERT INTO platforms_categories").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateContacts queries
	mockSql.ExpectExec("INSERT INTO contacts").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigratePlatformsToContacts queries
	mockSql.ExpectQuery("SELECT id FROM platforms").WithArgs("1a1").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectExec("UPDATE contacts").WithArgs(1, "a1a").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateArticles queries
	mockSql.ExpectExec("INSERT INTO articles").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateArticlesToPlatforms queries
	mockSql.ExpectQuery("SELECT id FROM articles").WithArgs("123").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectQuery("SELECT id FROM platforms").WithArgs("abc").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectExec("INSERT INTO platforms_articles").WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateArticlesTags queries
	mockSql.ExpectExec("INSERT IGNORE INTO tags").WithArgs("tag1", "tag2", "tag1").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectQuery("SELECT id, tag FROM tags").WillReturnRows(sqlmock.NewRows([]string{"id", "tag"}).AddRow(1, "tag1").AddRow(2, "tag2"))
	mockSql.ExpectQuery("SELECT id FROM articles").WithArgs("123").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectExec("INSERT IGNORE INTO articles_tags").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectQuery("SELECT id FROM articles").WithArgs("456").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	mockSql.ExpectExec("INSERT IGNORE INTO articles_tags").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateProjects queries
	mockSql.ExpectExec("INSERT INTO projects").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("UPDATE projects").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateProjectsToPlatforms queries
	mockSql.ExpectQuery("SELECT id FROM projects").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectQuery("SELECT id FROM platforms").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectExec("INSERT INTO platforms_projects").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectQuery("SELECT id FROM projects").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	mockSql.ExpectQuery("SELECT id FROM platforms").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(2))
	mockSql.ExpectExec("INSERT INTO platforms_projects").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect MigrateProjectsTags queries
	mockSql.ExpectExec("INSERT IGNORE INTO tags").WithArgs("tag1").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectQuery("SELECT id, tag FROM tags").WillReturnRows(sqlmock.NewRows([]string{"id", "tag"}).AddRow(1, "tag1").AddRow(2, "tag2"))
	mockSql.ExpectQuery("SELECT id FROM projects").WithArgs("11a").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mockSql.ExpectExec("INSERT IGNORE INTO projects_tags").WillReturnResult(sqlmock.NewResult(0, 0))

	// Expect CleanupTables queries
	mockSql.ExpectExec("DELETE FROM contacts").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `contacts`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `articles`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `platforms`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `contacts`").WillReturnResult(sqlmock.NewResult(0, 0))
	mockSql.ExpectExec("ALTER TABLE `projects`").WillReturnResult(sqlmock.NewResult(0, 0))

	// Call Migrate function and fail test if any error occurs
	err = Migrate(mockedDatabase, &importer.ExpectedData)
	if err != nil {
		t.Fatalf("MigratePlatforms() error = %v", err)
	}
}
