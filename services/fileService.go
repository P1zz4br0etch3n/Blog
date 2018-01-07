/*
    Matrikelnummern: 5836402, 2416160
*/

package services

import (
	"io/ioutil"
	"path/filepath"
	"os"
	"encoding/json"
)

const (
	BlogDir     = "blogs"
	UserDir     = "users"
	SettingsDir = "conf"
	dataDir     = "data"
	perm        = 0600
)

func WriteJsonFile(dir, name string, o interface{}) error {
	jsonData, e := json.Marshal(o)
	if e != nil {
		return e
	}
	return saveDataFile(dir, name + ".json", jsonData)
}

func ReadJsonFile(dir, name string, v interface{}) error {
	jsonData, e := readDataFile(dir, name + ".json")
	if e != nil {
		return e
	}
	e = json.Unmarshal(jsonData, v)
	if e != nil {
		return e
	}
	return nil
}

func saveDataFile(dir, name string, data []byte) error {
	filedir := filepath.Join(dataDir, dir)
	e := mkdirIfNotExist(filedir)
	if e != nil {
		return e
	}
	filename := filepath.Join(filedir, name)
	return ioutil.WriteFile(filename, data, perm)
}

func readDataFile(dir, name string) ([]byte, error) {
	path := filepath.Join(dataDir, dir, name)
	file, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}
	return file, nil
}

func mkdirIfNotExist(dir string) error {
	_, err := ioutil.ReadDir(dir)
	if err != nil {
		return os.MkdirAll(dir, perm)
	}
	return nil
}