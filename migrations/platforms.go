package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigratePlatforms(db *sqlx.DB, data *importer.LoadedData) error {
	platforms := data.Platforms

	// Build insert query with bound args for all platforms
	query := fmt.Sprintf(`
	INSERT INTO platforms (name, website, source, privacy, country, notes, comment, object_id)
	VALUES 
	(?, ?, ?, ?, ?, ?, ?, ?)%s`, strings.Repeat(",(?, ?, ?, ?, ?, ?, ?, ?)", len(platforms)-1))

	args := []any{}

	for _, platform := range platforms {
		args = append(args, platform.Name, platform.Website, platform.Source, platform.Privacy, platform.Country, platform.Notes, platform.Comment, platform.OldId.ObjectID)
	}

	_, err := db.Exec(query, args...)

	return err
}
