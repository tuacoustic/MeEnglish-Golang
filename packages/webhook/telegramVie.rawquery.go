package webhook

import "fmt"

func QueryTelegramStudyCommand() string {
	return fmt.Sprintf(`
select count(id) as count
from telegram_study_command
where active = true
`)
}
