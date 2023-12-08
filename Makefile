

build:
	go run .

public: build
	mkdir -p public && cp -a books static index.html public/