package story

import (
	"encoding/json"
)

type StoryDetails struct {
	Title   string
	Story   []string
	Options []StoryOptions
}

type StoryOptions struct {
	Text string
	Arc  string
}

func LoadStory(storyData []byte) (map[string]StoryDetails, error) {
	var story map[string]StoryDetails
	err := json.Unmarshal(storyData, &story)
	if err != nil {
		return nil, err
	}

	return story, nil
}
