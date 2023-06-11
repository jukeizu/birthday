TAG=$(shell git describe --tags --always)
VERSION=$(TAG:v%=%)
REPO=jukeizu/birthday
GO=GO111MODULE=on go
BUILD=GOARCH=amd64 $(GO) build -ldflags="-s -w -X main.Version=$(VERSION)" 
PROTOFILES=$(wildcard .protobuf/birthday/v1/*.proto)
PROTOPBDEST="../../../api/protobuf-spec"
PBFILES=$(patsubst %.proto,%.pb.go, $(PROTOFILES))

.PHONY: all deps test proto build clean $(PROTOFILES)

all: deps test build 
deps:
	$(GO) mod download

test:
	$(GO) vet ./...
	$(GO) test -v -race ./...

build:
	$(BUILD) -o bin/birthday-$(VERSION) .

build-linux:
	CGO_ENABLED=0 GOOS=linux $(BUILD) -a -installsuffix cgo -o bin/birthday .

docker-build:
	docker build -t $(REPO):$(VERSION) .

docker-deploy:
	docker push $(REPO):$(VERSION)

proto: $(PBFILES)

%.pb.go: %.proto
	cd $(dir $<) && mkdir -p $(PROTOPBDEST)/$(patsubst %.proto,%pb, $(notdir $<)) && protoc $(notdir $<) --go_out=plugins=grpc:$(PROTOPBDEST)/$(patsubst %.proto,%pb, $(notdir $<))

clean:
	@find bin -type f ! -name '*.toml' -delete -print
