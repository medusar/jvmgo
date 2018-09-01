package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
)

//jar包类型的Entry,
//读取jar包里的每个文件，从而找到需要的类
type ZipEntry struct {
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	log.Println("zip entry path:" + absPath)
	return &ZipEntry{absPath: absPath}
}

func (z *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(z.absPath)
	if err != nil {
		return nil, nil, err
	}

	defer r.Close()
	for _, f := range r.File {
		log.Println("zip inner file:" + f.Name)
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, z, nil
		}
	}
	//没有读取到返回error
	return nil, nil, errors.New("class not found:" + className)
}

func (z *ZipEntry) String() string {
	return z.absPath
}
