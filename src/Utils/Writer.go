package Utils

import (
	"bufio"
	"fmt"
	"os"
)

func WriteFile(filename string, content string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("fail to create file", err)
	}
	write := bufio.NewWriter(file)
	write.WriteString(content)
	write.Flush()
}
