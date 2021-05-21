package webhook

var (
	TelegramMaxLimitRequestResp = `
	{
		"ok": false,
		"error_code": 429,
		"description": "Too Many Requests: retry after 5",
		"parameters": {
		  "retry_after": 5
		}
	}
	`
	TelegramStartResp = `
	{
		"update_id": 107460166,
		"message": {
		  "message_id": 255,
		  "from": {
			"id": 664743441,
			"is_bot": false,
			"first_name": "Tu",
			"last_name": "Dinh",
			"username": "tuacoustic",
			"language_code": "en"
		  },
		  "chat": {
			"id": 664743441,
			"first_name": "Tu",
			"last_name": "Dinh",
			"username": "tuacoustic",
			"type": "private"
		  },
		  "date": 1621411217,
		  "text": "/start",
		  "entities": [
			{
			  "offset": 0,
			  "length": 6,
			  "type": "bot_command"
			}
		  ]
		}
	  }
	`
	TelegramComandResp = `
	{
		"update_id": 107460168,
		"message": {
		  "message_id": 258,
		  "from": {
			"id": 664743441,
			"is_bot": false,
			"first_name": "Tu",
			"last_name": "Dinh",
			"username": "tuacoustic",
			"language_code": "en"
		  },
		  "chat": {
			"id": 664743441,
			"first_name": "Tu",
			"last_name": "Dinh",
			"username": "tuacoustic",
			"type": "private"
		  },
		  "date": 1621411955,
		  "text": "/command",
		  "entities": [
			{
			  "offset": 0,
			  "length": 8,
			  "type": "bot_command"
			}
		  ]
		}
	  }
	`
	TelegramNonCommandResp = `
	{
		"update_id": 107460167,
		"message": {
		  "message_id": 257,
		  "from": {
			"id": 664743441,
			"is_bot": false,
			"first_name": "Tu",
			"last_name": "Dinh",
			"username": "tuacoustic",
			"language_code": "en"
		  },
		  "chat": {
			"id": 664743441,
			"first_name": "Tu",
			"last_name": "Dinh",
			"username": "tuacoustic",
			"type": "private"
		  },
		  "date": 1621411684,
		  "text": "Answer A"
		}
	  }
	`
)
