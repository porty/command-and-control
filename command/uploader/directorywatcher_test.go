package uploader

import (
	"testing"
)

func TestSimpleWatchOne(t *testing.T) {
	d := MockDirectory{}
	w := NewDirectoryWatcher(&d)

	files := w.watchOne()
	if len(files) != 0 {
		t.Fatal("Should be empty")
	}

	d.addFile("sup.txt")

	files = w.watchOne()
	if len(files) != 0 {
		t.Fatal("Should be empty")
	}

	d.addFile("sup.jpg")
	files = w.watchOne()
	if len(files) != 1 {
		t.Fatal("Should have 1 file")
	}

	if files[0].GetImagePath() != d.Path()+"/"+"sup.jpg" {
		t.Fatal("Path should be something like sup.jpg")
	}
	if files[0].GetTextPath() != d.Path()+"/"+"sup.txt" {
		t.Fatal("Path should be something like sup.txt")
	}

	files = w.watchOne()
	if len(files) != 1 {
		t.Fatal("Should have 1 file")
	}
}

func TestSimpleWatchOneWithIgnores(t *testing.T) {
	d := MockDirectory{}
	w := NewDirectoryWatcher(&d)

	files := w.watchOneWithIgnores()
	if len(files) != 0 {
		t.Fatal("Should be empty")
	}

	d.addFile("sup.txt")

	files = w.watchOneWithIgnores()
	if len(files) != 0 {
		t.Fatal("Should be empty")
	}

	d.addFile("sup.jpg")
	files = w.watchOneWithIgnores()
	if len(files) != 1 {
		t.Fatal("Should have 1 file")
	}

	if files[0].GetImagePath() != d.Path()+"/"+"sup.jpg" {
		t.Fatal("Path should be something like sup.jpg")
	}
	if files[0].GetTextPath() != d.Path()+"/"+"sup.txt" {
		t.Fatal("Path should be something like sup.txt")
	}

	w.AddIgnore("sup")

	files = w.watchOneWithIgnores()
	if len(files) != 0 {
		t.Fatal("Should have 0 files")
	}

	w.RemoveIgnore("sup")

	files = w.watchOneWithIgnores()
	if len(files) != 1 {
		t.Fatal("Should have 1 file")
	}
}
