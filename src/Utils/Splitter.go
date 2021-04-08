package Utils

import (
	"strings"
)

func UtilsSplit(InSring string) ([]string, error) {
	s := strings.Split(InSring, "\n")
	return s, nil
}

//func main() {
//	final := "test \n test"
//	s, err1 := UtilsSplit(final)
//	if err1 != nil {
//		log.Fatal()
//		return
//	}
//	fmt.Println(s[0])
//}
