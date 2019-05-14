GOBIN=go
GOBUILD=${GOBIN} build
GOCLEAN=${GOBIN} clean
GOTEST=${GOBIN} test
#GOGET=${GOBIN} get -u -v

APP=genPrime
                              
all: build
build: 
	$(GOBUILD) -o $(APP) -v
test: 
	$(GOTEST) -v
clean: 
	$(GOCLEAN)
	rm -f $(APP)


#run:
#	$(GOBUILD) -o $(APP) -v ./...
#	./$(APP)
#
#deps:
#	$(GOGET) github.com/markbates/goth
#	$(GOGET) github.com/markbates/pop
#	$(GOGET) github.com/mdempsky/gocode
#	$(GOGET) github.com/uudashr/gopkgs/cmd/gopkgs
#	$(GOGET) github.com/ramya-rao-a/go-outline
#	$(GOGET) github.com/acroca/go-symbols
#	$(GOGET) golang.org/x/tools/cmd/guru
#	$(GOGET) golang.org/x/tools/cmd/gorename
#	$(GOGET) github.com/go-delve/delve/cmd/dlv
#	$(GOGET) github.com/rogpeppe/godef
#	$(GOGET) github.com/sqs/goreturns
#	$(GOGET) golang.org/x/lint/golint
