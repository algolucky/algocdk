build:
	go build -o dist/

container:
	docker build . -t algocdk

install:
	go install

pre-commit:
	pre-commit install
	pre-commit run --all

.PHONY: build container install pre-commit
