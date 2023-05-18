package imgconv

import (
	"image"
	"io"
	"os"

	_ "github.com/sunshineplan/tiff"

	"github.com/disintegration/imaging"
)

type decodeConfig struct {
	autoOrientation bool
}

var defaultDecodeConfig = decodeConfig{
	autoOrientation: true,
}

// DecodeOption sets an optional parameter for the Decode and Open functions.
type DecodeOption func(*decodeConfig)

// AutoOrientation returns a DecodeOption that sets the auto-orientation mode.
// If auto-orientation is enabled, the image will be transformed after decoding
// according to the EXIF orientation tag (if present). By default it's enabled.
func AutoOrientation(enabled bool) DecodeOption {
	return func(c *decodeConfig) {
		c.autoOrientation = enabled
	}
}

// Decode reads an image from r.
// If want to use custom image format packages which were registered in image package, please
// make sure these custom packages imported before importing imgconv package.
func Decode(r io.Reader, opts ...DecodeOption) (image.Image, error) {
	cfg := defaultDecodeConfig
	for _, option := range opts {
		option(&cfg)
	}

	return imaging.Decode(r, imaging.AutoOrientation(cfg.autoOrientation))
}

// DecodeConfig decodes the color model and dimensions of an image that has been encoded in a
// registered format. The string returned is the format name used during format registration.
func DecodeConfig(r io.Reader) (image.Config, string, error) {
	return image.DecodeConfig(r)
}

// OpenFromPath loads an image from a file path.
func OpenFromPath(file string, opts ...DecodeOption) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Decode(f, opts...)
}

// OpenFromURL loads an image for a given url
func OpenFromURL(url string, opts ...DecodeOption) (image.Image, error) {
	f, err := Get(url)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return Decode(f, opts...)
}

// OpenFromFile loads an image from a file. A file is any object implementing the
// io.Reader interface. f is not closed after OpenFromFile returns.
func OpenFromFile(f io.Reader, opts ...DecodeOption) (image.Image, error) {
	return Decode(f, opts...)
}

// Write image according format option
func Write(w io.Writer, base image.Image, option *FormatOption) error {
	return option.Encode(w, base)
}

// SaveToPath saves the image to the path according to the format option
func SaveToPath(output string, base image.Image, option *FormatOption) error {
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer f.Close()

	return option.Encode(f, base)
}

// SaveToWriter saves the image to the writer according to the format option
func SaveToWriter(w io.Writer, base image.Image, option *FormatOption) error {
	return option.Encode(w, base)
}
