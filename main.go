package main

import (
	"io/ioutil"
	"net/http"
)

// create a function that fetches the latest hacker news stories and returns json
func FetchNewStories() ([]byte, error) {
	// fetch the latest hacker news stories
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/newstories.json")
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return b, nil
}

func main() {
	stories, err := FetchNewStories()
	if err != nil {
		panic(err)
	}

  // reverse the stories
  for i, j := 0, len(stories)-1; i < j; i, j = i+1, j-1 {
    stories[i], stories[j] = stories[j], stories[i]
  }

	for _, story := range stories[:10] {
    println(string(story))
		resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + string(story) + ".json")
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		println(string(b))
	}
}
