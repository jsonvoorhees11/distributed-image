package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const splittedFileSize = 10000

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

func writeFile(data []byte, filePath string) error {
	f, err := os.Create(filePath)
	check(err)
	defer func() {
		cerr := f.Close()
		if cerr == nil {
			err = cerr
		}
	}()
	n, err := f.Write(data)
	fmt.Printf("Bytes written: %v\n", n)
	return err
}

func getHashCode(input string, length ...int) string {
	h := sha1.New()
	io.WriteString(h, input)
	if len(length) > 0 {
		return hex.EncodeToString(h.Sum(nil)[0:length[0]])
	} else {
		return hex.EncodeToString(h.Sum(nil))
	}
}

func splitFile(reader *bufio.Reader, buffer []byte, imgPath string) {
	splittedDir := getSplittedDirectory(imgPath)
	hashedFileName := getHashCode(imgPath, 5)
	fmt.Printf("Splitted files directory: %v", splittedDir+hashedFileName)
	for i := 0; ; i++ {
		_, err := reader.Read(buffer)
		if err != nil {
			fmt.Println("End of file")
			break
		}
		//Format of new file /path-to-new-file/1.439248dff
		newFileDir := splittedDir + strconv.Itoa(i) + "." + hashedFileName
		writeFile(buffer, newFileDir)
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
	b := make([]byte, splittedFileSize)

	splitFile(r, b, imagePath)
}
