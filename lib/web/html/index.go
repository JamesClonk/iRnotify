package html

import (
	"fmt"
	"net/http"

	"github.com/JamesClonk/iRnotify/lib/web"
)

func Index(rw http.ResponseWriter, req *http.Request) {
	page := &Page{
		Title:  "iRnotify",
		Active: "home",
	}
	web.Render().HTML(rw, http.StatusOK, "index", page)
}

func NotFound(rw http.ResponseWriter, req *http.Request) {
	page := &Page{
		Title: "iRnotify - Not Found",
	}
	web.Render().HTML(rw, http.StatusNotFound, "not_found", page)
}

func ErrorHandler(rw http.ResponseWriter, req *http.Request) {
	Error(rw, fmt.Errorf("Internal Server Error"))
}
func Error(rw http.ResponseWriter, err error) {
	page := &Page{
		Title:   "iRnotify - Error",
		Content: err,
	}
	web.Render().HTML(rw, http.StatusInternalServerError, "error", page)
}
