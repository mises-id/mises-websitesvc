truss:
	truss proto/websitesvc.proto  --pbpkg github.com/mises-id/mises-websitesvc/proto --svcpkg github.com/mises-id/mises-websitesvc --svcout . -v 
run:
	go run cmd/cli/main.go