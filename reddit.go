// Package reddit implements a basic client for the Reddit API.
package reddit

import (
	"net/http"
	"encoding/json"
    "fmt"
    "errors"
)

func (i Item) String() string {
    com := ""
    switch i.Comments {
    case 0:
        // nothing
    case 1:
        com = " (1 comment)"
    default:
        com = fmt.Sprintf(" (%d comments)", i.Comments)
    }
    return fmt.Sprintf("%s%s\n%s", i.Title, com, i.URL)
}

// Item describes a Reddit item.
type Item struct {
		Title string
		URL   string
        Comments int `json:"num_comments"`
	}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Get fetches the most recent Items posted to the specified subreddit.
func Get(reddit string) ([]Item, error) {
    url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    // Executes at the end of function call.
    defer resp.Body.Close()
    if resp.StatusCode != http.StatusOK {
        // resp.Status is a string and return type is error.
        return nil, errors.New(resp.Status)
    }
    r := new(Response)
    err = json.NewDecoder(resp.Body).Decode(r)
    if err != nil {
        return nil, err
    }
    items := make([]Item, len(r.Data.Children))
    for i, child := range r.Data.Children {
        items[i] = child.Data
    }
    return items, nil
}
