package main

import (
	"time"
	. "woven_test/woven_challenge"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BillFor", func() {
	constantUsers := []User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}

	userSignedUp := []User{
		{
			Id:            1,
			Name:          "Employee #1",
			ActivatedOn:   time.Date(2018, 11, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            2,
			Name:          "Employee #2",
			ActivatedOn:   time.Date(2018, 12, 4, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
		{
			Id:            3,
			Name:          "Employee #3",
			ActivatedOn:   time.Date(2019, 1, 10, 0, 0, 0, 0, time.UTC),
			DeactivatedOn: time.Time{},
			CustomerId:    1,
		},
	}

	newPlan := Subscription{
		Id:                    1,
		CustomerId:            1,
		MonthlyPriceInDollars: 4,
	}

	var noUsers []User

	It("works when the customer has no active users during the month", func() {
		Expect(BillFor("2019-01", &newPlan, &noUsers)).To(BeNumerically("~", 0.00, 0.005))
	})
	It("works when everything stays the same for a month", func() {
		Expect(BillFor("2019-01", &newPlan, &constantUsers)).To(BeNumerically("~", 8.0, 0.005))
	})
	It("works when a user is activated during the month", func() {
		Expect(BillFor("2019-01", &newPlan, &userSignedUp)).To(BeNumerically("~", 10.84, 0.005))
	})

})
