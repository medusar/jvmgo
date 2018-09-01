package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//通过类名读取文件内容,如果err!=nil表示读取成功，没有找到则返回error
	readClass(className string) ([]byte, Entry, error)
	String() string
}

//根据路径类型来生成相应的Entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, "jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirectoryEntry(path)
}
