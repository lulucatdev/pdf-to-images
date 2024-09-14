# PDF to Images Converter

This is a command-line tool that converts PDF files to high-quality JPEG images. It's written in Go and uses the `pdftoppm` utility for the conversion process.

## Features

- Converts PDF files to JPEG images
- Outputs high-quality images (300 DPI, 100% quality)
- Creates a separate directory for each converted PDF
- Renames output files to match page numbers

## Prerequisites

Before using this tool, make sure you have the following installed:

- Go (version 1.22.3 or later)
- `pdftoppm` utility (usually part of the Poppler package)

## Installation

1. Clone this repository:
   ```
   git clone https://github.com/yourusername/pdf-to-images.git
   cd pdf-to-images
   ```

2. Build and install the binary:
   ```
   make install
   ```

This will compile the program and install it to `/usr/local/bin`.

## Usage

To convert a PDF file to images, use the following command:

```
pdf-to-images <path-to-pdf-file>
```

For example:

```
pdf-to-images /path/to/your/document.pdf
```

The program will create a new directory named after the PDF file (without the .pdf extension) in the same location as the PDF. The converted images will be saved in this directory as JPEG files, named according to their page numbers.

## Uninstallation

To uninstall the tool, run:

```
make uninstall
```

## Development

The main logic of the program is in the `main.go` file:

```go:main.go
startLine: 1
endLine: 81
```

To build the program without installing it, run:

```
make build
```

To clean up build artifacts, run:

```
make clean
```

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.