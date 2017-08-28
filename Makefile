install-deps:
	go get github.com/bwmarrin/discordgo

build:
	cd src && go build -o ../.build/gobot

clean:
	rm .build/*