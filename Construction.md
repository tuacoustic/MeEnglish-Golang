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
https://api.telegram.org/bot1824373162:AAHrLY0caFSNJVaZI17B7pPxzQ_dw73YRBU/sendMessage?chat_id=664743441&text=Xin chào&reply_markup={"keyboard":[[{"text":"Test"}]], "resize_keyboard": true, "one_time_keyboard": true}
```
**3. Gủi Webhook**
```
https://api.telegram.org/bot1824373162:AAHrLY0caFSNJVaZI17B7pPxzQ_dw73YRBU/setWebhook?url=urlParams
```
****