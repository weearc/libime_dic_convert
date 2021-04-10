package main

import (
	"./UTF16Convert"
	"./Utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	filename := "test_file.txt"
	final, err := UTF16Convert.Replacement(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	s, err1 := Utils.UtilsSplit(final)
	if err1 != nil {
		log.Fatal(err1)
		return
	}

	for index := range s {
		sTest := strings.Split(s[index], "\t")
		fmt.Println(sTest[0])
		fmt.Println(Utils.IsChinese(sTest[0]))
		if Utils.IsChinese(sTest[0]) {
			Utils.WriteFile("test.txt", Utils.PinyinReformat(sTest[0], sTest[1]))
		}

	}
	//fmt.Println(s[10000])
	//sTest := s[10000]
	//sOut := strings.Split(sTest, "\t")
	//fmt.Println(sOut[0]) //原文
	////fmt.Println(sOut[1]) //词频
	////fmt.Println(sOut[2]) //拼音
	//fmt.Println(Utils.PinyinReformat(sOut[0], sOut[1]))
	//fmt.Println(Utils.IsChinese(sOut[0]))
	//fmt.Println(Utils.IsChinese("test#钇"))
}
