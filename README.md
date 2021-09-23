# psynder
Tinder for dogs

### run server
```
./build_server.sh
./bin/server
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
### deploy app to a physical device
```
./build_android_app.sh
adb -d install bin/myapplication.apk
```
