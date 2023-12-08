

build:
	go run .

public: build
	cp -a books static index.html public/

clean:
	rm -rf public && git checkout public