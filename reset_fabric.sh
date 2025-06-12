#!/bin/bash

# Остановка и удаление всех сервисов
docker-compose down

# Удаление всех связанных томов
for volume in $(docker volume ls -q); do
  if [[ $volume == *"fabric-network"* ]]; then
    docker volume rm $volume
    echo "Removed volume: $volume"
  fi
done

# Удаление всех криптографических материалов
rm -rf ./crypto-config/*
rm -rf ./system-genesis-block/*
rm -rf ./channel-artifacts/*

echo "Fabric environment has been completely reset"