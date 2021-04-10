package Utils

import (
	_ "fmt"
	"github.com/mozillazg/go-pinyin"
)

func PinyinConvert(source string) [][]string {
	hans := source
	Instance := pinyin.NewArgs()
	Instance.Style = pinyin.Normal
	return pinyin.Pinyin(hans, Instance)
}

func PinyinRejoint(source string) string {
	pinyinOut := PinyinConvert(source)
	var stringJoint string
	for index := range pinyinOut {
		stringJoint += pinyinOut[index][0] + "'"
	}
	return stringJoint[:len(stringJoint)-1]
}

func PinyinReformat(word string, freq string) string {
	pinyin := PinyinRejoint(word)
	return word + " " + pinyin + " " + freq + "\n"

}

//func main()  {
//	//hans := "测试一下，球墨铸铁"
//	hans := "钇"
//	s := PinyinConvert(hans)
//	fmt.Println(s)
//}
