package importer

import (
	"reflect"
	"testing"

	"github.com/webstradev/rsdb-migrator/types"
)

func TestImportFiles(t *testing.T) {
	expectedData := LoadedData{
		Articles: []types.Article{
			{
				OldId:       types.ObjectReference{ObjectID: "123"},
				Title:       "Title",
				Link:        "http://www.example.com",
				Description: "Some description",
				Date:        "09/09/2017",
				Body:        "This would contain some long body of text",
				Platforms:   []types.ObjectReference{{ObjectID: "abc"}},
				Tags:        []string{"tag1", "tag2"},
			},
			{
				OldId:       types.ObjectReference{ObjectID: "456"},
				Title:       "Title2",
				Link:        "http://www.example.com",
				Description: "Some other description",
				Date:        "09/09/2019",
				Body:        "This would contain some long body of text",
				Platforms:   []types.ObjectReference{},
				Tags:        []string{"tag1"},
			},
		},
		Contacts: []types.Contact{
			{
				OldId:   types.ObjectReference{ObjectID: "a1a"},
				Title:   "Chief Test Officer",
				Source:  "test",
				Privacy: "Private",
				Name:    "Testie McTestFace",
				Email:   "test@company.org",
				Phone2:  "",
				Address: "1 Mock Road, Testville, Testopia",
				Notes:   "",
				Phone:   "000-000-000",
			},
		},
		Platforms: []types.Platform{
			{
				OldId:    types.ObjectReference{ObjectID: "1a1"},
				Website:  "https://www.example.com",
				Source:   "testdata",
				Privacy:  "Private",
				Name:     "Test Platform",
				Notes:    "",
				Comment:  "",
				Country:  "Testopia",
				Category: "Category1",
				Contacts: []types.ObjectReference{{ObjectID: "a1a"}},
				Projects: []types.ObjectReference{},
				Articles: []types.ObjectReference{{ObjectID: "456"}},
			},
		},
		Projects: []types.Project{
			{
				OldId:       types.ObjectReference{ObjectID: "11a"},
				Title:       "Test Project",
				Link:        "https://www.example.com",
				Description: "Test Description",
				Date:        "09/09/2017",
				Body:        "A body would go here",
				Platforms:   []types.ObjectReference{{ObjectID: "1a1"}},
			},
			{
				OldId:       types.ObjectReference{ObjectID: "22a"},
				Title:       "Test Project2",
				Link:        "https://www.example.com",
				Description: "Test Description",
				Date:        "09/09/2019",
				Body:        "A body would go here",
				Platforms:   []types.ObjectReference{{ObjectID: "1a1"}},
			},
		},
	}

	// Run import files function on testdata
	got, err := ImportFiles("../testdata")
	if err != nil {
		t.Fatalf("ImportFiles() error = %v", err)
	}

	// Compare expected and got data
	if !reflect.DeepEqual(expectedData, *got) {
		t.Errorf("ImportFiles() got = %v, want %v", got, expectedData)
	}
}
