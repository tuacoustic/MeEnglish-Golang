package webhook

import "fmt"

func QueryExistTelegramID(telegramID uint64) string {
	return fmt.Sprintf(`
select count(id) as count
from telegram_users
where telegram_id = "%d"
`, telegramID)
}

func QueryAllTelegramUsers() string {
	return fmt.Sprintf(`
select count(id) as count
from telegram_users
`)
}
