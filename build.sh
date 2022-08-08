cd server
go build -o ../zavod *.go

service zavod restart
service zavod status