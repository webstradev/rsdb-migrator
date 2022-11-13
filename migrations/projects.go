package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

// Need to migrate dates from 2017 to 09/09/2017 and 2019 to 09/09/2019
func MigrateProjects(db *sqlx.DB, data *importer.LoadedData) error {
	projects := data.Projects

	// Build insert query with bound args for all projects
	query := fmt.Sprintf(`
	INSERT INTO projects (title, link, description, date, body, object_id)
	VALUES 
	(?, ?, ?, STR_TO_DATE(?, ?), ?, ?)%s`, strings.Repeat(",(?, ?, ?, STR_TO_DATE(?, ?), ?, ?)", len(projects)-1))

	args := []any{}

	for _, project := range projects {
		args = append(args, project.Title, project.Link, project.Description, project.Date, "%d/%m/%Y", project.Body, project.OldId.ObjectID)
	}

	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	_, err = db.Exec("UPDATE projects SET date = NULL WHERE date = 0000-00-00")
	if err != nil {
		return err
	}

	return nil
}
