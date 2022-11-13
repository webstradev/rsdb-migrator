package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigrateArticles(db *sqlx.DB, data *importer.LoadedData) error {
	articles := data.Articles

	// Build insert query with bound args for all articles
	// STR_TO_DATE is used to parse the dates from strings in the "dd/mm/yyy" to the sql date types
	query := fmt.Sprintf(`
	INSERT INTO articles (title, link, description, date, body, object_id)
	VALUES 
	(?, ?, ?, STR_TO_DATE(?, ?), ?, ?)%s`, strings.Repeat(",(?, ?, ?, STR_TO_DATE(?, ?), ?, ?)", len(articles)-1))

	args := []any{}

	for _, article := range articles {
		args = append(args, article.Title, article.Link, article.Description, article.Date, "%d/%m/%Y", article.Body, article.OldId.ObjectID)
	}

	_, err := db.Exec(query, args...)
	return err
}
