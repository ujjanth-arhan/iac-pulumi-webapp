#!/bin/bash

touch ./tests/integration/.env
echo ENVIRONMENT="PRODUCTION" >> ./tests/integration/.env
echo PSGR_CONNECTION="host=localhost user=postgres password=password sslmode=disable" >> ./tests/integration/.env
echo DB_NAME="cloudcourse" >> ./tests/integration/.env
echo DB_CONNECTION="host=localhost user=postgres password=password sslmode=disable dbname=cloudcourse" >> ./tests/integration/.env
echo USERS_FILE="/opt/users.csv" >> ./tests/integration/.env
echo BCRYPT_COST=10 >> ./tests/integration/.env
echo LOG_FILE="log" >> ./tests/integration/.env
echo STATSD_SERVER="127.0.0.1:8125" >> ./tests/integration/.env
echo STATSD_PREFIX="dev.client" >> ./tests/integration/.env