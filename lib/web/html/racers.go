package html

import (
	"net/http"
	"strconv"

	"github.com/JamesClonk/iRnotify/lib/racers"
	"github.com/JamesClonk/iRnotify/lib/web"
	"github.com/gorilla/mux"
)

func Racers(rw http.ResponseWriter, req *http.Request) {
	page := &Page{
		Title:  "iRnotify - Racers",
		Active: "racers",
	}

	data, err := racers.GetFriends()
	if err != nil {
		Error(rw, err)
		return
	}

	page.Content = struct {
		Racers []racers.Racer
	}{
		data.Racers,
	}

	web.Render().HTML(rw, http.StatusOK, "racers", page)
}

func Racer(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if len(id) > 0 {
		page := &Page{
			Title:  "iRnotify - Racer",
			Active: "racers",
		}

		custId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			Error(rw, err)
			return
		}

		stats, err := racers.GetStats(int(custId))
		if err != nil {
			Error(rw, err)
			return
		}

		page.Content = struct {
			Racer racers.Racer
			Stats []racers.CareerStats
		}{
			racers.GetRacer(int(custId)),
			stats,
		}

		web.Render().HTML(rw, http.StatusOK, "racer", page)
		return
	}
	web.Render().JSON(rw, http.StatusNotFound, nil)
}
