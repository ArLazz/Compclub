
build:
	docker build -t compclub -f Dockerfile . 
run:
	docker run compclub

start: build run

build_tests:
	docker build -t compclub_tests -f Dockerfile.test .

run_tests:
	docker run compclub_tests

tests: build_tests run_tests
