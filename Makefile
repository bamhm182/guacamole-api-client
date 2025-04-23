# Import environment file
ifeq (,$(strip $(GUACAMOLE_URL)))
	include .env
	export $(shell sed 's/=.*//' .env)
else ifeq (,$(strip $(GUACAMOLE_USERNAME)))
	include .env
	export $(shell sed 's/=.*//' .env)
else ifeq (,$(strip $(GUACAMOLE_PASSWORD)))
	include .env
	export $(shell sed 's/=.*//' .env)
endif

.PHONY:	lint test

all: lint test

lint:
	go vet ./guacamole
	go fmt ./guacamole

test: lint
	go test -count=1 -v -cover --race -tags="unittests" ./guacamole

test_specific: lint
	go test -count=1 -v -cover --race -tags="specific" ./guacamole
