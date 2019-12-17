MAKEFLAGS += --silent

all: 
		install

install: 
		cd tasksource && go get -d -v
		cd tasksink && go get -d -v

clean:
		cd tasksource && go clean 
		cd tasksink && go clean

build:
		cd tasksource && go build 
		cd tasksink && go build

.PHONY: install clean build