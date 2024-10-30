Команды для проверки:
curl -X GET http://localhost:8080/admin -H "Authorization: your_jwt_token_here"
curl -X POST http://localhost:8080/protected -H "Authorization: your_jwt_token_here" -H "X-CSRF-Token: your_csrf_token_here"
# Создайте закрытый ключ для CA (корневой сертификат)
openssl genpkey -algorithm RSA -out ca.key -aes256
# Убедитесь, что запомнили пароль, так как он понадобится для подписи сертификатов

# Создайте самоподписанный сертификат CA, действительный на 365 дней
openssl req -x509 -new -key ca.key -sha256 -days 365 -out ca.crt
# Создайте закрытый ключ сервера
openssl genpkey -algorithm RSA -out server.key

# Создайте запрос на сертификат сервера
openssl req -new -key server.key -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 365 -sha256
# Создайте закрытый ключ для клиента
openssl genpkey -algorithm RSA -out client.key

# Создайте запрос на сертификат клиента
openssl req -new -key client.key -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 365 -sha256
