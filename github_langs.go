package githublangsgo

import (
	"embed"
	"errors"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

//go:embed languages.toml
var languagesFileContent embed.FS

var languages map[string]Language

type Language struct {
	Type               string   `toml:"type"`
	Color              string   `toml:"color,omitempty"`
	Extensions         []string `toml:"extensions,omitempty"`
	TmScope            string   `toml:"tm_scope"`
	AceMode            string   `toml:"ace_mode"`
	LanguageID         int      `toml:"language_id"`
	Aliases            []string `toml:"aliases,omitempty"`
	CodemirrorMode     string   `toml:"codemirror_mode,omitempty"`
	CodemirrorMimeType string   `toml:"codemirror_mime_type,omitempty"`
	Interpreters       []string `toml:"interpreters,omitempty"`
	Group              string   `toml:"group,omitempty"`
	Filenames          []string `toml:"filenames,omitempty"`
	Wrap               bool     `toml:"wrap,omitempty"`
	FsName             string   `toml:"fs_name,omitempty"`
	Searchable         bool     `toml:"searchable,omitempty"`
}

var (
	ErrReadLanguagesFile  = errors.New("failed to read languages.toml")
	ErrUnmarshalLanguages = errors.New("failed to unmarshal languages.toml")
	ErrLanguageNotFound   = errors.New("language not found")
)

func loadLanguages() error {
	data, err := languagesFileContent.ReadFile("languages.toml")
	if err != nil {
		return ErrReadLanguagesFile
	}

	if err := toml.Unmarshal(data, &languages); err != nil {
		return ErrUnmarshalLanguages
	}

	return nil
}

func GetAllLanguages() (map[string]Language, error) {
	// Load languages if not already loaded
	if languages == nil {
		if err := loadLanguages(); err != nil {
			return map[string]Language{}, err
		}
	}

	return languages, nil
}

func GetLanguage(languageName string) (Language, error) {
	allLanguages, err := GetAllLanguages()

	if err != nil {
		return Language{}, err
	}

	for lang, info := range allLanguages {
		if strings.EqualFold(lang, languageName) {
			return info, nil
		}
		for _, alias := range info.Aliases {
			if strings.EqualFold(alias, languageName) {
				return info, nil
			}
		}
	}

	return Language{}, ErrLanguageNotFound
}
