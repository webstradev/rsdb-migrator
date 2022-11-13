package importer

import (
	"encoding/json"
	"io"
	"os"

	"github.com/webstradev/rsdb-migrator/types"
)

type LoadedData struct {
	Articles  []types.Article
	Contacts  []types.Contact
	Platforms []types.Platform
	Projects  []types.Project
}

func ImportFiles() (*LoadedData, error) {
	// Initialize empty data obejct to easily get the typed empty arrays for the functions below
	data := LoadedData{}

	// Load all the required files and parse the json into the arrays of the data object
	err := LoadAndParseFile("./mongodumps/articles.json", &data.Articles)
	if err != nil {
		return nil, err
	}

	err = LoadAndParseFile("./mongodumps/contacts.json", &data.Contacts)
	if err != nil {
		return nil, err
	}

	err = LoadAndParseFile("./mongodumps/platforms.json", &data.Platforms)
	if err != nil {
		return nil, err
	}

	err = LoadAndParseFile("./mongodumps/projects.json", &data.Projects)
	if err != nil {
		return nil, err
	}

	data.HandleEdgeCases()

	return &data, nil
}

// This function loads in the specified file and Unmarshals it into an array of the specified result type
func LoadAndParseFile[T any](filename string, results *[]T) error {
	// Open the jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return err
	}

	// Close the file when we finish this functions
	defer jsonFile.Close()

	// Read our opened jsonFile as a byte array.
	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	// Unmarshal the json resuls into the passed struct
	err = json.Unmarshal(byteValue, &results)
	if err != nil {
		return err
	}

	return nil
}
