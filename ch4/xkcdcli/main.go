package main

import (
	"encoding/json"
	"fmt"
	"github.com/igor-kupczynski/gopl/ch4/xkcd"
	"log"
	"os"
	"strings"
)

const store = "xkcd.json"

var results = make(chan *xkcd.Comic)

var (
	reindex bool
	search  bool
	query   []string
)

func main() {
	for idx, x := range os.Args[1:] {
		if x == "--reindex" {
			reindex = true
		}
		if x == "--search" {
			search = true
			for _, q := range os.Args[idx+2:] {
				query = append(query, strings.ToLower(q))
			}
			break
		}
	}

	if reindex {
		log.Printf("Reindexing...\n")
		n := maxNum()
		index := buildIndex(n)
		saveIndex(index)
	}

	if search {
		index := loadIndex()
		log.Printf("Searching for any of %v...\n", query)
		r := find(index, query)
		log.Printf("Found %d matches:\n", len(r))
		for _, c := range r {
			println()
			printXkcd(c)
		}
	}
}

func maxNum() int {
	curr, err := xkcd.DefaultClient.Current()
	if err != nil {
		log.Printf("Can't process current: %v\n", err)
		return -1
	}
	return curr.Num
}

func buildIndex(maxNum int) map[int]*xkcd.Comic {
	var index = make(map[int]*xkcd.Comic, 3000)
	for i := 1; i <= maxNum; i++ {
		go func(num int) {
			c, err := xkcd.DefaultClient.Get(num)
			if err != nil {
				log.Printf("Can't process [XKCD#%d]: %v\n", num, err)
				results <- nil
			}
			results <- c
		}(i)
	}
	for i := 1; i <= maxNum; i++ {
		c := <-results
		if c == nil {
			continue
		}
		index[c.Num] = c
	}
	return index
}

func saveIndex(index map[int]*xkcd.Comic) {
	f, err := os.Create(store)
	if err != nil {
		log.Printf("Can't create file to save index: %v\n", err)
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	for _, c := range index {
		if err := enc.Encode(c); err != nil {
			log.Printf("Can't encode [%v]: %v\n", c, err)
			return
		}
	}
}

func loadIndex() map[int]*xkcd.Comic {
	f, err := os.Open(store)
	if err != nil {
		log.Printf("Can't open file to load index: %v\n", err)
		return nil
	}
	defer f.Close()

	index := make(map[int]*xkcd.Comic, 3000)

	dec := json.NewDecoder(f)
	for dec.More() {
		var c xkcd.Comic
		if err := dec.Decode(&c); err != nil {
			log.Printf("Can't decode: %v\n", err)
			return index
		}
		index[c.Num] = &c
	}
	return index
}

func find(index map[int]*xkcd.Comic, query []string) []*xkcd.Comic {
	r := make([]*xkcd.Comic, 0)
	for _, comic := range index {
		title := strings.ToLower(comic.Title)
		for _, q := range query {
			if strings.Contains(title, q) {
				r = append(r, comic)
				break
			}
		}
	}
	return r
}

func printXkcd(c *xkcd.Comic) {
	fmt.Printf(
		":: XKCD#%d %s\nUrl:\t%s\nScript:\t%s\nAlt:\t%s\n",
		c.Num, c.Title, c.Link(), c.Transcript, c.Alt)
}
