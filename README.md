**Welcome to me-english**

Details: https://github.com/tudinhacoustic/me-english

Run local source:
1. Create bin folder:
>mkdir bin
2. Make Environment Files:
>pwd

>me-english/

>nano .env
```
API_PORT=4040

```
3. Run source
>go run server.go
4. Run hot reload:
>make hot-reload

Run docker:
1. Build Dockerfile
>make docker-build
2. Init Swarm
>docker swarm init
3. Run Service
>make docker-run-service

Struct API resp:
1. Success
```
{
    "status": number,
    "data": []
}
```
2. Failed
```
{
    "status": number,
    "data": [
        "error_code": string,
        "message": string,
    ]
}
```
Cách đặt tên cache
```
Đặt key: GO_Tên serive/url
Ví dụ: GO_PRODUCT/product/get-all?page=1&limit=100&groupCategoryID=1
(*) Lưu ý: GO là bắt buộc
```