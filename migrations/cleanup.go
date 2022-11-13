package migrations

import "github.com/jmoiron/sqlx"

func CleanupTables(db *sqlx.DB) error {
	// Delete old contacts that are not linked to a platform
	_, err := db.Exec("DELETE FROM contacts WHERE platform_id IS NULL")
	if err != nil {
		return err
	}

	// Temporarily make platform_id nullable due to some platformless contacts in the mongo database
	_, err = db.Exec("ALTER TABLE `contacts` CHANGE COLUMN `platform_id` `platform_id` INT(11) NOT NULL AFTER `deleted_at`;")
	if err != nil {
		return err
	}

	// Drop object_id column which is no longer necessary when the migration is complete
	_, err = db.Exec("ALTER TABLE `articles` DROP COLUMN `object_id`;")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `platforms` DROP COLUMN `object_id`;")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `contacts` DROP COLUMN `object_id`;")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `projects` DROP COLUMN `object_id`;")
	if err != nil {
		return err
	}

	return nil
}
