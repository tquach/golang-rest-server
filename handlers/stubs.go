package handlers

import (
	"html/template"
	"net/http"

	"github.com/martini-contrib/render"
)

type stubrender struct {
	status   int
	response interface{}
}

func (r *stubrender) JSON(status int, v interface{}) {
	r.status = status
	r.response = v
}
func (r *stubrender) HTML(status int, name string, v interface{}, htmlOpt ...render.HTMLOptions) {}
func (r *stubrender) Error(status int)                                                           {}
func (r *stubrender) Redirect(location string, status ...int)                                    {}
func (r *stubrender) Template() *template.Template {
	return nil
}
func (r *stubrender) Data(int, []byte) {
}
func (r *stubrender) Header() (h http.Header) {
	return
}
func (r *stubrender) Status(int) {
}
func (r *stubrender) XML(int, interface{}) {
}
