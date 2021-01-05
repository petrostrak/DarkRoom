package views

import "html/template"

// NewView func to instantiate view objects
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")

	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
	}
}

// the view structure
type View struct {
	Template *template.Template
}
