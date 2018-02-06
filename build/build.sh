rm -rf static/adm/tmp/*
rm -rf static/adm/log/*
rm -rf bo.bo
GOOS=linux GOARCH=amd64 go build