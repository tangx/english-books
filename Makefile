

build:
	go run .

public: build
	cp -a books static index.html public/