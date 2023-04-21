package service

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ehrlich-b/go-unit-tests/internal/interfaces"
)

type Downloader struct {
	fs         interfaces.FS
	httpClient *http.Client
}

func NewDownloader(fs interfaces.FS, httpClient *http.Client) *Downloader {
	return &Downloader{
		fs:         fs,
		httpClient: httpClient,
	}
}

func (d *Downloader) Download(httpUrl, outFile string) error {
	resp, err := d.httpClient.Get(httpUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP error: %v", resp.Status)
	}

	file, err := d.fs.OpenFile(outFile, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}
