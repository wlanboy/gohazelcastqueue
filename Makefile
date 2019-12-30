MAKEFLAGS += --silent

all: 
		install

install: 
		cd tasksource && go get -d -v
		cd tasksink && go get -d -v
		cd taskmapadder && go get -d -v
		cd taskmaplistener && go get -d -v

clean:
		cd tasksource && go clean 
		cd tasksink && go clean
		cd taskmapadder && go clean
		cd taskmaplistener && go clean

build:
		cd tasksource && go build 
		cd tasksink && go build
		cd taskmapadder && go build
		cd taskmaplistener && go build

.PHONY: install clean build