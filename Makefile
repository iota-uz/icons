# Variables
PHOSPHOR_ICONS_URL = https://phosphoricons.com/assets/phosphor-icons.zip
PHOSPHOR_ICONS_ZIP = phosphor-icons.zip
PHOSPHOR_ICONS_DIR = phosphor-icons

# Targets
.PHONY: all download generate clean

all: download generate

deps:
	go mod download

download:
	@echo "Downloading and extracting Phosphor Icons..."
	wget $(PHOSPHOR_ICONS_URL) -O $(PHOSPHOR_ICONS_ZIP)
	mkdir -p $(PHOSPHOR_ICONS_DIR)
	unzip -o $(PHOSPHOR_ICONS_ZIP) -d $(PHOSPHOR_ICONS_DIR)
	rm $(PHOSPHOR_ICONS_ZIP)

generate:
	@echo "Generating templ files..."
	go run cmd/generate/main.go
	templ fmt .
	templ generate
