#! /bin/sh
# shellcheck disable=SC2077
# shellcheck disable=SC2140
if [ "$(uname)"=="Darwin" ]; then
  cd miniprogram && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
else
  cd miniprogram && go build .
fi

