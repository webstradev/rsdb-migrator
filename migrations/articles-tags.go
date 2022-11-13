package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigrateArticlesTags(db *sqlx.DB, data *importer.LoadedData) error {
	articles := data.Articles

	// Create a list of all tags for all articles
	tags := []string{}
	for _, article := range articles {
		tags = append(tags, article.Tags...)
	}

	// Build insert query and bound args for all tags (ignoring duplicates)
	query := fmt.Sprintf(`
	INSERT IGNORE INTO tags (tag)
	VALUES 
	(?)%s`, strings.Repeat(",(?)", len(tags)-1))

	args := []any{}
	for _, tag := range tags {
		// Trim spaces and convert to lowercase
		args = append(args, strings.ToLower(strings.TrimSpace(tag)))
	}

	_, err := db.Exec(query, args...)
	if err != nil {
		return err
	}

	// Get all tags from the database
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
	for _, dbCategory := range dbTags {
		tagToId[dbCategory.Tag] = dbCategory.ID
	}

	// Start articles loop
	for _, article := range articles {
		// Go to the next article if this article doesn't have any tags
		if len(article.Tags) == 0 {
			continue
		}

		// Get the database primary key for this article
		articleId := 0
		err = db.Get(&articleId, "SELECT id FROM articles WHERE object_id = ?", article.OldId.ObjectID)
		if err != nil {
			return err
		}

		// Go to next article if we cant find the article in the database
		if articleId == 0 {
			continue
		}

		// Build insert query for all the tags for this article (ignoring duplicate combinations)
		query := fmt.Sprintf(`
		INSERT IGNORE INTO articles_tags (article_id, tag_id)
		VALUES 
		(?, ?)%s`, strings.Repeat(",(?,?)", len(article.Tags)-1))

		args := []any{}
		for _, tag := range article.Tags {
			// Trim spaces and convert to lowercase
			args = append(args, articleId, tagToId[strings.ToLower(strings.TrimSpace(tag))])
		}

		_, err = db.Exec(query, args...)
		if err != nil {
			return err
		}
	}
	return nil
}
