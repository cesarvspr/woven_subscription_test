package woven_challenge

import (
	"strconv"
	"strings"
	"time"
)

type Subscription struct {
	Id                    int
	CustomerId            int
	MonthlyPriceInDollars int
}

type User struct {
	Id            int
	Name          string
	ActivatedOn   time.Time
	DeactivatedOn time.Time
	CustomerId    int
}

func BillFor(yearMonth string, activeSubscription *Subscription, users *[]User) float64 {

	if len(*users) == 0 {
		return 0
	}

	data := strings.Split(yearMonth, "-")
	var result float64
	var months = map[string]int{}

	for i := time.January; i <= time.December; i++ {
		months[i.String()] = int(i)
	}

	year, _ := strconv.Atoi(data[0])
	month, _ := strconv.Atoi(data[1])

	//get how many days a month had
	firstDay := FirstDayOfMonth(time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC))
	lastDay := LastDayOfMonth(time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC))
	days := lastDay.Sub(firstDay).Hours() / 24

	// gey daily rate with days/price
	dailyRate := float64(activeSubscription.MonthlyPriceInDollars) / days

	// get total price for month
	for _, user := range *users {
		if months[user.ActivatedOn.Month().String()] == month {
			last := LastDayOfMonth(user.ActivatedOn)
			daysUsed := last.Sub(user.ActivatedOn).Hours() / 24
			result += daysUsed * dailyRate
			continue
		}

		if user.DeactivatedOn.IsZero() {
			result += dailyRate * days
			continue
		}
	}
	return result
}

/*******************
* Helper functions *
*******************/

/*
Takes a time.Time object and returns a time.Time object
which is the first day of that month.

FirstDayOfMonth(time.Date(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb 7
=> time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC))               // Feb 1
*/
func FirstDayOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

/*
Takes a time.Time object and returns a time.Time object
which is the end of the last day of that month.

LastDayOfMonth(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))  // Feb  7
=> time.Time(2019, 2, 28, 23, 59, 59, 0, time.UTC)           // Feb 28
*/
func LastDayOfMonth(t time.Time) time.Time {
	return FirstDayOfMonth(t).AddDate(0, 1, 0).Add(-time.Second)
}

/*
Takes a time.Time object and returns a time.Time object
which is the next day.

NextDay(time.Time(2019, 2, 7, 0, 0, 0, 0, time.UTC))   // Feb 7
=> time.Time(2019, 2, 8, 0, 0, 0, 0, time.UTC)         // Feb 8

NextDay(time.Time(2019, 2, 28, 0, 0, 0, 0, time.UTC))  // Feb 28
=> time.Time(2019, 3, 1, 0, 0, 0, 0, time.UTC)         // Mar  1
*/
func NextDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}
