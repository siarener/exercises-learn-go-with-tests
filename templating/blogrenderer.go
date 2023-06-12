package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Body, Description string

	Tags []string
}

type PostRenderer struct {
	templ *template.Template
}

var (
	/* Package embed provides access to files embedded in the running Go program.

	Go source files that import "embed" can use the //go:embed directive to initialize a variable of type
	string, []byte, or FS with the contents of files read from the package directory or subdirectories at
	compile time.

	(source: https://pkg.go.dev/embed)
	*/
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{templ: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, p Post) error {
	if err := r.templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
