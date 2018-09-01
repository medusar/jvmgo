package classpath

import (
	"io/ioutil"
	"path/filepath"
)

//目录类型Entry
//直接把目录和类名拼接起来读取文件
type DirEntry struct {
	absDir string
}

func newDirectoryEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

func (e *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(e.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, e, err
}

func (e *DirEntry) String() string {
	return e.absDir
}
