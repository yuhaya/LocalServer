package main

import (
	_ "LocalServer/routers"
	//	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
	"os"
	"path/filepath"
	"time"
)

/**
 * 模板判断文件是否存在
 */
func isfile(paths ...string) bool {

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	var x string
	for _, n := range paths {
		x += n
	}
	path_total := dir + x
	//	fmt.Println(path_total)
	_, err := os.Stat(path_total)
	if err != nil && os.IsNotExist(err) {
		//		fmt.Printf("\n=====%s====not found\n", path_total)
		return false
	}
	//	fmt.Printf("\n=====%s====found\n", path_total)
	return true
}

/**
 * 格式化时间
 */
func date_format(date time.Time) string {
	str := date.Format("2006-01-02")
	return str
}

/**
 * 判断map中得key是否存在
 */
func map_exist(map_val map[string][]string, key string) bool {
	_, ok := map_val[key]
	return ok
}

func main() {
	beego.AddFuncMap("isfile", isfile)
	beego.AddFuncMap("date_format", date_format)
	beego.AddFuncMap("map_exist", map_exist)

	beego.AddFuncMap("i18n", i18n.Tr)
	orm.Debug = true
	beego.Run()
}
