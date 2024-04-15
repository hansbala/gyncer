#!/bin/bash

# NOTE: Always run this script from the root repository directory
# Provides easy access to the database for local development

function print_usage() {
    echo "Usage: ./database/local_db_tools.sh [reset|mysql|insert_dummy_user]"
    echo "  reset: Reset the database to its initial state"
    echo "  mysql: Run the mysql command with the given arguments"
    echo "  insert_dummy_user: Insert dummy user data into the database"
}

function reset_db() {
    echo "Resetting database..."
    docker-compose down
    docker kill gyncer-mysql-1
    docker rm gyncer-mysql-1
    docker volume rm gyncer_mysql_data
    docker-compose up -d
    echo "Database reset complete."
}

function run_mysql() {
    # get MYSQL_ROOT_PASSWORD from .env
    MYSQL_ROOT_PASSWORD=$(grep MYSQL_ROOT_PASSWORD .env | cut -d '=' -f2)
    docker-compose exec mysql mysql -u root -p"$MYSQL_ROOT_PASSWORD"
}

function insert_dummy_user() {
    # get MYSQL_ROOT_PASSWORD from .env
    MYSQL_ROOT_PASSWORD=$(grep MYSQL_ROOT_PASSWORD .env | cut -d '=' -f2)
    MYSQL_DATABASE=$(grep MYSQL_DATABASE .env | cut -d '=' -f2)
    docker-compose exec -T mysql mysql -u root -p"$MYSQL_ROOT_PASSWORD" "$MYSQL_DATABASE" < ./database/test_data/insert_dummy_user.sql
}

# check the argument and execute the corresponding function
if [ "$1" == "reset" ]; then
    reset_db
elif [ "$1" == "mysql" ]; then
    run_mysql
elif [ "$1" == "insert_dummy_user" ]; then
    insert_dummy_user
else
    print_usage
fi