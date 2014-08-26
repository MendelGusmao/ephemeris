package fake

import (
	"github.com/martini-contrib/render"
	"html/template"
	"net/http"
)

type renderer struct{}

func (r renderer) JSON(status int, v interface{}) {

}

func (r renderer) HTML(status int, name string, v interface{}, htmlOpt ...render.HTMLOptions) {

}

func (r renderer) XML(status int, v interface{}) {

}

func (r renderer) Data(status int, v []byte) {

}

func (r renderer) Error(status int) {

}

func (r renderer) Status(status int) {

}

func (r renderer) Redirect(location string, status ...int) {

}

func (r renderer) Template() *template.Template {
	return &template.Template{}
}

func (r renderer) Header() http.Header {
	return make(http.Header)
}
