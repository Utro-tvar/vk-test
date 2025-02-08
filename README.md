# Тестовое задание для VK

**Задание:**

Написать приложение для регулярного пингования всех контейнеров в клиенте Docker с выводом в таблицу.

Проект состоит из 4 сервисов:
  - frontend - сервис, отвечающий за вывод страницы с таблицей об ip контейнера, времени пинга до него в миллисекундах и датой последнего успешного соединения
  - pinger - регулярно пингует все контейнеры в Docker и отправляет результат на backend 
  - postgres - база данных PostgreSQL
  - backend - сервис, обрабатывающий запросы от pinger и frontend

После запуска в течение 15-20 секунд pinger начнет пинговать контейнеры. Результат будет доступен по адресу http://localhost:3000/

Для сборки и запуска:
1. Изменить .env файлы в папках сервисов:  

   ./pinger/.env:
   ```dotenv
   #in seconds
   PING_PERIOD=300
   ```
   
   ./postgres/.env:
   
   ```dotenv
   POSTGRES_USER=postgres
   POSTGRES_DB=postgres
   POSTGRES_PASSWORD=postgres
   
   #database parameters
   DB_USER=service
   DB_PASS=password
   DB_NAME=service
   ```

   ./backend/.env:  
   здесь должны быть указаны те же значения, что и в ./postgres/.env
   ```dotenv
   #database parameters
   DB_USER=service
   DB_PASS=password
   DB_NAME=service
   ```

   ./frontend/.env:  
   REACT_APP_PERIOD - период обновления информации в таблице
   ```dotenv
   #in milliseconds
   REACT_APP_PERIOD=50000
   ```

2. Выполнить
   
   ```bash
   docker compose up 
   ```


frontend развернут на базе nginx и через него же отправляет запросы на backend. backend и postgres изолированы отдельной сетью, pinger использует сеть хоста.
