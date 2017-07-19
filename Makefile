plugin:
	CGO_ENABLED=0 GO15VENDOREXPERIMENT=1 go build -ldflags "-s -w -X main.version='`date +%F_%T`'"

docker: plugin
	docker build -t krakonos/drone-hg2 .
	docker push krakonos/drone-hg2

