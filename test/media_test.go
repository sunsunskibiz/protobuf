package test_test

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMediaCSV(t *testing.T) {
	filePath := "pokemon.csv"
	url := "http://localhost:8080/v1/csv"

	file, err := os.Open(filePath)
	require.NoError(t, err)
	defer func() { assert.NoError(t, file.Close()) }()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	require.NoError(t, err)

	_, err = io.Copy(part, file)
	require.NoError(t, err)

	err = writer.Close()
	require.NoError(t, err)

	request, err := http.NewRequest(http.MethodPost, url, body)
	require.NoError(t, err)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}

	response, err := client.Do(request)
	require.NoError(t, err)
	defer response.Body.Close()

	rawResponse, err := io.ReadAll(response.Body)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Contains(t, string(rawResponse), "Mu,1000")
}
