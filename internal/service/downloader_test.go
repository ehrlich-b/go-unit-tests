package service

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ehrlich-b/go-unit-tests/internal/interfaces/mocks"
	"github.com/stretchr/testify/assert"
)

func TestDownloadAndSave(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("content"))
	}))
	defer testServer.Close()

	httpClient := &http.Client{}
	fileSystem := &mocks.FS{}

	file := &mocks.WriteCloser{}
	file.On("Write", []byte("content")).Return(len("content"), nil)
	file.On("Close").Return(nil)

	fileSystem.On("OpenFile", "example.html", os.O_CREATE|os.O_WRONLY, os.ModePerm).Return(file, nil)

	exampleDownloader := NewDownloader(fileSystem, httpClient)

	err := exampleDownloader.Download(testServer.URL, "example.html")
	assert.NoError(t, err)

	file.AssertExpectations(t)
	fileSystem.AssertExpectations(t)
}
