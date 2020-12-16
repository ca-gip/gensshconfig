#!/bin/bash

checksum () {
   curl -s https://api.github.com/repos/ca-gip/gensshconfig/releases/latest \
  | grep browser_download_url \
  | grep checksums \
  | cut -d '"' -f 4 \
  | xargs curl -sL
}

if [ "$(uname)" == "Darwin" ]; then
  echo "Downloading Darwin Release"
  mkdir -p /var/tmp/gensshconfig
  curl -s https://api.github.com/repos/ca-gip/gensshconfig/releases/latest \
    | grep browser_download_url \
    | grep darwin_amd64 \
    | cut -d '"' -f 4 \
    | xargs curl -sL \
    | tar xf - -C /var/tmp/gensshconfig/
    sudo sh -c 'mv /var/tmp/gensshconfig/gensshconfig /usr/local/bin/ && chmod +x /usr/local/bin/gensshconfig'
    rm -rf /var/tmp/gensshconfig
elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
  echo "Downloading Linux Release"
  mkdir -p /tmp/gensshconfig
  curl -s  https://api.github.com/repos/ca-gip/gensshconfig/releases/latest \
    | grep browser_download_url \
    | grep linux_amd64 \
    | cut -d '"' -f 4 \
    | xargs curl -sL \
    | tar xzf - -C /tmp/gensshconfig
    sudo sh -c 'mv /tmp/gensshconfig/gensshconfig /usr/local/bin/ && chmod +x /usr/local/bin/gensshconfig'
    rm -rf /tmp/gensshconfig
else echo "Unsupported OS" && exit 1
fi

echo "Install done !"