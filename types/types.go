package types

// An ObjectReference is a reference to an object using its mongodb object key. This will only be stored in the database
type ObjectReference struct {
	ObjectID string `json:"$oid"`
}

type Article struct {
	OldId struct {
		ObjectID string `json:"$oid"`
	} `json:"_id"`
	Title       string            `json:"title" db:"title"`
	Link        string            `json:"link" db:"link"`
	Description string            `json:"description" db:"description"`
	Date        string            `json:"date" db:"date"`
	Body        string            `json:"body" db:"body"`
	Platforms   []ObjectReference `json:"platforms"`
	Tags        []string          `json:"tags"`
}

type Contact struct {
	OldId struct {
		ObjectID string `json:"$oid"`
	} `json:"_id"`
	Title   string `json:"title" db:"title"`
	Source  string `json:"source" db:"source"`
	Privacy string `json:"privacy" db:"privacy"`
	Name    string `json:"name" db:"name"`
	Email   string `json:"email" db:"email"`
	Phone   string `json:"phone" db:"phone"`
	Phone2  string `json:"phone2" db:"phone2"`
	Address string `json:"address" db:"address"`
	Notes   string `json:"notes" db:"notes"`
}

type Platform struct {
	OldId struct {
		ObjectID string `json:"$oid"`
	} `json:"_id"`
	Website  string            `json:"website" db:"website"`
	Source   string            `json:"source" db:"source"`
	Privacy  string            `json:"privacy" db:"privacy"`
	Name     string            `json:"name" db:"name"`
	Notes    string            `json:"generalNotes" db:"notes"`
	Comment  string            `json:"comment" db:"comment"`
	Country  string            `json:"country" db:"country"`
	Category string            `json:"category"`
	Contacts []ObjectReference `json:"contacts"`
	Projects []ObjectReference `json:"projects"`
	Articles []ObjectReference `json:"articles"`
}

type Project struct {
	OldId struct {
		ObjectID string `json:"$oid"`
	} `json:"_id"`
	Title       string            `json:"title" db:"title"`
	Link        string            `json:"link" db:"link"`
	Description string            `json:"description" db:"description"`
	Date        string            `json:"date" db:"date"`
	Body        string            `json:"body" db:"body"`
	Platforms   []ObjectReference `json:"platforms"`
	Tags        []string          `json:"tags"`
}
