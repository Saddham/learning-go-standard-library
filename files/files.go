package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
)

func main() {
	fmt.Println(checkFileExists("sampletext.txt"))
	fmt.Println(checkFileExists("files/sampletext.txt"))

	stats, err := os.Stat("files/sampletext.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("Modification time:", stats.ModTime())
	fmt.Println("File mode:", stats.Mode())

	if stats.Mode().IsRegular() {
		fmt.Println("This is a regular file")
	}

	fmt.Println("File size:", stats.Size())
	fmt.Println("Is dir:", stats.IsDir())

	// *** Writing to Files ***

	data := []byte("This is some text data\n")
	os.WriteFile("files/sampletext1.txt", data, 0644)

	fname := "files/testfile.txt"
	f, err := os.Create(fname)
	handleError(err)

	defer f.Close()

	fmt.Println("The file name is", f.Name())

	f.WriteString("This is some text\n")

	data2 := []byte{'a', 'b', 'c', '\n'}
	f.Write(data2)

	// os.Truncate(f.Name(), 10) // Truncate to 10 bytes

	f.Sync() // Flush, forces data to be written out to disk

	appendFileData(fname, "This data was appended")

	// *** Read File Content ****

	content, err := os.ReadFile(fname) // Returns bytes
	handleError(err)

	fmt.Println("File content:")
	fmt.Println(string(content))

	const bufferSize = 20
	f2, _ := os.Open(fname)
	defer f2.Close()

	buffer := make([]byte, bufferSize)
	for {
		n, err := f2.Read(buffer)

		if err != nil {
			if err != io.EOF {
				handleError(err)
			}
			break
		}

		fmt.Println("Bytes read:", n)
		fmt.Println("Content:", string(buffer[:n]))
	}

	// Read file content at a specific offset
	l := getFileStats(fname).Size()
	buffer2 := make([]byte, 10)
	_, err2 := f2.ReadAt(buffer2, l-int64(len(buffer2))) // 10 bytes from the end of the file
	handleError(err2)
	fmt.Println("File length:", l)
	fmt.Println("Content at the offset:", string(buffer2))

	// *** Working with Directories
	os.Mkdir("files/newdir", os.ModePerm)
	os.MkdirAll("files/newdir2/newdir3", os.ModePerm)

	defer os.Remove("files/newdir")
	defer os.RemoveAll("files/newdir2")

	dir, _ := os.Getwd() // Curr work dir
	fmt.Println("Current dir is", dir)

	exedir, _ := os.Executable() // Cur dir of the process
	fmt.Println("Exe dir is", exedir)

	contents, _ := os.ReadDir("files")
	for _, fi := range contents {
		fmt.Println(fi.Name(), fi.IsDir())
	}

	// Working with Temporary Directories
	tempPath := os.TempDir()
	fmt.Println("Default temp dir is", tempPath)

	tempContent := []byte("This is some temp content for the file")
	tempFile, err := os.CreateTemp(tempPath, "tempfile_*.txt") // * add some random chars if file exits
	handleError(err)

	if _, err := tempFile.Write(tempContent); err != nil {
		panic(err)
	}

	data3, _ := os.ReadFile(tempFile.Name())
	fmt.Printf("%s\n", data3)

	fmt.Println("Temp file is named:", tempFile.Name())
	defer os.Remove(tempFile.Name())

	tempDir, err := os.MkdirTemp("", "tempdir*") // empty param uses default temp dir
	handleError(err)
	fmt.Println("The temp dir is named:", tempDir)
	defer os.RemoveAll(tempDir)
}

func checkFileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func getFileStats(filePath string) fs.FileInfo {
	stats, err := os.Stat(filePath)
	handleError(err)

	return stats
}

func appendFileData(fname string, data string) {
	f, err := os.OpenFile(fname, os.O_APPEND|os.O_WRONLY, 0644)
	handleError(err)

	defer f.Close()

	_, err2 := f.WriteString(data)
	handleError(err2)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
