package main

import (
	"flag"
)

func main() {
	var folder string
	email := "youssefashraf397@yahoo.com"
	flag.StringVar(&folder, "add", "", "add a new folder to scan for Git repositories")
	flag.StringVar(&email, "email", "youssefashraf397@yahoo.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
