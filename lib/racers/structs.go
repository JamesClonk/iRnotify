package racers

import (
	"strconv"
	"time"
)

type RacingData struct {
	Timestamp time.Time `json:"timestamp"`
	Racers    []Racer   `json:"fsRacers"`
}

type Racer struct {
	Broadcast      interface{} `json:"broadcast"`
	DriverChanges  bool        `json:"driverChanges"`
	LastLogin      unixTime    `json:"lastLogin"`
	MaxUsers       int         `json:"maxUsers"`
	HasGrid        int         `json:"hasGrid"`
	TrackId        int         `json:"trackId"`
	SessionStatus  string      `json:"sessionStatus"`
	SessionTypeID  int         `json:"sessionTypeId"`
	PrivateSession interface{} `json:"privateSession"`
	SeriesID       int         `json:"seriesId"`
	RegOpen        bool        `json:"regOpen"`
	CatID          int         `json:"catId"`
	EventTypeID    int         `json:"eventTypeId"`
	SpotterAccess  int         `json:"spotterAccess"`
	LastSeen       unixTime    `json:"lastSeen"`
	SeasonID       int         `json:"seasonId"`
	Helmet         struct {
		Color1 string `json:"c1"`
		Color2 string `json:"c2"`
		Color3 string `json:"c3"`
	} `json:"helmet"`
	PrivateSessionID int      `json:"privateSessionId"`
	inGrid           int      `json:"inGrid"`
	CustID           int      `json:"custid"`
	Name             string   `json:"name"`
	StartTime        unixTime `json:"startTime"`
	UserRole         int      `json:"userRole"`
	SubSessionStatus string   `json:"subSessionStatus"`
}

type CareerStats struct {
	Wins                    int     `json:"wins"`
	TotalClubPoints         int     `json:"totalclubpoints"`
	WinPercentage           float64 `json:"winPerc"`
	Poles                   int     `json:"poles"`
	AverageStart            float64 `json:"avgStart"`
	AverageFinish           float64 `json:"avgFinish"`
	Top5Percentage          float64 `json:"top5Perc"`
	TotalLaps               int     `json:"totalLaps"`
	AverageIncidentsPerRace float64 `json:"avgIncPerRace"`
	AveragePointsPerRace    float64 `json:"avgPtsPerRace"`
	LapsLed                 int     `json:"lapsLed"`
	Top5                    int     `json:"top5"`
	LapsLedPercentage       float64 `json:"lapsLedPerc"`
	Category                string  `json:"category"`
	Starts                  int     `json:"starts"`
}

type unixTime struct {
	time.Time
}

func (u *unixTime) UnmarshalJSON(data []byte) error {
	unix, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}

	*u = unixTime{time.Unix(unix/1000, 0)}
	return nil
}
