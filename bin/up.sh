#!/bin/bash

set -u

## down
echo "docker-compose down"
docker-compose down
if [ $? -ne 0 ]; then
  echo "fail to down docker host"
  exit 1
fi

echo
echo "ok"

## up mysql
echo
echo "up mysql"
docker-compose up -d db
if [ $? -ne 0 ]; then
  echo "fail to up DB"
  exit 1
fi

## check to up mysql
docker-compose exec db mysql -u root -phoge -e "quit"
while [ $? -ne 0 ]
do
  sleep 1
  docker-compose exec db mysql -u root -phoge -e "quit"
done
echo
echo "ok"

## migrate
echo
echo "migrate DB"
migrate -source file://migrations -database "mysql://root:hoge@tcp(localhost:3306)/mah_jong" up
if [ $? -ne 0 ]; then
  echo "fail to migrate DB"
  echo "down docker host"
  docker-compose down
  exit 1
fi

migrate -source file://migrations -database "mysql://root:hoge@tcp(localhost:3306)/mah_jong_test" up
if [ $? -ne 0 ]; then
  echo "fail to migrate DB"
  echo "down docker host"
  docker-compose down
  exit 1
fi

echo
echo "ok"

## insert data
echo
echo "insert data"

docker-compose exec -T db mysql -u root -phoge < ./rdb/data.sql

echo
echo "ok"

# up app
echo
echo "up app"
docker-compose up -d app
if [ $? -ne 0 ]; then
  echo "fail to up app"
  echo "down docker host"
  docker-compose down
  exit 1
fi

echo
echo "done"
