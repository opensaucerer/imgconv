package imgconv

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestDecodeWrite(t *testing.T) {
	var formats = []string{
		"jpg",
		"png",
		"gif",
		"tif",
		"bmp",
		"webp",
		"pdf",
	}

	for _, i := range formats {
		b, err := os.ReadFile("testdata/video-001." + i)
		if err != nil {
			t.Fatal(err)
		}

		img, err := Decode(bytes.NewBuffer(b))
		if err != nil {
			t.Fatal("Failed to decode", i)
		}

		if err := Write(io.Discard, img, &FormatOption{}); err != nil {
			t.Fatal("Failed to write", i)
		}

		if _, _, err := DecodeConfig(bytes.NewBuffer(b)); err != nil {
			t.Fatal("Failed to decode", i, "config")
		}
	}

	if _, err := Decode(bytes.NewBufferString("Hello")); err == nil {
		t.Fatal("Decode string want error")
	}
}

func TestOpenSaveFromFile(t *testing.T) {
	if _, err := OpenFromPath("/invalid/path"); err == nil {
		t.Error("Open invalid path want error")
	}

	if _, err := OpenFromPath("go.mod"); err == nil {
		t.Error("Open invalid image want error")
	}

	img, err := OpenFromPath("testdata/video-001.png")
	if err != nil {
		t.Fatal("Fail to open image", err)
	}

	if err := Save("/invalid/path", img, defaultFormat); err == nil {
		t.Fatal("Save invalid path want error")
	}

	if err := Save("testdata/tmp", img, defaultFormat); err != nil {
		t.Fatal("Fail to save image", err)
	}
	if err := os.Remove("testdata/tmp"); err != nil {
		t.Fatal(err)
	}
}

func TestOpenSaveFromURL(t *testing.T) {
	if _, err := OpenFromURL("https://storage.googleapis.com/bpxls-original/footer-landscap"); err == nil {
		t.Error("Open invalid url want error")
	}

	if _, err := OpenFromURL("go.mod"); err == nil {
		t.Error("Open invalid image want error")
	}

	img, err := OpenFromURL("https://storage.googleapis.com/bpxls-original/wm-portrait-new.png")
	if err != nil {
		t.Fatal("Fail to open image", err)
	}

	if err := Save("/invalid/path", img, defaultFormat); err == nil {
		t.Fatal("Save invalid path want error")
	}

	if err := Save("testdata/water", img, &FormatOption{Format: PNG}); err != nil {
		t.Fatal("Fail to save image", err)
	}
	if err := os.Remove("testdata/water"); err != nil {
		t.Fatal(err)
	}
}
