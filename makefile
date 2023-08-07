GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

MOD_DIR := $(shell go env GOPATH)/pkg/mod
MCUBE_MODULE := "github.com/infraboard/mcube"
MCUBE_VERSION :=$(shell go list -m ${MCUBE_MODULE} | cut -d' ' -f2)
MCUBE_PKG_PATH := ${MOD_DIR}/${MCUBE_MODULE}@${MCUBE_VERSION}

GV="network:v1alpha1 servicemesh:v1alpha2 tenant:v1alpha1 tenant:v1alpha2 devops:v1alpha1 iam:v1alpha2 devops:v1alpha3 cluster:v1alpha1 storage:v1alpha1 auditing:v1alpha1 types:v1beta1 types:v1beta2 quota:v1alpha2 application:v1alpha1 notification:v2beta1 notification:v2beta2 gateway:v1alpha1 alerting:v2beta1"


ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: init
# init env
init:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/favadi/protoc-go-inject-tag@latest

.PHONY: api
# generate api proto
api:
	protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api \
	       $(API_PROTO_FILES)

.PHONY: pb
# copy mcube protobuf files to third_party/pb
pb:
	@mkdir -pv third_party/mcube/pb
	@sudo cp -r ${MCUBE_PKG_PATH}/pb/* third_party/mcube/pb
	@sudo rm -rf third_party/mcube/pb/*/*.go

.PHONY: run
# run server
run:
	@go run main.go start

.PHONY: inject
# generate inject proto
inject:
	@protoc-go-inject-tag -input=api/*/*.pb.go
	@mcube generate enum -p -m api/*/*.pb.go

.PHONY: set
# Deprecated clientSet cause we will replace code-generate with controller-runtime cache
set:  ;$(info $(M)...Begin to find or download controller-gen.)  @
	./hack/generate_client.sh ${GV}

.PHONY: all
# generate all proto
all:
	make api;
	make inject;

help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
