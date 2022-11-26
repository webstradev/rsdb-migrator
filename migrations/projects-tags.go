package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigrateProjectsTags(db *sqlx.DB, data *importer.LoadedData) error {
	projects := data.Projects

	// Create a list of all tags for all projects
	tags := []string{}
	for _, project := range projects {
		tags = append(tags, project.Tags...)
	}

	if len(tags) == 0 {
		return nil
	}

	// Build an insert query with bound args for all tags (ignoring duplicates)
	query := fmt.Sprintf(`
	INSERT IGNORE INTO tags (tag)
	VALUES 
	(?)%s`, strings.Repeat(",(?)", len(tags)-1))

	args := []any{}
	for _, tag := range tags {
		args = append(args, strings.ToLower(strings.TrimSpace(tag)))
	}

	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Get all the tags from the database
	dbTags := []struct {
		ID  int64  `db:"id"`
		Tag string `db:"tag"`
	}{}
	tagToId := map[string]int64{}

	err = db.Select(&dbTags, "SELECT id, tag FROM tags")
	if err != nil {
		return err
	}

	// Create a map from tag to database primary key
	for _, dbTag := range dbTags {
		tagToId[dbTag.Tag] = dbTag.ID
	}

	// Start projects loop
	for _, project := range projects {
		// Go to the next projects  if there are no tags for this project
		if len(project.Tags) == 0 {
			continue
		}

		// Get database primary key for this project
		projectId := 0
		err = db.Get(&projectId, "SELECT id FROM projects WHERE object_id = ?", project.OldId.ObjectID)
		if err != nil {
			return err
		}

		// Go to the next project if we cant find the project in the database
		if projectId == 0 {
			continue
		}

		// Build insert query for all tags for this project (ignoring duplicates)
		query := fmt.Sprintf(`
		INSERT IGNORE INTO projects_tags (project_id, tag_id)
		VALUES 
		(?, ?)%s`, strings.Repeat(",(?,?)", len(project.Tags)-1))

		args := []any{}
		for _, tag := range project.Tags {
			// Convert all tags to lower case and trim spaces
			args = append(args, projectId, tagToId[strings.ToLower(strings.TrimSpace(tag))])
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
