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

	sub, err := kbc.ListenForNewTextMessages()
	if err != nil {
		fail("Error listening: %s", err.Error())
	}

	for {
		msg, err := sub.Read()
		if err != nil {
			fail("failed to read message: %s", err.Error())
		}

		if msg.Message.Content.TypeName != "text" {
			continue
		}

		fmt.Println("received: ", msg.Message.Content.Text.Body)

		// uncomment to send chat messages in the original message's channel
		/*
			if _, err = kbc.SendMessage(msg.Message.Channel, msg.Message.Content.Text.Body); err != nil {
				fail("error echo'ing message: %s", err.Error())
			}
		*/
	}

}
