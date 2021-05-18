package vocabulary

type OxfordRespJSON struct {
	ID       string                 `json:"id"`
	Metadata OxfordMetadataRespJSON `json:"metadata"`
	Results  []OxfordResultRespJSON `json:"results"`
}

type OxfordMetadataRespJSON struct {
	Provider string `json:"provider"`
}

type OxfordResultRespJSON struct {
	Language       string                           `json:"language"`
	LexicalEntries []OxfordResultLexicalEntriesJSON `json:"lexicalEntries"`
}

type OxfordResultLexicalEntriesJSON struct {
	Entries         []OxfordResultLexical2EntriesJSON             `json:"Entries"`
	LexicalCategory OxfordResultLexicalEntriesLexicalCategoryJSON `json:"LexicalCategory"`
}

type OxfordResultLexicalEntriesLexicalCategoryJSON struct {
	ID string `json:"id"`
}

type OxfordResultLexical2EntriesJSON struct {
	Senses         []OxfordResultLexical2EntriesSensesJSON               `json:"senses"`
	Pronunciations []OxfordResultLexical2EntriesSensesPronunciationsJSON `json:"pronunciations"`
}

type OxfordResultLexical2EntriesSensesJSON struct {
	Definitions []string                                        `json:"definitions"`
	Examples    []OxfordResultLexical2EntriesSensesExamplesJSON `json:"examples"`
}

type OxfordResultLexical2EntriesSensesExamplesJSON struct {
	Text string `json:"text"`
}

type OxfordResultLexical2EntriesSensesPronunciationsJSON struct {
	Dialects         []string `json:"dialects"`
	PhoneticSpelling string   `json:"phoneticSpelling"`
	AudioFile        string   `json:"audioFile"`
}

type OxfordCRUDJSON struct {
	ID       string
	Provider string
	Language string
	// Definition []string
	DefinitionNoun      []string
	DefinitionVerb      []string
	DefinitionAdjective []string
	DefinitionAdverb    []string
	DefinitionPhrasal   []string
	// Examples         []OxfordResultLexical2EntriesSensesExamplesJSON
	ExamplesNoun      []OxfordResultLexical2EntriesSensesExamplesJSON
	ExamplesVerb      []OxfordResultLexical2EntriesSensesExamplesJSON
	ExamplesAdjective []OxfordResultLexical2EntriesSensesExamplesJSON
	ExamplesAdverb    []OxfordResultLexical2EntriesSensesExamplesJSON
	ExamplesPhrasal   []OxfordResultLexical2EntriesSensesExamplesJSON
	AudioFile         string
	Dialects          []string
	PhoneticSpelling  string
	LexicalCategory   []string
}
