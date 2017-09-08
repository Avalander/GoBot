install-deps:
	go get github.com/bwmarrin/discordgo

build:
	go build -o .build/gobot src/main.go

build-test:
	go build -o .build/deleteme src/deleteme.go

run-test: clean build-test
	.build/deleteme

clean:
	rm .build/*