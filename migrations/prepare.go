package migrations

import "github.com/jmoiron/sqlx"

func PrepareTables(db *sqlx.DB) error {
	// Add object_id column which is the pkey in the mongodb structure
	_, err := db.Exec("ALTER TABLE `articles` ADD COLUMN `object_id` VARCHAR(100) NULL DEFAULT NULL;")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `contacts` ADD COLUMN `object_id` VARCHAR(100) NULL DEFAULT NULL;")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `platforms` ADD COLUMN `object_id` VARCHAR(100) NULL DEFAULT NULL;")
	if err != nil {
		return err
	}
	_, err = db.Exec("ALTER TABLE `projects`	ADD COLUMN `object_id` VARCHAR(100) NULL DEFAULT NULL;")
	if err != nil {
		return err
	}

	// Make platform_id nullable for the migration
	_, err = db.Exec("ALTER TABLE `contacts` CHANGE COLUMN `platform_id` `platform_id` INT(11) NULL AFTER `deleted_at`;")
	if err != nil {
		return err
	}

	return nil
}
