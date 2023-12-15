package blog_renderer

import (
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {

	var (
		post = Post {
			Title: "hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts a single post int HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, post)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders a index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}

		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		err := postRenderer.RenderIndex(&buf, posts)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})

	t.Run("it renders an index of posts", func(t *testing.T) {
		buf := bytes.Buffer{}
		posts := []Post{{Title: "Hello World"}, {Title: "Hello World 2"}}

		if err := postRenderer.RenderIndex(&buf, posts); err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var (
		post = Post {
			Title: "hello world",
			Body: "This is a post",
			Description: "This is a description",
			Tags: []string{"go", "tdd"},
		}
	)

	postRenderer, err := NewPostRenderer()

	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	
	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, post)
	}
}


