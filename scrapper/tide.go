package scrapper

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Tide struct {
	Time time.Time `json:"time"`
	Data int       `json:"data"`
}

func ScrapTides(c *colly.Collector, url string) (map[string][]Tide, error) {
	var err error
	tides := map[string][]Tide{}
	date := strings.Split(url, "=")
	l := len(date)
	if l > 0 {
		strlen := len(date[1])
		if date[1][strlen-1] == '#' {
			date[1] = date[1][:strlen-1]
		}
	} else {
		return nil, errors.New("url not valid")
	}

	// get year, month, day
	t := strings.Split(date[1], "-")
	if len(t) != 3 {
		return nil, errors.New("url not valid")
	}

	c.OnHTML("script", func(h *colly.HTMLElement) {
		content := h.Text

		for {
			if len(content) == 0 {
				break
			}

			startIDX := strings.Index(content, `{"name":`)
			if startIDX == -1 {
				return
			}

			data := content[startIDX:]
			content = data
			endIDX := strings.Index(data, "]}")

			if endIDX == -1 {
				return
			}
			data = data[:endIDX+1]

			content = content[len(data):]

			s := strings.Split(data, `:`)
			if len(s) > 0 {
				keys := strings.Split(s[1], ",")
				key := keys[0][1 : len(keys[0])-1]

				values := strings.Split(s[len(s)-1], ",")
				values[0] = values[0][1:]
				values[len(values)-1] = values[len(values)-1][:len(values[len(values)-1])-1]

				pasutData := make([]Tide, len(values))
				for i := range values {
					v := 0
					if values[i] == "null" {
						continue
					}

					v, err = strconv.Atoi(values[i])
					if err != nil {
						log.Panic(err)
					}
					sourceTime := fmt.Sprintf("%v-%v-%v %v:00", t[0], t[1], t[2], i+1)
					parseTime, errParse := time.Parse(time.DateTime, sourceTime)
					if errParse != nil {
						err = errParse
					}
					pasutData[i] = Tide{
						Time: parseTime,
						Data: v,
					}
				}

				tides[key] = pasutData
			}
		}

	})

	if err != nil {
		return nil, err
	}

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit(url)

	if len(tides) == 0 {
		return nil, errors.New("tidak ada yang didapat")
	}

	return tides, nil
}
