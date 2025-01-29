# doc
- https://www.figma.com/file/Ba8Yxlt23XGKT6VB6akzJG/Untitled?type=design&node-id=60%3A4818&mode=dev&t=2ydAJnVEBM3hQIjC-1
- https://docs.google.com/document/d/192Pq68kyw4Ej0ufcADf1-RnUuZA83kXPSjjCKdANws0/edit
- https://docs.google.com/presentation/d/1F-KwGOIrPEFocuENmtEnSBUJAS6cPkfTqYUrbnDS8Xw/edit#slide=id.p



```nginx
server {
    server_name api-test.ykypcz.vip;
            
    client_max_body_size 150M;
     location / {
    location /api {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://127.0.0.1:7702;
    }

    location /admin {
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_pass http://127.0.0.1:7703;
    }
    listen 80;
}
```