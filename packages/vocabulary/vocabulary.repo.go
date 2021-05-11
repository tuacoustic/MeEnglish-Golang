package vocabulary

type VocabularyRepository interface {
	AddVocab(OxfordCRUDJSON) (bool, error)
}
