package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
	"unicode"
)

func main() {
	wg := new(sync.WaitGroup)
	sem := make(chan bool, 23)

	loc, err := time.LoadLocation("America/Vancouver")

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, path := range os.Args {
		wg.Add(1)
		go touch(sem, wg, path, loc)
	}

	wg.Wait()
}

func touch(sem chan bool, wg *sync.WaitGroup, path string, loc *time.Location) {
	sem <- true
	defer func() { <-sem }()
	defer func() { wg.Done() }()

	name := filepath.Base(path)
	extn := filepath.Ext(name)
	base := name[0 : len(name)-len(extn)]

	// NOTES(cixtor): TYPE_yyyymmdd_hhiiss.ext
	parts := strings.SplitN(base, "_", 3)

	if len(parts) < 3 {
		fmt.Printf("invalid filename `%s`\n", name)
		return
	}

	if len(parts[1]) != 8 {
		fmt.Printf("invalid filename date `%s`\n", name)
		return
	}

	if !isTimePortion(parts[2]) {
		parts[2] = "090000"
	}

	// NOTES(cixtor): create fake datetime string.
	date := fmt.Sprintf(
		"%s-%s-%sT%s:%s:%s-08:00",
		parts[1][0:4], // year
		parts[1][4:6], // month
		parts[1][6:8], // day
		parts[2][0:2], // hours
		parts[2][2:4], // minutes
		parts[2][4:6], // seconds
	)

	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		fmt.Printf("invalid datetime in `%s`: %s\n", name, err)
		return
	}

	fmt.Printf("%s >>> %s\n", name, t)

	// NOTES(cixtor): change access and modification times of the named file.
	if err := os.Chtimes(path, t, t); err != nil {
		fmt.Printf("cannot change time `%s`: %s\n", name, err)
		return
	}
}

func isTimePortion(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
