package webhook

import "fmt"

func QueryTelegramStudyCommand() string {
	return fmt.Sprintf(`
select count(id) as count
from telegram_study_command
where active = true
`)
}

func QueryCountVocabByGroup(AwlGroupID uint64) string {
	return fmt.Sprintf(`
select count(id) as count
from vocabulary
where awl_group_id = %d	
`, AwlGroupID)
}

func QueryCountAwlGroup() string {
	return fmt.Sprintf(`
select awl_group_id as count
from vocabulary
group by awl_group_id
`)
}
