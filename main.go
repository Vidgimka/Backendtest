package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

/**/

type GeoData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    []struct {
		Login            string  `json:"login"`
		Session_id       int     `json:"session_id"`
		Subnet           string  `json:"subnet"`
		Mountpoint       string  `json:"mountpoint"`
		Station          string  `json:"station"`
		Ntrip_agent      string  `json:"ntrip_agent"`
		Connect_time     int     `json:"connect_time"`
		Time_span        int     `json:"time_span"`
		Recieved_data    float64 `json:"recieved_data"`
		Sent_data        float64 `json:"sent_data"`
		Status_code      int     `json:"status_code"`
		Latency          int     `json:"latency"`
		Sv_num           int     `json:"sv_num"`
		Lat              float64 `json:"lat"`
		Lon              float64 `json:"lon"`
		Height           int     `json:"height"`
		Station_distance float64 `json:"station_distance"`
	} `json:"data"`
}

func ReadFileData() GeoData {
	URL := "https://svtp.prin.ru:8044/api/events/online-user"

	resp, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Response status:", resp.Status)
	/*
		io.Copy(os.Stdout, resp.Body) //тест корректного чтения  данных с API
		data := make([]byte, 1014)     // создаем байтовую переменную, чтобы записать респонс от API
			n, err := resp.Body.Read(data) // записываем тело ответа в переменную
			fmt.Println(string(data[:n]))  //вывод в консоль*/
	d, _ := io.ReadAll(resp.Body)

	var UsersOnline GeoData
	if err := json.Unmarshal(d, &UsersOnline); err != nil {
		//panic(err)
		log.Fatal(err.Error())
	}
	//fmt.Println(UsersOnline)
	return UsersOnline

}

func main() {
	var UsOn GeoData
	UsOn = ReadFileData()
	fmt.Println(UsOn)
}
