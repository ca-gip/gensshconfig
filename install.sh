#!/bin/bash

if [ "$(uname)" == "Darwin" ]; then
  echo "Downloading Darwin Release"
  curl -s https://api.github.com/repos/ca-gip/gensshconfig/releases/latest \
    | grep browser_download_url \
    | grep darwin_amd64 \
    | cut -d '"' -f 4 \
    | xargs curl -sL -o /usr/local/bin/gensshconfig
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  echo "Downloading Linux Release"
  curl -s https://api.github.com/repos/ca-gip/gensshconfig/releases/latest \
    | grep browser_download_url \
    | grep linux_amd64 \
    | cut -d '"' -f 4 \
    | xargs curl -sL -o /usr/local/bin/gensshconfig
else echo "Unsupported OS" && exit 1
fi

chmod +x /usr/local/bin/gensshconfig
echo "Install done !"