package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getSplittedDirectory(imgPath string) string {
	folders := strings.Split(imgPath, "/")
	folders[len(folders)-1] = ""
	folders[len(folders)-2] = "splitted"
	return strings.Join(folders, "/")
}

func splitFile(reader *bufio.Reader, buffer []byte, imgPath string) {

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("End of file")
			break
		}

		fmt.Println(buffer[0:n])
	}
}

func main() {
	fptr := flag.String("fpath", "test.jpg", "file path to read from")
	flag.Parse()
	imagePath := *fptr
	f, err := os.Open(imagePath)
	check(err)

	defer func() {
		check(f.Close())
	}()

	r := bufio.NewReader(f)
	b := make([]byte, 2000)

	splitFile(r, b, imagePath)
}
