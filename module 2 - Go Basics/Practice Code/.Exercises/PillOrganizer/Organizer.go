package main

import (
	"fmt"
	"sort"
	"time"
)

type pill struct {
	name string
}

type user struct {
	name      string
	pills     []medicalRecipe
	schedules []schedule
}

type schedule struct {
	pill     string
	quantity float64
	timep    time.Time
}

type medicalRecipe struct {
	pill      pill
	frecuency int     //hours
	quantity  float64 //dose of pills
	duration  int     //days
}

//method that calculates the schedule for the pills with the current time
func (u *user) calculateSchedule() {
	for _, v := range u.pills {
		tempVar := time.Now()
		tempDay := time.Now() //.Format(ANSIC)
		for x := v.duration; x >= 0; x-- {
			//while infinite loop :/
			for tempVar.Before(tempDay) {
				//we need to add the frecuency to the current time to start adding dates to the schedule
				tempSchedule := schedule{pill: v.pill.name, quantity: v.quantity, timep: tempVar}
				//add new date
				u.schedules = append(u.schedules, tempSchedule)
				//increase with the frecuency of the medical recipe
				tempVar = tempVar.Add(time.Hour * time.Duration(v.frecuency))
			}
			tempDay = tempDay.AddDate(0, 0, 1) //add 1 day before to the temporal day variable
		}
	}

	sort.Slice(u.schedules, func(i int, j int) bool {
		return u.schedules[i].timep.Before(u.schedules[j].timep)
	})

}
func (u user) printSchedule() {
	for _, v := range u.schedules {
		fmt.Println("pill:", v.pill, "date:", v.timep.Format(time.ANSIC))
	}
}

func main() {
	dolex := pill{name: "dolex"}
	acetaminofem := pill{name: "acetaminofem"}
	//1 dolex every 8 hours per 4 days
	dolexRecipe := medicalRecipe{pill: dolex, quantity: 1, frecuency: 8, duration: 4}
	acetaminofemRecipe := medicalRecipe{pill: acetaminofem, quantity: 2, frecuency: 4, duration: 6}

	//user
	patient1 := user{name: "patient1"}
	patient1.pills = append(patient1.pills, dolexRecipe, acetaminofemRecipe)
	patient1.calculateSchedule()
	patient1.printSchedule()
}
