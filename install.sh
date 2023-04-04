#!/bin/bash

# Determine the user's operating system
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
  OS="linux"
elif [[ "$OSTYPE" == "darwin"* ]]; then
  OS="macos"
elif [[ "$OSTYPE" == "win32" ]]; then
  OS="windows"
else
  echo "Unsupported operating system: $OSTYPE"
  exit 1
fi

# Determine the latest release version
RELEASE=$(curl -s https://api.github.com/repos/ratnadeep007/dev-gpt/releases/latest | grep "tag_name" | cut -d'"' -f4)

# Download the appropriate executable for the user's operating system
URL="https://github.com/ratnadeep007/dev-gpt/releases/download/$RELEASE/myapp-$OS"
curl -L -o myapp "$URL"

# Make the executable executable
chmod +x myapp

# Move the executable to a directory in the user's PATH
sudo mv devgpt /usr/local/bin/
