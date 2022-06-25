package ringfit_record_bot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetIfttt(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Url string `json:"url"`
	}

	fmt.Println(r.Body)
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Fprint(w, "can not decode url")
		return
	}

	if data.Url == "" {
		fmt.Fprint(w, "url is empty value")
		return
	}
	fmt.Println("request is valid")

	var photoUrl string = fetchPhotoURL(data.Url)
	var excResult ExerciseResult = getExerciseResult(photoUrl)
	fmt.Println(excResult)
}
