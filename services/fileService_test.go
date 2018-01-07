/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"de/vorlesung/projekt/2416160-5836402/models"
	"time"
	"github.com/stretchr/testify/assert"
	"testing"
	"path/filepath"
	"os"
	"io/ioutil"
)

const TestDir string = "test"

func TestWriteJsonFile(t *testing.T) {
	blog := models.BlogPost{
		PostID:   "999999999",
		Content:  "text",
		Date:     time.Now(),
		Author:   "p1zz4br0etch3n",
		Comments: nil,
	}

	WriteJsonFile(TestDir, blog.PostID, blog)
	path := filepath.Join(dataDir, TestDir, blog.PostID+".json")

	assert.FileExists(t, path, "File not found.")

	os.RemoveAll(dataDir)
}

func TestReadJsonFile(t *testing.T) {
	var blog models.BlogPost
	path := filepath.Join(dataDir, TestDir, "1000.json")

	e := os.MkdirAll(filepath.Join(dataDir, TestDir), perm)
	assert.NoError(t, e)

	e = ioutil.WriteFile(path, []byte("{\"PostID\":\"1000\",\"Content\":\"Beispieltext\",\"Date\":\"2018-01-05T22:14:25.8565125+01:00\",\"Author\":\"p1zz4br0etch3n\",\"Comments\":null}"), perm)
	assert.NoError(t, e)

	e = ReadJsonFile(TestDir, "1000", &blog)
	assert.NoError(t, e)

	assert.Equal(t, "1000", blog.PostID)

	os.RemoveAll(dataDir)
}

func TestDeleteDataFile(t *testing.T) {
	filename := "test.txt"
	path := filepath.Join(dataDir, TestDir)

	e := os.MkdirAll(path, perm)
	assert.NoError(t, e)

	e = ioutil.WriteFile(filepath.Join(path, filename), []byte("Test321"), perm)
	assert.NoError(t, e)

	e = deleteDataFile(TestDir, filename)
	assert.NoError(t, e)

	_, err := os.Stat(filepath.Join(path, filename))
	assert.Error(t, err)

	os.RemoveAll(dataDir)
}
