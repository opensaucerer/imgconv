# Image Converter

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]
[![Go](https://github.com/opensaucerer/imgconv/workflows/Test/badge.svg)][actions]
[![CoverageStatus](https://coveralls.io/repos/github/opensaucerer/imgconv/badge.svg?branch=main&service=github)][coveralls]
[![GoReportCard](https://goreportcard.com/badge/github.com/opensaucerer/imgconv)][goreportcard]

[godev]: https://pkg.go.dev/github.com/opensaucerer/imgconv
[actions]: https://github.com/opensaucerer/imgconv/actions 'GitHub Actions Page'
[coveralls]: https://coveralls.io/github/opensaucerer/imgconv?branch=main
[goreportcard]: https://goreportcard.com/report/github.com/opensaucerer/imgconv

Package imgconv provides basic image processing functions (resize, add watermark, format converter.).

All the image processing functions provided by the package accept any image type that implements `image.Image` interface
as an input, include jpg(jpeg), png, gif, tif(tiff), bmp, webp(decode only) and pdf.

## Installation

    Required go version for building: go1.16 and up
    go get -u github.com/opensaucerer/imgconv

## Documentation

https://pkg.go.dev/github.com/opensaucerer/imgconv

## License

[The MIT License (MIT)](https://raw.githubusercontent.com/opensaucerer/imgconv/main/LICENSE)

## Credits

This repo relies on the following third-party projects:

- [disintegration/imaging](https://github.com/disintegration/imaging)
- [pdfcpu/pdfcpu](https://github.com/pdfcpu/pdfcpu)
- [hhrutter/tiff](https://github.com/hhrutter/tiff)

## Usage examples

A few usage examples can be found below. See the documentation for the full list of supported functions.

### Image resizing

```go
// Resize srcImage to size = 128x128px.
dstImage128 := imgconv.Resize(srcImage, &imgconv.ResizeOption{Width: 128, Height: 128})

// Resize srcImage to width = 800px preserving the aspect ratio.
dstImage800 := imgconv.Resize(srcImage, &imgconv.ResizeOption{Width: 800})

// Resize srcImage to 50% size preserving the aspect ratio.
dstImagePercent50 := imgconv.Resize(srcImage, &imgconv.ResizeOption{Percent: 50})
```

### Add watermark

```go
// srcImage add a watermark at randomly position.
dstImage := imgconv.Watermark(srcImage, &WatermarkOption{Mark: markImage, Opacity: 128, Random: true})

// srcImage add a watermark at fixed position with offset.
dstImage := imgconv.Watermark(srcImage, &WatermarkOption{Mark: markImage, Opacity: 128, Offset: image.Pt(5, 5)})
```

### Format convert

```go
// Convert srcImage to dst with jpg format.
imgconv.Write(dstWriter, srcImage, &imgconv.FormatOption{Format: imgconv.JPEG})
```

## Example code

```go
package main

import (
	"io"
	"log"

	"github.com/opensaucerer/imgconv"
)

func main() {
	// Open a test image.
	src, err := imgconv.OpenFromFile("testdata/video-001.png")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, &imgconv.ResizeOption{Width: 200})

	// Add random watermark set opacity = 128.
	dst := imgconv.Watermark(src, &imgconv.WatermarkOption{Mark: mark, Opacity: 128, Random: true})

	// Write the resulting image as TIFF.
	err = imgconv.Write(io.Discard, dst, &imgconv.FormatOption{Format: imgconv.TIFF})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
}
```
