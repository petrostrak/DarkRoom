package views

import (
	"DarkRoom/context"
	"bytes"
	"html/template"
	"io"
	"net/http"
	"path/filepath"
)

var (
	LayoutDir   string = "views/layouts/"
	TemplateDir string = "views/"
	TemplateExt string = ".gohtml"
)

// NewView func to instantiate view objects
func NewView(layout string, files ...string) *View {
	// fmt.Println("Before:", files)
	addTemplatePath(files)
	// fmt.Println("after path:", files)
	addTemplateExt(files)
	// fmt.Println("after ext:", files)
	files = append(files, layoutFiles()...)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

// the view structure
type View struct {
	Template *template.Template
	Layout   string
}

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v.Render(w, r, nil)
}

// Render is used to render the view with the predefined layout
func (v *View) Render(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	var vd Data
	switch d := data.(type) {
	case Data:
		vd = d
		// do nothing
	default:
		vd = Data{
			Yield: data,
		}
	}
	vd.User = context.User(r.Context())
	var buf bytes.Buffer
	if err := v.Template.ExecuteTemplate(&buf, v.Layout, vd); err != nil {
		http.Error(w, "Something went wrong. If the proplem persists, please email us at support@darkroom.com", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

// returns all layouts
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files
}

// takes in a slice of strings representing file path
// for templates and it prepends the TemplateDir to each
// in the slice
// e.g. the input {"home"} would result in the output
// {"views/home"} if TemplateDir == "views/"
func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}
}

// takes in a slice of strings representing file path for
// templates and it appends the TemplateExt extension to each
// string in the slice
// e.g. the input {"home"} would result in the output
// {"home.gohtml"} if TemplateDir == ".gohtml"
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
