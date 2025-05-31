package getitems

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Item struct {
	ID    int    `json:"id"`
	Sid   int    `json:"sid"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}


type MarketPriceInfo struct {
	Name    string         `json:"name"`
	ID      int            `json:"id"`
	Sid     int            `json:"sid"`
	History map[string]int `json:"history"` // История цен в виде карты (timestamp -> цена)
}


type Methods map[string]string

type APIFunc func()

const baseUrl = "https://api.arsha.io"


func GetMarketPriceInfo(id int, sid int) (map[string]int, error) {
	url := baseUrl + fmt.Sprintf("/v2/ru/GetMarketPriceInfo?id=%d&sid=%d&lang=ru", id, sid) // ?id=12237&sid=0
	method := "GET"
	/*payload := strings.NewReader(fmt.Sprintf(`{
		"id": %d,
		"sid": %d
	}`, id, sid))*/

	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "BlackDesert")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		e := fmt.Errorf("ошибка при чтении тела ответа: %v", err)
		return nil, e
	}

	var MarketPriceInfo MarketPriceInfo
	err = json.Unmarshal(body, &MarketPriceInfo)
	if err != nil {
		fmt.Println("Error parse data")
		return nil, err
	}

	return MarketPriceInfo.History, nil

}


func GetWorldMarketList(mainCategory, subCategory int) []Item {
	url := baseUrl + fmt.Sprintf("/v2/ru/GetWorldMarketList?mainCategory=%d&subCategory=%d&lang=ru", mainCategory, subCategory)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Ошибка при создании запроса:", err)
		return []Item{}
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("User-Agent", "BlackDesert")
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Ошибка при выполнении запроса:", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {

	}

	var items []Item

	err = json.Unmarshal(body, &items)
	if err != nil {
		fmt.Println("parse JSON data error :", err)
		return nil
	}

	fmt.Println(res)
	fmt.Println(string(body))

	return items
}
