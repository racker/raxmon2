all: deps
	go build .

deps: 
	go get -u github.com/vaughan0/go-ini
	go get -u github.com/codegangsta/cli
	go get -u github.com/rphillips/gorax/identity
	go get -u github.com/rphillips/gorax/monitoring

