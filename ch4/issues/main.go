package main

import (
	"fmt"
	"github.com/igor-kupczynski/gopl/ch4/github"
	"log"
	"os"
	"time"
)

const (
	month = time.Hour * 24 * 30
	year  = time.Hour * 24 * 365
)

type AgeLevel struct {
	Age      time.Duration
	AgeDesc  string
	LessThan bool
	Items    []*github.Issue
}

var levels = [...]*AgeLevel{
	{Age: month, AgeDesc: "month", LessThan: true},
	{Age: year, AgeDesc: "year", LessThan: true},
	{Age: year, AgeDesc: "year", LessThan: false},
}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// Segregate issues into levels by age
	now := time.Now()
	for _, item := range result.Items {
		for _, l := range levels {
			if item.CreatedAt.Add(l.Age).After(now) == l.LessThan {
				l.Items = append(l.Items, item)
				break
			}
		}
	}

	// Print each level
	for _, l := range levels {
		// Header
		var lt = "less than"
		if !l.LessThan {
			lt = "more than"
		}
		fmt.Printf("> %d issues %s a %s old", len(l.Items), lt, l.AgeDesc)

		// Issues
		for _, item := range l.Items {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}

		fmt.Println()
	}
}
