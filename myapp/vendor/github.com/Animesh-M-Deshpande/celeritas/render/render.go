package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/CloudyKit/jet/v6"
)

type Render struct {
	Renderer   string
	RootPath   string
	Secure     bool
	Port       string
	ServerName string
	JetViews   *jet.Set
}

type TemplateData struct {
	IsAuthenticated bool
	IntMap          map[string]int
	StringMap       map[string]string
	FloatMap        map[string]float32
	Data            map[string]interface{}
	CSRFToken       string
	Port            string
	ServerName      string
	Secure          bool
}

func (c *Render) Page(w http.ResponseWriter, r *http.Request, view string, variables, data interface{}) error {

	switch strings.ToLower(c.Renderer) {
	case "go":
		c.GoPage(w, r, view, data)
	case "jet":
		c.JetPage(w, r, view, variables, data)
	}

	return nil
}

//JetPage renders the template using the JetPage templating engine
func (c *Render) JetPage(w http.ResponseWriter, r *http.Request, templateName string, variables, data interface{}) error {

	var vars jet.VarMap

	if variables == nil { // initializes if its nil
		vars = make(jet.VarMap)
	} else {
		vars = variables.(jet.VarMap) // cast it to correct format
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}
	fmt.Println("Jet page renderer, page is:", fmt.Sprintf("%s/views/%s.jet", c.RootPath, templateName))
	t, err := c.JetViews.GetTemplate(fmt.Sprintf("%s/views/%s.jet", c.RootPath, templateName))

	if err != nil {
		log.Println(err)
		return err
	}

	if err = t.Execute(w, vars, td); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

//renders standard  Go template
func (c *Render) GoPage(w http.ResponseWriter, r *http.Request, view string, data interface{}) error {

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/views/%s.page.tmpl", c.RootPath, view))

	if err != nil {
		return err
	}

	td := &TemplateData{}

	if data != nil {
		td = data.(*TemplateData)
	}

	err = tmpl.Execute(w, &td)

	if err != nil {
		return err
	}

	return nil
}
