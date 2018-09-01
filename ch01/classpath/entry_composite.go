package classpath

import (
	"errors"
	"strings"
)

//这个类型定义为数组
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

//遍历所有的Entry，指导有一个能加载到为止
func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found:" + className)
}

func (c CompositeEntry) String() string {
	//数组还是slice
	strs := make([]string, len(c))
	for i, entry := range c {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
