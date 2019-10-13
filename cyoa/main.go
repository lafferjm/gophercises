package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	"gophercises/cyoa/story"
)

func main() {
	storyFile := flag.String("story", "gopher.json", "JSON file containing the story details")
	flag.Parse()

	storyData, err := ioutil.ReadFile(*storyFile)
	if err != nil {
		panic(err)
	}

	storyDetails, err := story.LoadStory(storyData)
	if err != nil {
		panic(err)
	}

	fmt.Println(storyDetails["intro"].Options[0])
}
