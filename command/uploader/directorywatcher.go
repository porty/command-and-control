package uploader

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

//
// FileInfo
//

// FileInfo - To get paths for a file
type FileInfo struct {
	dir  string
	name string
}

// NewFileInfo - Make a new FileInfo
func NewFileInfo(dir string, name string) FileInfo {
	return FileInfo{dir, name}
}

// GetImagePath - Get full path to jpeg file
func (f FileInfo) GetImagePath() string {
	return path.Join(f.dir, f.name+".jpg")
}

// GetTextPath - Get full path to text file
func (f FileInfo) GetTextPath() string {
	return path.Join(f.dir, f.name+".txt")
}

//
// DirectoryWatcher
//

// DirectoryLister - Thing that lists a directory
type DirectoryLister interface {
	ReadDir() ([]os.FileInfo, error)
	Path() string
}

// Directory - A DirectoryLister for a real directory
type Directory string

// NewDirectory - Create a Directory from a file system path
func NewDirectory(path string) Directory {
	return Directory(path)
}

// ReadDir - see ioutil.ReadDir()
func (d Directory) ReadDir() ([]os.FileInfo, error) {
	return ioutil.ReadDir(d.Path())
}

// Path - The path of this directory
func (d Directory) Path() string {
	return string(d)
}

// DirectoryWatcher - Watches a directory for new files
type DirectoryWatcher struct {
	dir       *DirectoryLister
	FileChan  chan FileInfo
	handled   []FileInfo
	trigger   chan byte
	spincount uint32
}

// NewDirectoryWatcher - Get a new DirectoryWatcher
func NewDirectoryWatcher(dir DirectoryLister /*string*/) DirectoryWatcher {
	closed := make(chan byte)
	close(closed)
	return DirectoryWatcher{
		&dir,
		make(chan FileInfo),
		make([]FileInfo, 0, 100),
		closed,
		0,
	}
}

func (w *DirectoryWatcher) addHandled(fileInfo *FileInfo) {
	w.handled = append(w.handled, *fileInfo)
}

func (w *DirectoryWatcher) setTrigger(c chan byte) {
	w.trigger = c
}

func things(s string) (valid bool, shortened string, other string) {
	if strings.HasSuffix(s, ".jpg") {
		valid = true
		shortened = s[:len(s)-4]
		other = shortened + ".txt"
	} else if strings.HasSuffix(s, ".txt") {
		valid = true
		shortened = s[:len(s)-4]
		other = shortened + ".jpg"
	}
	return
}

func (w *DirectoryWatcher) watchOne() (ret []FileInfo) {

	ret = make([]FileInfo, 0, 10)
	memory := make(map[string]bool)

	//raw, err := ioutil.ReadDir(w.dir)
	raw, err := (*w.dir).ReadDir()

	if err != nil {
		panic(err)
	}

	for _, f := range raw {
		valid, shortened, other := things(f.Name())
		if valid {
			if memory[other] {
				ret = append(ret, NewFileInfo((*w.dir).Path(), shortened))
				delete(memory, other)
			} else {
				memory[f.Name()] = true
			}
		}
	}

	return
}

// Watch - Start watching that directory for new files - never exits
func (w *DirectoryWatcher) Watch() {
	for {
		w.spincount++
		for _, f := range w.watchOne() {
			w.FileChan <- f
		}
		<-w.trigger
	}
}
