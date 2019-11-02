package fileManip

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"

	strHelper "github.com/jsonvoorhees11/distributed-image/string-helpers"
)

const splittedFileSize = 10000

func SplitFile(imgPath string) {
	f, err := os.Open(imgPath)
	check(err)
	defer func() {
		check(f.Close())
	}()

	reader := bufio.NewReader(f)
	buffer := make([]byte, splittedFileSize)

	splittedDir := getSplittedDirectory(imgPath)
	hashedFileName := strHelper.GetHashCode(imgPath, 5)
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

func MergeFiles(srcDir string, dir string, fileName string) error {
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
