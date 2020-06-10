package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	city2 "github.com/alessandroarosio/weather-workshop/src/model/city"
	services2 "github.com/alessandroarosio/weather-workshop/src/services"
	"io"
	"os"
	"strconv"
	"sync"
)

func SaveCsvToDb() {
	file, err := os.Open("cities_20000.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// skip 1st line
	row1, _ := bufio.NewReader(file).ReadSlice('\n')
	file.Seek(int64(len(row1)), io.SeekStart)

	reader := csv.NewReader(file)
	record, _ := reader.ReadAll()

	buffer := make(chan bool, 10)
	var wg sync.WaitGroup

	for _, value := range record {
		buffer <- true
		wg.Add(1)
		go saveConcurrent(buffer, value, &wg)
	}
	wg.Wait()
	close(buffer)

	fmt.Println("cities saved!")
}

func saveConcurrent(buffer chan bool, record []string, wg *sync.WaitGroup) {
	id, _ := strconv.ParseInt(record[0], 10, 64)

	city := city2.City{
		Id:          id,
		CityName:    record[1],
		StateCode:   record[2],
		CountryCode: record[3],
		Country:     record[4],
		Lat:         record[5],
		Lon:         record[6],
	}
	services2.CityService.SaveCity(city)
	<- buffer
	wg.Done()
}
