package webhook

import "fmt"

func QueryTelegramStudyGroupCommand(telegram_id uint64, command string) string {
	return fmt.Sprintf(`
select count(id) as count
from telegram_study_command
where active = true and telegram_id = %d and command = '%s'
order by updated_at desc
`, telegram_id, command)
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

func IncreaseScroreByOne(telegram_id uint64, vocabulary_id uint64) string {
	return fmt.Sprintf(`
UPDATE study_vocab_lists 
SET score = score + 1  
WHERE (telegram_id = %d and vocabulary_id = %d and active = 1)
`, telegram_id, vocabulary_id)
}

func IncreaseScroreByTwo(telegram_id uint64, vocabulary_id uint64) string {
	return fmt.Sprintf(`
UPDATE study_vocab_lists 
SET score = score + 1  
WHERE (telegram_id = %d and vocabulary_id = %d and active = 1)
`, telegram_id, vocabulary_id)
}

// func QueryTelegramStudyGroupCommand(telegram_id uint64) string {
// 	return fmt.Sprintf(`
// select count(id) as count
// from telegram_study_command
// where active = true and telegram_id = %d
// order by updated_at desc
// `, telegram_id)
// }
