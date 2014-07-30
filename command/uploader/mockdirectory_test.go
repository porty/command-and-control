package uploader

import (
	"os"
	"testing"
	"time"
)

type MockDirectory struct {
	infos []os.FileInfo
}

func NewMockDirectory() MockDirectory {
	return MockDirectory{
		make([]os.FileInfo, 0, 10),
	}
}

func (d MockDirectory) ReadDir() ([]os.FileInfo, error) {
	return d.infos, nil
}

func (d MockDirectory) Path() string {
	return "/home/ubuntu/images"
}

func (d *MockDirectory) addFile(name string) {
	d.infos = append(d.infos, MockFileInfo{
		name,
		1234,
		0,
		time.Now(),
	})
}

func (d *MockDirectory) addDir(name string) {
	d.infos = append(d.infos, MockFileInfo{
		name,
		0,
		os.ModeDir,
		time.Now(),
	})
}

type MockFileInfo struct {
	name string
	size int64
	mode os.FileMode
	time time.Time
}

func (f MockFileInfo) Name() string {
	return f.name
}

func (f MockFileInfo) Size() int64 {
	return f.size
}

func (f MockFileInfo) Mode() os.FileMode {
	return f.mode
}

func (f MockFileInfo) ModTime() time.Time {
	return f.time
}

func (f MockFileInfo) IsDir() bool {
	return f.mode.IsDir()
}

func (f MockFileInfo) Sys() interface{} {
	return nil
}

func TestMockDirectory(t *testing.T) {
	r := MockDirectory{}
	files, err := r.ReadDir()
	if len(files) != 0 {
		t.Fatal("Files should be empty")
	}
	if err != nil {
		t.Fatal("There should be no error")
	}

	r.addFile("butts")

	files, err = r.ReadDir()
	if len(files) != 1 {
		t.Fatal("Files should have 1")
	}
	if err != nil {
		t.Fatal("There should be no error")
	}
}
