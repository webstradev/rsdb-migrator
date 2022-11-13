package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigrateArticlesToPlatforms(db *sqlx.DB, data *importer.LoadedData) error {
	articles := data.Articles

	// Start article loop
	for _, article := range articles {
		// Go to next platform if there are no Platforms
		if len(article.Platforms) == 0 {
			continue
		}

		// Get the article primary key from the database
		articleId := 0
		err := db.Get(&articleId, "SELECT id FROM articles WHERE object_id = ?", article.OldId.ObjectID)
		if err != nil {
			return err
		}

		// Go to next article if we cant find the article in the database
		if articleId == 0 {
			continue
		}

		// Create a list of platform object ids linked to this article
		platformObjectIds := []any{}
		for _, platform := range article.Platforms {
			platformObjectIds = append(platformObjectIds, platform.ObjectID)
		}

		platformIds := []any{}

		// Get platform primary keys from the database for all the platforms linked to this article
		err = db.Select(&platformIds, fmt.Sprintf(`SELECT id FROM platforms WHERE object_id IN (?%s)`, strings.Repeat(",?", len(article.Platforms)-1)), platformObjectIds...)
		if err != nil {
			return err
		}

		// Go to the next platform if we cant find any matching platforms in the database
		if len(platformIds) == 0 {
			continue
		}

		// Build a query with bound args for all of the platforms linked to this article
		query := fmt.Sprintf(`
		INSERT INTO 
			platforms_articles (article_id, platform_id) 
		VALUES
			(?,?)%s`, strings.Repeat(",(?,?)", len(platformIds)-1))

		args := []any{}
		for _, plt := range platformIds {
			args = append(args, articleId, plt)
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
