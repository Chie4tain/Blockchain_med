#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Usage: $0 <path1> <path2> ..."
  exit 1
fi

# Определяем конфигурационный файл оболочки пользователя
USER_SHELL=$(basename "$SHELL")
SHELL_RC="$HOME/.bashrc"
[ "$USER_SHELL" = "zsh" ] && SHELL_RC="$HOME/.zshrc"

# Обрабатываем каждый путь
for path_arg in "$@"; do
  # Конвертируем относительный путь в абсолютный
  if [ -d "$path_arg" ]; then
    path=$(cd "$path_arg" && pwd)
  else
    echo "Warning: Directory '$path_arg' does not exist!" >&2
    continue
  fi

  # Экранирование спецсимволов для регулярных выражений
  escaped_path=$(printf '%q' "$path" | sed 's/[][\.*^$(){}?+|/]/\\&/g')
  
  # Проверка существования пути в конфигурации
  if grep -Eq "export PATH=.*$escaped_path" "$SHELL_RC"; then
    echo "Already exists: $path"
  else
    # Добавляем путь в начало PATH для приоритета
    echo "export PATH=\"$path:\$PATH\"" >> "$SHELL_RC"
    echo "Added: $path"
  fi
done

echo "Paths updated. To apply changes:"
echo "1. Restart your terminal"
echo "2. Or run: source $SHELL_RC"