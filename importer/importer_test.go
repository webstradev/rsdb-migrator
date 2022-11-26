package importer

import (
	"reflect"
	"testing"
)

func TestImportFiles(t *testing.T) {
	// Run import files function on testdata
	got, err := ImportFiles("../testdata")
	if err != nil {
		t.Fatalf("ImportFiles() error = %v", err)
	}

	// Compare expected and got data
	if !reflect.DeepEqual(ExpectedData, *got) {
		t.Errorf("ImportFiles() got = %v, want %v", got, ExpectedData)
	}
}
