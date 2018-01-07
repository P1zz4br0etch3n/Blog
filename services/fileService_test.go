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
	blog := models.Blog{
		Id:       "999999999",
		Title:    "titel",
		Text:     "text",
		Time:     time.Now(),
		Author:   nil,
		Comments: nil,
	}

	WriteJsonFile(TestDir, blog.Id, blog)
	path := filepath.Join(dataDir, TestDir, blog.Id + ".json")

	assert.FileExists(t, path, "File not found.")

	os.RemoveAll(dataDir)
}

func TestReadJsonFile(t *testing.T) {
	var blog models.Blog
	path := filepath.Join(dataDir, TestDir, "1000.json")

	e := os.MkdirAll(filepath.Join(dataDir, TestDir), perm)
	assert.NoError(t, e)

	e = ioutil.WriteFile(path, []byte("{\"id\":\"1000\",\"Title\":\"Titel\",\"Text\":\"Beispieltext\",\"Time\":\"2018-01-05T22:14:25.8565125+01:00\",\"Author\":{\"Username\":\"p1zz4br0etch3n\",\"Password\":\"test\"},\"Comments\":null}"), perm)
	assert.NoError(t, e)

	e = ReadJsonFile(TestDir, "1000", &blog)
	assert.NoError(t, e)

	assert.Equal(t, "1000", blog.Id)

	os.RemoveAll(dataDir)
}