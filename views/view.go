package views

import (
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"path/filepath"

	"github.com/petrostrak/darkroom/context"

	"github.com/gorilla/csrf"
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
	t, err := template.New("").Funcs(template.FuncMap{
		"csrfField": func() (template.HTML, error) {
			return "", errors.New("csrf field is not implemented")
		},
	}).ParseFiles(files...)
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
	if alert := getAlert(r); alert != nil {
		vd.Alert = alert
		clearAlert(w)
	}
	vd.User = context.User(r.Context())
	var buf bytes.Buffer
	csrfField := csrf.TemplateField(r)
	tpl := v.Template.Funcs(template.FuncMap{
		"csrfField": func() template.HTML {
			return csrfField
		},
	})
	if err := tpl.ExecuteTemplate(&buf, v.Layout, vd); err != nil {
		http.Error(w, "Something went wrong. If the proplem persists, please email us at support@github.com/petrostrak/darkroom.com", http.StatusInternalServerError)
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
