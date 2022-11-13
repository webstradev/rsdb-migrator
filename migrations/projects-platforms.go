package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigrateProjectsToPlatforms(db *sqlx.DB, data *importer.LoadedData) error {
	projects := data.Projects

	// Start projects loop
	for _, project := range projects {
		// Go to next project if there are no platforms linked to this project
		if len(project.Platforms) == 0 {
			continue
		}

		// Get database primary key for this project
		projectId := 0
		err := db.Get(&projectId, "SELECT id FROM projects WHERE object_id = ?", project.OldId.ObjectID)
		if err != nil {
			return err
		}

		// Go to the next project if we cant find the project in the database
		if projectId == 0 {
			continue
		}

		// Create a list of all platform object ids
		platformObjectIds := []any{}
		for _, platform := range project.Platforms {
			platformObjectIds = append(platformObjectIds, platform.ObjectID)
		}
		platformIds := []any{}

		// Get all the database primary keys for the platforms linked to this project
		err = db.Select(&platformIds, fmt.Sprintf(`SELECT id FROM platforms WHERE object_id IN (?%s)`, strings.Repeat(",?", len(project.Platforms)-1)), platformObjectIds...)
		if err != nil {
			return err
		}

		// Go to the next project if we cannot find any linked platforms in the database
		if len(platformIds) == 0 {
			continue
		}

		// Build insert query for all the linked platforms for this project
		query := fmt.Sprintf(`
		INSERT INTO 
			platforms_projects (project_id, platform_id) 
		VALUES
			(?,?)%s`, strings.Repeat(",(?,?)", len(platformIds)-1))

		args := []any{}
		for _, plt := range platformIds {
			args = append(args, projectId, plt)
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
