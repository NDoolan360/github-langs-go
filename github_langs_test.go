package githublangsgo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLanguage(t *testing.T) {
	// Test name and alias
	for _, name := range []string{"Go", "go", "Golang", "golang"} {
		err := CompareLanguage(name, Language{
			Type:               "programming",
			Color:              "#00ADD8",
			Aliases:            []string{"golang"},
			Extensions:         []string{".go"},
			TmScope:            "source.go",
			AceMode:            "golang",
			CodemirrorMode:     "go",
			CodemirrorMimeType: "text/x-go",
			LanguageID:         132,
		})
		if err != nil {
			t.Error(err)
		}
	}

	// Test Non-existing language
	err := CompareLanguage("Not-A-Real-Language", Language{})
	if err != ErrLanguageNotFound {
		t.Error(err)
	}
}

func CompareLanguage(name string, exepected Language) error {
	if lang, err := GetLanguage(name); err != nil {
		return err
	} else if !reflect.DeepEqual(lang, exepected) {
		return fmt.Errorf("Got %v, expected %v", lang, exepected)
	}

	return nil
}
