package main

import (
	"io/ioutil"
	"net/http"
)

func FetchNewStories() ([]byte, error) {
	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/newstories.json")
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	return b, nil
}

func reverse(s []byte) []byte {
  for i := 0; i < len(s)/2; i++ {
    j := len(s) - 1 - i
    s[i], s[j] = s[j], s[i]
  }
  return s
}

func main() {
	stories, err := FetchNewStories()
	if err != nil {
		panic(err)
	}

  reverse(stories)

	for _, story := range stories[:10] {
		resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + string(story) + ".json")
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		println(string(b))
	}
}
