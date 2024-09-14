BINARY_NAME=pdf-to-images
INSTALL_DIR=/usr/local/bin

build:
	go build -o $(BINARY_NAME) main.go

install: build
	sudo install -m 755 $(BINARY_NAME) $(INSTALL_DIR)

uninstall:
	sudo rm -f $(INSTALL_DIR)/$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

.PHONY: build install uninstall clean
