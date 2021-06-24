package webhook

type enumStudyCommandStruct struct {
	GetCommand   string
	StudyCommand string
}

var (
	EnumStudyCommand = enumStudyCommandStruct{
		GetCommand:   "GET_GROUP",
		StudyCommand: "STUDY_GROUP",
	}
)
