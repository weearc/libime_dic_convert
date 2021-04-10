package main

import (
	"github.com/mozillazg/go-pinyin"
)


func PinyinConvert(source string) [][]string {
	hans := source
	Instance := pinyin.NewArgs()
	Instance.Style = pinyin.Normal
	return pinyin.Pinyin(hans, Instance)
}

//func main()  {
//	//hans := "测试一下，球墨铸铁"
//	hans := "钇"
//	s := PinyinConvert(hans)
//	fmt.Println(s)
//}
