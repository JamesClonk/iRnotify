package monitoring

import (
	"log"
	"time"

	"github.com/JamesClonk/iRnotify/lib/env"
	"github.com/JamesClonk/iRnotify/lib/racers"
)

type Monitor struct{}

func New() *Monitor {
	return &Monitor{}
}

func (m *Monitor) Start() {
	monitorRacers()
}

func monitorRacers() {
	go func() {
		for {
			time.Sleep(1 * time.Minute)

			data, err := racers.GetRacers(env.Get("IRACING_NAME", "Fabio+Berchtold"))
			if err != nil {
				log.Println(err)
				continue
			}

			for _, racer := range data.Racers {
				if racer.UserRole == 0 && racer.SubSessionStatus == "subses_running" {
					//log.Printf("Session running! \n%#v\n", racer)
					if err := notify(racer); err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
}
