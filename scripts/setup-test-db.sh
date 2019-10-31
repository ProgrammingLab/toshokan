#!/bin/sh

if [ ! -z "$1" ] && [ $1 != '-y' ]; then
    echo "unknown option: $1"
    exit 1
fi

docker-compose exec pg psql -c 'create database toshokan_test;' -U postgres
docker-compose exec pg psql -c 'create extension pgcrypto;' -U postgres toshokan_test

docker-compose build
if [ $? -gt 0 ]; then
    exit $?
fi

docker-compose run --rm ridgepole -c env:TEST_DATABASE_URL --apply --dry-run
if [ $? -gt 0 ]; then
    exit $?
fi

while [ -z "$1" ]; do
    read -p 'Are you sure you want to apply changes? [Y/n]' ans
    case $ans in
        [Y] ) break;;
        [Nn] ) exit;;
        * ) echo "Input [Y/n]";;
    esac
done

docker-compose run --rm ridgepole -c env:TEST_DATABASE_URL --apply
