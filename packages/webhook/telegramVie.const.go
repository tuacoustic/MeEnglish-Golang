package webhook

type enumStudyCommandStruct struct {
	GetCommand   string
	StudyCommand string
}

type enumAnswerCommandStruct struct {
	ButtonCommand string
	TextCommand   string
}

var (
	EnumStudyCommand = enumStudyCommandStruct{
		GetCommand:   "GET_GROUP",
		StudyCommand: "STUDY_GROUP",
	}
	EnumAnswerCommand = enumAnswerCommandStruct{
		TextCommand: "text",
	}
)
