#!/bin/bash

# url_database="postgres://root:postgres@localhost:5432/meetup_dev_demo?sslmode=disable"
path_relative="$(pwd)"/scripts/migrations
env_file=".env"

generate_data() {
  if [ -f "$env_file" ]; then
    url_database=$(grep "^POSTGRESQL_URL=" "$env_file" | cut -d '=' -f 2- | sed 's/"//g')
    if [ -z "$url_database" ]; then
      echo "A url do banco de dado não está definida no arquivo .env"
    fi
  else
    echo "O arquivo .env não existe"
  fi
}

generate_data

migrate_up() {
  sudo docker run --rm -v "$path_relative":/migrations --network host migrate/migrate -path=/migrations/ -database "$url_database" -verbose up
}

migrate_down() {
  sudo docker run --rm -v "$path_relative":/migrations --network host migrate/migrate -path=/migrations/ -database "$url_database" -verbose down
}

migrate_create() {
  read -p "Digite o nome da migrate: " migrate
  sudo docker-compose run --rm migrate create -ext sql -dir /migrations "$migrate"
}

failure_process() {
  if [ $? -eq 0 ]; then
    echo "Processo realizado com sucesso"
  else
    echo "Houve uma falha no processo"
  fi
}

if [ $# -eq 0 ]; then
  migrate_up
  migrate_down
  migrate_create
else
  case $1 in
  "migrate_up")
    migrate_up
    failure_process
    ;;
  "migrate_down")
    migrate_down
    failure_process
    ;;
  "migrate_create")
    migrate_create
    failure_process
    ;;
  *)
    echo "Função não encontrada"
    ;;
  esac
fi
