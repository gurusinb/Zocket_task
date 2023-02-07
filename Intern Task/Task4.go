package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func downloadFile(url string, fileName string, c chan string) {
	response, err := http.Get(url)
	if err != nil {
		c <- fmt.Sprintf("Error while downloading %s: %s", fileName, err)
		return
	}
	defer response.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		c <- fmt.Sprintf("Error while creating file %s: %s", fileName, err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		c <- fmt.Sprintf("Error while writing file %s: %s", fileName, err)
		return
	}

	c <- fmt.Sprintf("%s has been downloaded successfully", fileName)
}

func main() {
	files := []struct {
		url      string
		fileName string
	}{
		{"https://example.com/file1.pdf", "file1.pdf"},
		{"https://example.com/file2.pdf", "file2.pdf"},
		{"https://example.com/file3.pdf", "file3.pdf"},
	}

	c := make(chan string)

	for _, file := range files {
		go downloadFile(file.url, file.fileName, c)
	}

	for range files {
		fmt.Println(<-c)
	}
}
