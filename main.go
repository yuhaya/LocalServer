package main

import (
	_ "LocalServer/routers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"os"
	"os/exec"
	"path/filepath"
)

/**
 * 模板判断文件是否存在
 */
func isfile(paths ...string) bool {
	file, _ := exec.LookPath(os.Args[0])
	path_root, _ := filepath.Abs(file)
	var x string
	for _, n := range paths {
		x += n
	}
	path_total := path_root + "/.." + x
	fmt.Println(path_total)
	_, err := os.Stat(path_total)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	beego.AddFuncMap("isfile", isfile)
	beego.AddFuncMap("i18n", i18n.Tr)
	orm.Debug = true
	beego.Run()
}
