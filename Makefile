install-deps:
	go get github.com/bwmarrin/discordgo

build:
	mkdir -p .build
	go build -o .build/gobot src/main.go

build-test:
	mkdir -p .build
	go build -o .build/deleteme src/deleteme.go src/module-context.go

run-test: clean build-test
	.build/deleteme

clean:
	rm -r .build