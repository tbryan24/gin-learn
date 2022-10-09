GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	RUN_USER_API=.\service\user\api\main.go
else
	RUN_USER_API=./service/user/api/main.go
endif

.PHONY: run.user.api
# init env
run.user.api:
	@echo $(RUN_USER_API)
	go run $(RUN_USER_API)

.PHONY: run.user.api