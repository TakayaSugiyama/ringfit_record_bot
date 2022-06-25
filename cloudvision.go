package ringfit_record_bot

import (
	"context"
	"fmt"
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"

	vision "cloud.google.com/go/vision/apiv1"
)

type ExerciseResult struct {
	km   float64
	kcal float64
	min  float64
}

func getExerciseResult(photoUrl string) ExerciseResult {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	image := vision.NewImageFromURI(photoUrl)

	annotations, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatal(err)
	}

	var exeResult ExerciseResult
	kcalRule := regexp.MustCompile(`kcal`)
	kmRule := regexp.MustCompile(`km`)
	timeRule := regexp.MustCompile(`\d`)

	if len(annotations) == 0 {
		fmt.Println("not found")
	} else {
		for i, annotation := range annotations {
			if i == 0 {
				continue
			}
			description := annotation.Description
			if kcalRule.MatchString(description) {
				kcalString := strings.Replace(description, "kcal", "", 1)
				parseKcal, _ := strconv.ParseFloat(kcalString, 64)
				exeResult.kcal = parseKcal
			} else if kmRule.MatchString(description) {
				kmString := strings.Replace(description, "km", "", 1)
				parseKm, _ := strconv.ParseFloat(kmString, 64)
				exeResult.km = parseKm
			} else if timeRule.MatchString(description) {
				timeArray := strings.Split(description, "")
				min, _ := strconv.Atoi(timeArray[0] + timeArray[1])
				joined := strings.Join(timeArray[2:len(timeArray)-1], "")
				sec, _ := strconv.Atoi(joined)

				time := float64(min*60+sec) / 60.0
				exeResult.min = math.Round(time*100) / 100
			}
		}
	}
	return exeResult
}
