/*
    Matrikelnummern: 5836402, 2416160
*/

package main

import (
	"io/ioutil"
	"encoding/json"
	"path/filepath"
	"os"
)

const BLOG_DIR string = "blogs"
const PERM os.FileMode = 0600

func saveBlog(b *Blog) error {
	_, err := ioutil.ReadDir("blogs")
	if err != nil {
		os.MkdirAll(filepath.Join(BLOG_DIR), PERM)
	}
	filename := filepath.Join(BLOG_DIR, b.Id + ".json")

	blogJson, err := json.Marshal(b)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, blogJson, PERM)
}

func loadBlog(id string) (*Blog, error) {
	filename := filepath.Join(BLOG_DIR, id + ".json")
	blogJson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var blog Blog
	json.Unmarshal(blogJson, &blog)
	return &blog, nil
}