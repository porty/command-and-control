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

// GetBaseName - e.g. "100" for "100.jpg"
func (f FileInfo) GetBaseName() string {
	return f.name
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
	ignores   map[string]bool
}

// NewDirectoryWatcher - Get a new DirectoryWatcher
func NewDirectoryWatcher(dir DirectoryLister /*string*/) DirectoryWatcher {
	closed := make(chan byte)
	close(closed)
	ignores := make(map[string]bool)
	return DirectoryWatcher{
		&dir,
		make(chan FileInfo),
		make([]FileInfo, 0, 100),
		closed,
		0,
		ignores,
	}
}

func (w *DirectoryWatcher) addHandled(fileInfo *FileInfo) {
	w.handled = append(w.handled, *fileInfo)
}

func (w *DirectoryWatcher) setTrigger(c chan byte) {
	w.trigger = c
}

func getPathInfo(s string) (valid bool, shortened string, other string) {
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
		valid, shortened, other := getPathInfo(f.Name())
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

func (w *DirectoryWatcher) watchOneWithIgnores() (ret []FileInfo) {
	ret = make([]FileInfo, 0, 10)
	for _, f := range w.watchOne() {
		if !w.ignores[f.GetBaseName()] {
			ret = append(ret, f)
		}
	}
	return
}

// AddIgnore - Adding "100" will ignore files "100.jpg" and "100.txt"
func (w *DirectoryWatcher) AddIgnore(baseName string) {
	w.ignores[baseName] = true
}

// RemoveIgnore - opposite to AddIgnore
func (w *DirectoryWatcher) RemoveIgnore(baseName string) {
	delete(w.ignores, baseName)
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
