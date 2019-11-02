package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sort"
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

func getFiles(dir string, fileName string) ([]string, error) {
	var files []string
	f, err := os.Open(dir)
	if err != nil {
		return files, err
	}

	defer func() {
		check(f.Close())
	}()

	fileInfo, err := f.Readdir(-1)
	if err != nil {
		return files, err
	}
	for _, file := range fileInfo {
		if strings.Contains(file.Name(), fileName) {
			files = append(files, file.Name())
		}
	}
	sort.Strings(files)
	return files, nil
}

func mergeFiles(srcDir string, dir string, fileName string) error {
	files, err := getFiles(srcDir, fileName)
	check(err)
	f, err := os.OpenFile(dir+"/"+fileName+".jpg", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)
	defer func() {
		cerr := f.Close()
		if cerr == nil {
			err = cerr
		}
	}()
	fmt.Println(files)
	for _, file := range files {
		in, err := os.Open(srcDir + "/" + file)
		stat, err := in.Stat()
		check(err)
		fmt.Println("size of " + file + ": " + strconv.FormatInt(stat.Size(), 10))
		check(err)
		_, cerr := io.Copy(f, in)
		check(cerr)
		in.Close()
	}

	return err
}

func main1() {
	// fptr := flag.String("fpath", "test.jpg", "file path to read from")
	// flag.Parse()
	// imagePath := *fptr
	// f, err := os.Open(imagePath)
	// check(err)

	// defer func() {
	// 	check(f.Close())
	// }()

	// r := bufio.NewReader(f)
	// b := make([]byte, splittedFileSize)

	// splitFile(r, b, imagePath)

	// str, err := getFiles("/home/victor/gocode/src/github.com/jsonvoorhees11/distributed-image/splitted", "a64a2dd2c7")
	// check(err)
	// for i, val := range str {
	// 	if i > 0 && val != str[i] {
	// 		panic("Wrong")
	// 	}
	// }

	err := mergeFiles(
		"/home/victor/gocode/src/github.com/jsonvoorhees11/distributed-image/splitted",
		"/home/victor/gocode/src/github.com/jsonvoorhees11/distributed-image/merged",
		"a64a2dd2c7")

	check(err)
}
