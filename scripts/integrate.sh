cd src
go mod download
go build -v -o /usr/local/bin/app ./...
app &
curl --retry 10 --retry-connrefused --retry-delay 1 -sf http://localhost:8080/deals > /dev/null
cd ../tests
go mod download
go test ./...
