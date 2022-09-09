truss:
	truss proto/websitesvc.proto  --pbpkg github.com/mises-id/mises-websitesvc/proto --svcpkg github.com/mises-id/mises-websitesvc --svcout . -v 
run:
	go run cmd/cli/main.go
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/cli/main.go
upload:
	scp ./main mises_backup:/apps/mises-websitesvc/
replace:
	ssh mises_backup "mv /apps/mises-websitesvc/main /apps/mises-websitesvc/mises-websitesvc"
restart:
	ssh mises_backup "sudo supervisorctl restart websitesvc"
deploy: build \
	upload \
	replace 
	