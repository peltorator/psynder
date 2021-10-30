# psynder
Tinder for dogs


## Generate keys
```
openssl genrsa -out app.rsa 2048
openssl rsa -in app.rsa -pubout > app.rsa.pub
```

### run server
```
docker-compose build --no-cache
docker-compose up
```
### request examples
signup
```
curl -v --insecure https://localhost:8080/signup -H 'Content-Type: application/json' -d '{"email":"rediska@yandex-team.ru", "password":"qwerty123"}'
```
login
```
curl -v --insecure https://localhost:8080/login -H 'Content-Type: application/json' -d '{"email":"rediska@yandex-team.ru", "password":"qwerty123"}'
```
loadpsynas
```
curl -v --insecure https://localhost:8080/loadpsynas -H "Authorization: Bearer $TOKEN" -H 'Content-Type: application/json' -d '{"count":1}'
```

### deploy app to a physical device
```
./build_android_app.sh
adb -d install bin/myapplication.apk
```
