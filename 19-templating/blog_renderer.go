package blog_renderer

import (
	"io"
	"html/template"
	"embed"
	"strings"
)

type Post struct {
	Title string
	Description string
	Body string
	Tags []string
}

type PostViewModel struct {
	Title, SanitisedTitle, Description, Body string
	Tags []string
}

func (p Post) SanitisedTitle() string {
	return strings.ToLower(strings.Replace(p.Title, " ", "-", -1))
}


var (
	//go:embed "templates/*"
	postTemplates embed.FS 
)

type PostRenderer struct {
	templ *template.Template
}

func NewPostRenderer() (*PostRenderer, error) {
	templ, err:= template.ParseFS(postTemplates, "templates/*.gohtml")
	
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	return r.templ.ExecuteTemplate(w, "blog.gohtml", post)
}

func (r *PostRenderer) RenderIndex(w io.Writer, posts []Post) error {
	return r.templ.ExecuteTemplate(w, "index.gohtml", posts)
}