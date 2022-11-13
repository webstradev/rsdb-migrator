package migrations

import (
	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func Migrate(db *sqlx.DB, data *importer.LoadedData) error {
	// Prepare some database schema alterations that are requires for the migration
	err := PrepareTables(db)
	if err != nil {
		return err
	}

	// Migrate all platforms and platform categories
	err = MigratePlatforms(db, data)
	if err != nil {
		return err
	}
	err = MigratePlatformsToCategories(db, data)
	if err != nil {
		return err
	}

	// Migrate contacts and map contacts to platforms
	err = MigrateContacts(db, data)
	if err != nil {
		return err
	}
	err = MigratePlatformsToContacts(db, data)
	if err != nil {
		return err
	}

	// Migrate articles and map articles to platforms
	err = MigrateArticles(db, data)
	if err != nil {
		return err
	}
	err = MigrateArticlesToPlatforms(db, data)
	if err != nil {
		return err
	}

	// Migrate tags and map articles to tags
	err = MigrateArticlesTags(db, data)
	if err != nil {
		return err
	}

	// Migrate project and map projects to platforms
	err = MigrateProjects(db, data)
	if err != nil {
		return err
	}
	err = MigrateProjectsToPlatforms(db, data)
	if err != nil {
		return err
	}

	// Migrate tags and map projects to tags
	err = MigrateProjectsTags(db, data)
	if err != nil {
		return err
	}

	// Clean up some stale records and undo schema alterations
	err = CleanupTables(db)
	if err != nil {
		return err
	}

	return nil
}
