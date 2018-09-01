package classpath

import (
	"log"
	"os"
	"path/filepath"
)

// classpath entry
type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

//解析jre路径和classpath路径
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (c *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")

	log.Println("jre lib path:" + jreLibPath)
	c.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	log.Println("ext lib path:" + jreExtPath)
	c.extClasspath = newWildcardEntry(jreExtPath)
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}

	panic("Can not find jre folder")
}

//判断文件是否存在的好方法
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (c *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	c.userClasspath = newEntry(cpOption)
}

//总是提示我写注释
func (c *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := c.bootClasspath.readClass(className); err == nil {
		log.Println("class:" + className + " loaded by bootClasspath")
		return data, entry, err
	}
	if data, entry, err := c.extClasspath.readClass(className); err == nil {
		log.Println("class:" + className + " loaded by extClasspath")
		return data, entry, err
	}
	return c.userClasspath.readClass(className)
}

func (c *Classpath) String() string {
	return c.userClasspath.String()
}
