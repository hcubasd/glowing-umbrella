cd src
go mod download
go build -buildvcs=false -v -o /usr/local/bin/app ./...
app &
until curl -fSs http://localhost:8080/deals; do sleep 1; done
cd ../tests
go mod download
go test ./...
