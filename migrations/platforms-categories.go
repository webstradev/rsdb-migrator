package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigratePlatformsToCategories(db *sqlx.DB, data *importer.LoadedData) error {
	platforms := data.Platforms

	// Get database categories and create a map from category to cateogry primary key
	dbCategories := []struct {
		ID       int64  `db:"id"`
		Category string `db:"category"`
	}{}
	categoryToId := map[string]int64{}

	err := db.Select(&dbCategories, "SELECT id, category FROM categories")
	if err != nil {
		return err
	}

	for _, dbCategory := range dbCategories {
		categoryToId[dbCategory.Category] = dbCategory.ID
	}

	// Start platform loop
	for _, platform := range platforms {
		// Go to next platform if there is no category
		if platform.Category == "" {
			continue
		}

		// Categories is stored as an array of strings in the mongodb database
		categories := strings.Split(platform.Category, ",")

		// Get database primary key for this platform
		platformId := 0
		db.Get(&platformId, "SELECT id FROM platforms WHERE object_id = ?", platform.OldId.ObjectID)

		// Go to next platform if we cant find the platform in the database
		if platformId == 0 {
			continue
		}

		// Build an insert query with bound args for this platforms categories
		query := fmt.Sprintf(`
		INSERT INTO 
			platforms_categories (platform_id, category_id)
		VALUES
			(?,?)%s`, strings.Repeat(",(?,?)", len(categories)-1))

		args := []any{}

		for _, cat := range categories {
			args = append(args, platformId, categoryToId[cat])
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
