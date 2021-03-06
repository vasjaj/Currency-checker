# Currency-checker

App that checks currency exchange rates via https://www.bank.lv/vk/ecb_rss.xml

App consits of two main functions - endpoint, data population

### Endpoint

Has two routes:

- / - shows the newest data for all currencies
- /:currency - show historical data for some currency, example - /AUD

### Population

Downloads data from https://www.bank.lv/vk/ecb_rss.xml and saves it to database

## Configuration

Application requires two env files for database and endpoit:

- ./configs/db.env : MYSQL_ROOT_PASSWORD, MYSQL_DATABASE, MYSQL_USER, MYSQL_PASSWORD
- ./configs/web.env : DB_ROOT_PASSWORD, DB_DATABASE, DB_USER, DB_PASSWORD, DB_HOST

Note that USERNAME, PASSWORD and DATABASE should be the same for both config files

## Installation

- Clone app
- To start app, run ```docker-compose up -d```, but first time you will probably get 500 Error, because data is not loaded
- If data has not been downloaded or you want to get latest changes:
1. Compile app via ```go build main.go```
2. Run ```DB_DATABASE={db_name_goes_here} DB_USER={db_user_goes_here} DB_PASSWORD={db_password_goes_here} DB_HOST={db_host_goes_here} ./main populate ```

## Installation scenario example

1. ```git clone https://github.com/vasjaj/Currency-checker.git```
2. ```cd Currency-checker```
3. ```echo "MYSQL_ROOT_PASSWORD=password\nMYSQL_DATABASE=cdb\nMYSQL_USER=user\nMYSQL_PASSWORD=password" > ./configs/db.env```
4. ```echo "DB_ROOT_PASSWORD=password\nDB_DATABASE=cdb\nDB_USER=user\nDB_PASSWORD=password\nDB_HOST=mariadb" > ./configs/web.env```
5. ```docker-compose up -d```
6. ```go build main.go```
7. ```DB_ROOT_PASSWORD=password DB_DATABASE=cdb DB_USER=user DB_PASSWORD=password DB_HOST=127.0.0.1 ./main populate```
8. Note that DB container doesnt work alread after the start, so some errors may occur, but will disappear in as short time
9. Check localhost:8080 and localhost:8080/AUD
