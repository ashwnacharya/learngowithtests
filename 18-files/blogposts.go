package blogposts

import (
	"testing/fstest"
	"io/fs"
	"bufio"
	"strings"
	"bytes"
	"fmt"
)

type Post struct {
	Title string
	Description string
	Tags []string
	Body string
}

func NewPostsFromFS(filesystem fstest.MapFS) ([]Post, error) {
	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		return nil, err
	}

	var posts []Post 

	for _, f := range dir {
		
		post, err := getPost(filesystem, f)

		if err != nil {
			return nil, err
		}
		
		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(filesystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := filesystem.Open(f.Name())

	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()
	return newPost(postFile)

}

const (
	titleSeparator = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparatator = "Tags: "
)

func newPost(postFile fs.File) (Post, error) {

	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	titleLine := readMetaLine(titleSeparator)
	descriptionLine := readMetaLine(descriptionSeparator)
	tags:= strings.Split(readMetaLine(tagsSeparatator), ", ")
	body := readBody(scanner)
	

	post := Post{Title: titleLine, Description: descriptionLine, Tags: tags, Body: body}
	return post, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()

	buf := bytes.Buffer{}

	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	return strings.TrimSuffix(buf.String(), "\n")
}