package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/khulnasoft-bot/khulnasoft-bot/khulnasoft-chat-bot/kschat"
)

func fail(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(3)
}

func main() {
	var kbLoc string
	var kbc *kschat.API
	var err error

	flag.StringVar(&kbLoc, "khulnasoft", "khulnasoft", "the location of the Khulnasoft app")
	flag.Parse()

	if kbc, err = kschat.Start(kschat.RunOptions{KhulnasoftLocation: kbLoc}); err != nil {
		fail("Error creating API: %s", err.Error())
	}

	tlfName := fmt.Sprintf("%s,%s", kbc.GetUsername(), kbc.GetUsername())
	fmt.Printf("saying hello on conversation: %s\n", tlfName)
	if _, err = kbc.SendMessageByTlfName(tlfName, "hello! a bot sent this message."); err != nil {
		fail("Error sending message; %s", err.Error())
	}

}
