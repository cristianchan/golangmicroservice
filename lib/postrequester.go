package lib

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//PostRequest perfom post request
func PostRequest() {
	url := "https://www.googleapis.com/geolocation/v1/geolocate?key=AIzaSyA0uSmWKCub1EUkwxP7xd3S7xVqRIvGq84"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
