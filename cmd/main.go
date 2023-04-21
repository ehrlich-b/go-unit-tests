package main

import (
	"fmt"
	"net/http"

	"github.com/ehrlich-b/go-unit-tests/internal/interfaces"
	"github.com/ehrlich-b/go-unit-tests/internal/service"
)

func main() {
	httpClient := &http.Client{}
	fileSystem := interfaces.NewLocalFS()
	downloader := service.NewDownloader(fileSystem, httpClient)

	err := downloader.Download("http://example.com", "example.html")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Println("Downloaded and saved example.html")
	}
}
