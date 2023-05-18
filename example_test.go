package imgconv_test

import (
	"io"
	"log"

	"github.com/opensaucerer/imgconv"
)

func Example() {
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
