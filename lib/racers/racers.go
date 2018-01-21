package racers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var racingData RacingData

func getData(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Jar: cookieJar,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("GET: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status code: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}
	return data, nil
}

func parseRacingData(data []byte) error {
	racingData = RacingData{} // clear all previously stored values
	if err := json.Unmarshal(data, &racingData); err != nil {
		return err
	}

	for _, racer := range racingData.SearchRacers {
		racingData.Racers = append(racingData.Racers, racer)
	}
	for i := range racingData.Racers {
		racingData.Racers[i].Name = strings.Replace(racingData.Racers[i].Name, "+", " ", -1)
		racingData.Racers[i].TrackName = Track(racingData.Racers[i].TrackID)
		racingData.Racers[i].CarName = Car(racingData.Racers[i].CarID)
		racingData.Racers[i].SessionTypeName = SessionType(racingData.Racers[i].SessionTypeID, racingData.Racers[i].UserRole)
	}
	racingData.Timestamp = time.Now()

	return nil
}

func GetRacers(name string) (RacingData, error) {
	if time.Now().After(racingData.Timestamp.Add(3 * time.Minute)) {
		if err := updateRacingData(name); err != nil {
			return racingData, err
		}
	}
	return racingData, nil
}

func updateRacingData(name string) error {
	// login if older than 15 minutes
	if time.Now().After(racingData.Timestamp.Add(15 * time.Minute)) {
		if err := Login(); err != nil {
			return err
		}
	}

	data, err := getData(fmt.Sprintf("http://members.iracing.com/membersite/member/GetDriverStatus?friends=1&studied=1&searchTerms=%s", name))
	if err != nil {
		return err
	}

	if err := parseRacingData(data); err != nil {
		return err
	}
	return nil
}

func GetRacer(id int) Racer {
	for _, racer := range racingData.Racers {
		if racer.CustID == id {
			return racer
		}
	}
	return Racer{}
}

func GetStats(id int) ([]CareerStats, error) {
	data, err := getData(fmt.Sprintf("http://members.iracing.com/memberstats/member/GetCareerStats?custid=%d", id))
	if err != nil {
		return nil, err
	}

	stats := make([]CareerStats, 0)
	if err := json.Unmarshal(data, &stats); err != nil {
		return nil, err
	}

	for i := range stats {
		stats[i].Category = strings.Replace(stats[i].Category, "+", " ", -1)
	}

	return stats, nil
}
