**1. Handle Limited Message Error**
```
{
  "ok": false,
  "error_code": 429,
  "description": "Too Many Requests: retry after 5",
  "parameters": {
    "retry_after": 5
  }
}
```
**2. Gửi Button**
```
https://api.telegram.org/bot1824373162:AAHrLY0caFSNJVaZI17B7pPxzQ_dw73YRBU/sendMessage?chat_id=664743441&text=Examle: English&reply_markup={"keyboard":[[{"text":"Answer A","callback_data":"budget"},{"text":"Answer B"},{"text":"Answer C"},{"text":"Answer D"}],[{"text":"Quit the answer","callback_data":"budget"}]],"resize_keyboard":true,"one_time_keyboard":true,"selective":true}
```
**3. Gửi Webhook**
```
https://api.telegram.org/bot1824373162:AAHrLY0caFSNJVaZI17B7pPxzQ_dw73YRBU/setWebhook?url=urlParams
```
**4. Start message /start**
```
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
```
**5. Start message non command**
```
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
```
**6. Start message command**
```
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
```