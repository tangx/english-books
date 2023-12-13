
build:
	go run .

public: build
	cp -a targets static index.html public/

clean:
	rm -rf public && git checkout public