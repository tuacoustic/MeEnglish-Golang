package webhook

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"me-english/entities"
	"me-english/utils/config"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type textHandlingStruct struct {
	BotCommand     string
	GetStudyNowVie string
	AutoRemindVie  string
	InstructionVie string
	SupportVie     string
	DevelopVie     string
	DonateVie      string
	BackHome       string
	Continue       string
	QueryGroup     string
	QueryPage      string
	OnCurrentGroup string
	OnCurrentPage  string
	GroupStudy     string
	StartBot       string
	GetAudio       string
	GetImage       string
}

type commandGetGroupStruct struct {
	Group1  string
	Group2  string
	Group3  string
	Group4  string
	Group5  string
	Group6  string
	Group7  string
	Group8  string
	Group9  string
	Group10 string
}

var (
	HomeButtonText = "Trang Chủ"
	msg            = ""
	Limit_GetVocab = 15
	telegramParams = config.SendTelegramMsgStruct{
		ChatID:      uint64(664743441),
		Text:        "Lỗi",
		ReplyMarkup: "",
		ParseMode:   "markdown",
	}
	Command_Handling = textHandlingStruct{
		BotCommand:     "/",
		GetStudyNowVie: "học ngay",
		AutoRemindVie:  "nhắc học tự động",
		InstructionVie: "từ vựng đã lưu",
		SupportVie:     "gửi hỗ trợ",
		DevelopVie:     "cùng phát triển",
		DonateVie:      "ủng hộ tác giả",
		BackHome:       "trang chủ",
		Continue:       "tiếp tục",
		QueryGroup:     "gr",
		QueryPage:      "pg",
		OnCurrentGroup: ">gr",
		OnCurrentPage:  ">pg",
		GroupStudy:     "học theo group",
		StartBot:       "/start",
		GetAudio:       "/audio@",
		GetImage:       "/image@",
	}
	Home_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Học ngay"),
			tgbotapi.NewKeyboardButton("Nhắc học tự động"),
			tgbotapi.NewKeyboardButton("Từ vựng đã lưu"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Gửi hỗ trợ"),
			tgbotapi.NewKeyboardButton("Cùng phát triển"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ủng hộ tác giả"),
		),
	)
	Back_Home_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Tiếp tục"),
			tgbotapi.NewKeyboardButton("Nhắc học tự động"),
			tgbotapi.NewKeyboardButton("Từ vựng đã lưu"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Gửi hỗ trợ"),
			tgbotapi.NewKeyboardButton("Cùng phát triển"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Ủng hộ tác giả"),
		),
	)
	StudyNowVie_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Học Group 1"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("﹒1﹒"),
			tgbotapi.NewKeyboardButton("2 >"),
			tgbotapi.NewKeyboardButton("3 >"),
			tgbotapi.NewKeyboardButton("4 >"),
			tgbotapi.NewKeyboardButton("10 >>"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("﹒Gr1﹒"),
			tgbotapi.NewKeyboardButton("Gr2 >"),
			tgbotapi.NewKeyboardButton("Gr3 >"),
			tgbotapi.NewKeyboardButton("Gr10 >>"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(HomeButtonText),
		),
	)
	AnswerKey_Reply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Trang chủ"),
			tgbotapi.NewKeyboardButton("Học Group khác"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Answer A"),
			tgbotapi.NewKeyboardButton("Answer B"),
			tgbotapi.NewKeyboardButton("Answer C"),
			tgbotapi.NewKeyboardButton("Answer D"),
		),
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton("Quit the answer"),
		),
	)
	Command_GetGroup = "GET_GROUP"
	Donate_Text      = fmt.Sprintf(`
*Thông tin tài khoản*

*Nội địa*
ACB: 4680167 (DINH NGUYEN CAM TU)

*Visa*
Sacombank: 0602 3234 6739 (DINH NGUYEN CAM TU)

*Ví điện tử*
Momo: 0975089502 (DINH NGUYEN CAM TU) 
`)
	Support_Text = fmt.Sprintf(`
*Bạn vui lòng gửi nội dung, hình ảnh lỗi cho thông tin bên dưới*

Email: tudinhacoustic@gmail.com 
Website: tudinh.vn/support
`)
)

func GetStudyNowVie(AwlGroupID uint64, pageNumber uint64, groupVocab string) string {
	return fmt.Sprintf(`
*Group %d - Page %d*

Lưu ý: Các bạn có thể click bên cạnh từ vựng để thấy chi tiết từ vựng đó nhé. Mỗi trang sẽ chứa 15 từ vựng (có thể ít hơn) theo Group, chúc các bạn học tập hiệu quả 😉

%s

*Tip theo nút*:
1. Học theo Group đang xem
2. Trang từ vựng theo Group
3. Trang Group
4. Trở về Trang chính
`, AwlGroupID, pageNumber, groupVocab)
}

func GetStudyNowNullVocabVie(AwlGroupID uint64) string {
	return fmt.Sprintf(`
*Group %d*
Hiện *Group %d* chưa có từ vựng, Bạn vui lòng thử lại sau nhé 😘
`, AwlGroupID, AwlGroupID)
}

func StudyNowVieReplyDefault(AwlGroupID uint64) tgbotapi.ReplyKeyboardMarkup {
	var respReply tgbotapi.ReplyKeyboardMarkup
	var paginateGroupButton1to5 []tgbotapi.KeyboardButton
	var paginateGroupButton6to10 []tgbotapi.KeyboardButton

	for indexGroup1to5 := 1; indexGroup1to5 <= 5; indexGroup1to5++ {
		if indexGroup1to5 == int(AwlGroupID) {
			paginateGroupButton1to5 = append(paginateGroupButton1to5, tgbotapi.NewKeyboardButton(fmt.Sprintf(">Gr%d", indexGroup1to5)))
			continue
		}
		paginateGroupButton1to5 = append(paginateGroupButton1to5, tgbotapi.NewKeyboardButton(fmt.Sprintf("Gr%d", indexGroup1to5)))
	}
	for indexGroup6to10 := 6; indexGroup6to10 <= 10; indexGroup6to10++ {
		if indexGroup6to10 == int(AwlGroupID) {
			paginateGroupButton6to10 = append(paginateGroupButton6to10, tgbotapi.NewKeyboardButton(fmt.Sprintf(">Gr%d", indexGroup6to10)))
			continue
		}
		paginateGroupButton6to10 = append(paginateGroupButton6to10, tgbotapi.NewKeyboardButton(fmt.Sprintf("Gr%d", indexGroup6to10)))
	}
	respReply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(HomeButtonText),
		),
		paginateGroupButton1to5,
		paginateGroupButton6to10,
	)
	return respReply
}

func StudyNowVieReply(AwlGroupID uint64, currentPage uint64, pagination uint32) tgbotapi.ReplyKeyboardMarkup {
	var respReply tgbotapi.ReplyKeyboardMarkup
	var paginateLearnNowButton []tgbotapi.KeyboardButton
	var paginateGroupButton1to5 []tgbotapi.KeyboardButton
	var paginateGroupButton6to10 []tgbotapi.KeyboardButton
	rollGroup := fmt.Sprintf("Học theo Group %d", AwlGroupID)
	// Theo paginate
	var maxPaginationNumber int = int(pagination / uint32(Limit_GetVocab))
	if float64(pagination)/float64(Limit_GetVocab) > float64(maxPaginationNumber) {
		maxPaginationNumber = maxPaginationNumber + 1
	}
	for indexPage := 1; indexPage <= maxPaginationNumber; indexPage++ {
		if indexPage == int(currentPage) {
			paginateLearnNowButton = append(paginateLearnNowButton, tgbotapi.NewKeyboardButton(fmt.Sprintf(">Pg%d", indexPage)))
			continue
		}
		paginateLearnNowButton = append(paginateLearnNowButton, tgbotapi.NewKeyboardButton(fmt.Sprintf("Pg%d", indexPage)))
	}
	for indexGroup1to5 := 1; indexGroup1to5 <= 5; indexGroup1to5++ {
		if indexGroup1to5 == int(AwlGroupID) {
			paginateGroupButton1to5 = append(paginateGroupButton1to5, tgbotapi.NewKeyboardButton(fmt.Sprintf(">Gr%d", indexGroup1to5)))
			continue
		}
		paginateGroupButton1to5 = append(paginateGroupButton1to5, tgbotapi.NewKeyboardButton(fmt.Sprintf("Gr%d", indexGroup1to5)))
	}
	for indexGroup6to10 := 6; indexGroup6to10 <= 10; indexGroup6to10++ {
		if indexGroup6to10 == int(AwlGroupID) {
			paginateGroupButton6to10 = append(paginateGroupButton6to10, tgbotapi.NewKeyboardButton(fmt.Sprintf(">Gr%d", indexGroup6to10)))
			continue
		}
		paginateGroupButton6to10 = append(paginateGroupButton6to10, tgbotapi.NewKeyboardButton(fmt.Sprintf("Gr%d", indexGroup6to10)))
	}
	respReply = tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(HomeButtonText),
			tgbotapi.NewKeyboardButton(rollGroup),
		),
		paginateLearnNowButton,
		paginateGroupButton1to5,
		paginateGroupButton6to10,
	)
	return respReply
}

func VocabDetailText(AwlGroupID uint64, vocabData entities.FindVocab, listsVocabArr []string) string {
	// groupText := fmt.Sprintf("```GROUP %d```", AwlGroupID)
	var lexicalCategoryArr []string
	json.Unmarshal([]byte(vocabData.LexicalCategory), &lexicalCategoryArr)
	definition := getDefinition(vocabData)
	example := getExample(vocabData)
	textArrToString := strings.Join(listsVocabArr, "")
	encodeBase64 := b64.StdEncoding.EncodeToString([]byte(strings.ToLower(vocabData.Word)))
	return fmt.Sprintf(`
*GROUP %d*

🔑 *%s* (%s) (%s): %s
Audio: /audio@%s﹒
Image: /image@%s﹒

*Definition*
%s

*Example*
%s

%s
`, AwlGroupID, strings.Title(strings.ToLower(vocabData.Word)), vocabData.PhoneticSpelling, strings.Join(lexicalCategoryArr, "|"), vocabData.Vi, vocabData.Word, encodeBase64, strings.Join(definition, "\n"), strings.Join(example, "\n"), textArrToString)
}

func getDefinition(defiData entities.FindVocab) []string {
	var definitionArr []string
	var definitionNoun []string
	var definitionVerb []string
	var definitionAdjective []string
	var definitionAdverb []string
	var definitionPhrasal []string
	if defiData.DefinitionNoun != "" {
		json.Unmarshal([]byte(defiData.DefinitionNoun), &definitionNoun)
		definitionArr = append(definitionArr, fmt.Sprintf("Noun: %s", strings.Join(definitionNoun, " | ")))
	}
	if defiData.DefinitionVerb != "" {
		json.Unmarshal([]byte(defiData.DefinitionVerb), &definitionVerb)
		definitionArr = append(definitionArr, fmt.Sprintf("Verb: %s", strings.Join(definitionVerb, " | ")))
	}
	if defiData.DefinitionAdjective != "" {
		json.Unmarshal([]byte(defiData.DefinitionAdjective), &definitionAdjective)
		definitionArr = append(definitionArr, fmt.Sprintf("Adjective: %s", strings.Join(definitionAdjective, " | ")))
	}
	if defiData.DefinitionAdverb != "" {
		json.Unmarshal([]byte(defiData.DefinitionAdverb), &definitionAdverb)
		definitionArr = append(definitionArr, fmt.Sprintf("Adverb: %s", strings.Join(definitionAdverb, " | ")))
	}
	if defiData.DefinitionPhrasal != "" {
		json.Unmarshal([]byte(defiData.DefinitionPhrasal), &definitionPhrasal)
		definitionArr = append(definitionArr, fmt.Sprintf("Phrasal: %s", strings.Join(definitionPhrasal, " | ")))
	}
	return definitionArr
}

func getExample(exData entities.FindVocab) []string {
	var exampleArr []string
	var exampleNoun []string
	var exampleVerb []string
	var exampleAdjective []string
	var exampleAdverb []string
	var examplePhrasal []string
	if exData.ExampleNoun != "" {
		json.Unmarshal([]byte(exData.ExampleNoun), &exampleNoun)
		exampleArr = append(exampleArr, fmt.Sprintf("Noun: %s", strings.Join(exampleNoun, " | ")))
	}
	if exData.ExampleVerb != "" {
		json.Unmarshal([]byte(exData.ExampleVerb), &exampleVerb)
		exampleArr = append(exampleArr, fmt.Sprintf("Verb: %s", strings.Join(exampleVerb, " | ")))
	}
	if exData.ExampleAdjective != "" {
		json.Unmarshal([]byte(exData.ExampleAdjective), &exampleAdjective)
		exampleArr = append(exampleArr, fmt.Sprintf("Adjective: %s", strings.Join(exampleAdjective, " | ")))
	}
	if exData.ExampleAdverb != "" {
		json.Unmarshal([]byte(exData.ExampleAdverb), &exampleAdverb)
		exampleArr = append(exampleArr, fmt.Sprintf("Adverb: %s", strings.Join(exampleAdverb, " | ")))
	}
	if exData.ExamplePhrasal != "" {
		json.Unmarshal([]byte(exData.ExamplePhrasal), &examplePhrasal)
		exampleArr = append(exampleArr, fmt.Sprintf("Phrasal: %s", strings.Join(examplePhrasal, " | ")))
	}
	return exampleArr
}

func VocabAnswerLists(AwlGroupID uint64, vocabData entities.FindVocab, answerKeyLists []string) string {
	var lexicalCategoryArr []string
	json.Unmarshal([]byte(vocabData.LexicalCategory), &lexicalCategoryArr)
	definition := getDefinition(vocabData)
	example := getExample(vocabData)
	encodeBase64 := b64.StdEncoding.EncodeToString([]byte(strings.ToLower(vocabData.Word)))
	var addTheAnswer []string
	key := []string{"A", "B", "C", "D"}
	for index, value := range answerKeyLists {
		addTheAnswer = append(addTheAnswer, fmt.Sprintf("%s. %s", key[index], value))
	}
	return fmt.Sprintf(`
*📌 Bạn đang học Group %d*

Vui lòng cung cấp đáp án dưới:
🔑 ----- (##) (##): %s
Image: /image@%s

*Definition*
%s

*Example*
%s

%s
`, AwlGroupID, vocabData.Vi, encodeBase64, strings.Join(definition, "\n"), strings.Join(example, "\n"), strings.Join(addTheAnswer, "\n"))
}

// func AnswerKeyReply() tgbotapi.ReplyKeyboardMarkup {

// }
