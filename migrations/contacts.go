package migrations

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/webstradev/rsdb-migrator/importer"
)

func MigrateContacts(db *sqlx.DB, data *importer.LoadedData) error {
	contacts := data.Contacts

	// Build query with bound args for all contacts
	query := fmt.Sprintf(`
	INSERT INTO contacts (name, title, source, privacy, email, phone, phone2, address, notes, object_id)
	VALUES 
	(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)%s`, strings.Repeat(",(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", len(contacts)-1))

	args := []any{}

	for _, contact := range contacts {
		args = append(args, contact.Name, contact.Title, contact.Source, contact.Privacy, contact.Email, contact.Phone, contact.Phone2, contact.Address, contact.Notes, contact.OldId.ObjectID)
	}

	_, err := db.Exec(query, args...)
	return err
}
