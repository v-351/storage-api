# StorageAPI
### Запуск и остановка сервиса
```make up``` -> ```make down```
### Документация
http://localhost:8080/swagger/index.html (отключается в .env)  
Так же в корне есть Postman коллекция  
Методы расширены доп. параметрами для работы с несколькими складами
### Тесты
Для е2е тестирования используется gonkey  
Для получения сабмодуля и компиляции:  
```make prepare-test```  
Запуск тестов, при запущеном сервисе:  
```make run-test```