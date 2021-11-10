# psynder
Tinder for dogs


## Generate keys
```
openssl genrsa -out app.rsa 2048
openssl rsa -in app.rsa -pubout > app.rsa.pub
```

### Run server

(Twice)
```
docker-compose build --no-cache
docker-compose up
```

Or
```
docker-compose up --build
```

### Request examples
signup
```
curl -v --insecure https://localhost:443/signup -H 'Content-Type: application/json' -d '{"email":"rediska@yandex-team.ru", "password":"qwerty123"}'
```
login
```
curl -v --insecure https://localhost:443/login -H 'Content-Type: application/json' -d '{"email":"rediska@yandex-team.ru", "password":"qwerty123"}'
```
psyna-info
```
curl -v --insecure https://localhost:443/psyna-info  -H 'Content-Type: application/json'  -d '{"psynId":'1'}'

```

### Deploy app to a physical device
```
./build_android_app.sh
adb -d install bin/myapplication.apk
```
