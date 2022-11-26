package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigratePlatformsToContacts(db *sqlx.DB, data *importer.LoadedData) error {
	platforms := data.Platforms

	// Start platform loop
	for _, platform := range platforms {
		// Go to next platform if there are no contacts for this platform
		if len(platform.Contacts) == 0 {
			continue
		}

		// Get the platforms primary key in the database
		platformId := 0
		err := db.Get(&platformId, "SELECT id FROM platforms WHERE object_id = ?", platform.OldId.ObjectID)
		if err != nil {
			return err
		}

		// Go to next platform if we cant find the platform in the database
		if platformId == 0 {
			continue
		}

		// Build update query with bound args for this platform
		query := fmt.Sprintf(`
		UPDATE 
			contacts 
		SET
			platform_id = ?
		WHERE object_id IN (?%s)`, strings.Repeat(",?", len(platform.Contacts)-1))

		args := []any{platformId}

		for _, contact := range platform.Contacts {
			args = append(args, contact.ObjectID)
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
