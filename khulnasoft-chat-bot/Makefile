PROTOCOL_PATH=../client/protocol
AVDLC=$(PROTOCOL_PATH)/node_modules/.bin/avdlc

.DEFAULT_GOAL := types

types:
	@mkdir -p kschat/types/khulnasoft1
	@mkdir -p kschat/types/gregor1
	@mkdir -p kschat/types/chat1
	@mkdir -p kschat/types/stellar1
	$(AVDLC) -b -l go -t -m -o kschat/types/khulnasoft1 $(PROTOCOL_PATH)/avdl/khulnasoft1/*.avdl
	$(AVDLC) -b -l go -t -m -o kschat/types/gregor1 $(PROTOCOL_PATH)/avdl/gregor1/*.avdl
	$(AVDLC) -b -l go -t -m -o kschat/types/chat1 $(PROTOCOL_PATH)/avdl/chat1/*.avdl
	$(AVDLC) -b -l go -t -m -o kschat/types/stellar1 $(PROTOCOL_PATH)/avdl/stellar1/*.avdl
	goimports -w ./kschat/types/

clean:
	rm -rf kschat/types/
