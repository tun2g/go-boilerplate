server:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run main.go server

.PHONY: server