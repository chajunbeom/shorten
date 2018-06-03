default: clean deps init config build

config=config.json

# set variables for base information
output=bin
app=shorten

# set variables for dependencies
depsfile=gitdeps
depsvar=$(shell cat ${depsfile})

##################################################################
clean:
	# clean out_dir
	rm -rf $(output)

deps:
	go get -v ./...

init:
	@if [ -d $(output) ]; then rm -rf $(output); fi;
	@if [ ! -d $(output) ]; then mkdir -p $(output); mkdir -p $(output)/logs; fi;

config: init
	# copy the configuration files to the output directory.
	cp etc/conf/config.json $(output)/config.json
	cp -r etc/conf/docs $(output)/docs

build: init
	go build -o $(output)/$(app) cmd/main.go

run: stop
	cd bin && ./$(app) -c=config.json &

stop:
	- killall -9 $(app)