package vocabulary

type VocabularyRepository interface {
	AddVocab(AddVocabRequestStruct) (bool, error)
}
