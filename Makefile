
setup:
	chmod +x bin/*
build-docker:
	chmod +x bin/_buildImage.sh
	bin/_buildImage.sh
build-generator:
	cd gen && env GOOS=darwin GOARCH=amd64 go build -o ../generator.darwin cmd/gen/main.go && cd ..;
	cd gen && env GOOS=linux GOARCH=amd64 go build -o ../generator.linux cmd/gen/main.go && cd ..;
test:
	chmod +x bin/_test.sh
	bin/_test.sh
run:
	chmod +x bin/_run.sh
	bin/_run.sh
clean:
	rm generator.darwin && rm generator.linux;
	rm -R outputs/