package scrapper

// import (
// 	"fmt"
// 	"log"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/gocolly/colly"
// )

// type Weather struct {
// 	Time          time.Time
// 	Image         string
// 	Temperature   int
// 	State         string
// 	Humidity      int
// 	WindSpeed     int
// 	WindDirection string
// 	Visibility    float64
// 	LastUpdate    time.Time
// }

// func ScrapWeather(c *colly.Collector, link string) []Weather {
// 	weathers := []Weather{}

// 	c.OnHTML("div.md\\:flex.items-center.gap-6", func(e *colly.HTMLElement) {
// 		lastUpdatedAtr := e.ChildText("time.text-lg.md\\:text-xl")
// 		temperatureAtr := e.ChildText("p.font-bold")
// 		stateAtr := e.ChildText("p.text-black-primary")

// 		attrs := []string{}
// 		e.ForEach("span.text-black-primary", func(i int, h *colly.HTMLElement) {
// 			attrs = append(attrs, h.Text)
// 		})

// 		winDirection := attrs[2]
// 		attrs[2] = attrs[3]
// 		attrs[3] = winDirection

// 		sAttr := ""
// 		for i := range attrs {
// 			sAttr += attrs[i] + " "
// 		}

// 		sAttr = sAttr[:len(sAttr)-1]

// 		l := strings.Split(lastUpdatedAtr, " ")
// 		lastUpdate := strings.Join(l[3:], " ")
// 		hours := strings.Join(l[:3], " ")

// 		w := extractAttr(hours, "", lastUpdate, temperatureAtr, stateAtr, sAttr)
// 		weathers = append(weathers, w)

// 	})

// 	c.OnHTML("div.bg-white.pb-10", func(e *colly.HTMLElement) {
// 		lastUpdatedAtr := e.ChildText("span.mx-auto.md\\:w-max")
// 		hours := []string{}

// 		e.ForEach("div.p-5.md\\:p-8.rounded-2xl", func(i int, h *colly.HTMLElement) {
// 			hours = append(hours, h.ChildText("h4.text-base"))
// 		})

// 		day := e.ChildText("button")[:6]
// 		temperatur := make([]string, len(hours))
// 		img := make([][]string, len(hours))
// 		states := make([]string, len(hours))
// 		attrs := make([]string, len(hours))

// 		e.ForEach("div.swiper-slide", func(i int, h *colly.HTMLElement) {
// 			temperatur[i] = h.ChildText("p.font-bold")
// 			img[i] = h.ChildAttrs("path", "d")
// 			states[i] = h.ChildText("p.text-sm")
// 			attrs[i] = h.ChildText("p.text-black-primary") + " " + h.ChildText("span.text-black-primary")
// 		})

// 		for i := range hours {
// 			w := extractAttr(day, hours[i], lastUpdatedAtr, temperatur[i], states[i], attrs[i])
// 			weathers = append(weathers, w)
// 		}
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	// Start scraping
// 	c.Visit(link)

// 	return weathers
// }

// func extractAttr(t, hours, lastUpdatedAttr, temperaturAttr, stateAtr, attrs string) Weather {
// 	var err error
// 	trimedTempt := strings.Split(temperaturAttr, " ")
// 	var temperature int
// 	if len(trimedTempt) > 0 {
// 		temperature, err = strconv.Atoi(string(trimedTempt[0]))
// 		if err != nil {
// 			log.Fatal(err, "p")
// 		}
// 	}

// 	var humidity, windSpeed int
// 	var windDirection string
// 	var visibility float64
// 	if len(attrs) > 0 {
// 		trimedHumididy := strings.Split(attrs, "%")
// 		if len(trimedHumididy) > 0 {
// 			humidity, err = strconv.Atoi(trimedHumididy[0])
// 			if err != nil {
// 				log.Fatal(err, "e")
// 			}
// 		}

// 		attrs = attrs[len(trimedHumididy)+2:]
// 		trimedWindSpeed := strings.Split(attrs, " ")
// 		if len(trimedTempt) > 0 {
// 			windSpeed, err = strconv.Atoi(trimedWindSpeed[0])
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 		}

// 		windDirection = trimedWindSpeed[len(trimedWindSpeed)-1]

// 		v, err := strconv.ParseFloat(string(trimedWindSpeed[3]), 64)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		// if arrow to left will lower than real value, otherwise
// 		if trimedWindSpeed[2] == "<" {
// 			visibility = v - 0.1
// 		}
// 		if trimedWindSpeed[2] == ">" {
// 			visibility = v + 0.1
// 		}

// 	}

// 	var day, month, year, hour string
// 	trimedLastUpdated := strings.Split(lastUpdatedAttr, " ")
// 	if len(trimedLastUpdated) > 0 {
// 		day = trimedLastUpdated[1]
// 		month = changeMonth(trimedLastUpdated[2])
// 		year = trimedLastUpdated[3][:len(trimedLastUpdated[3])-1]
// 		hour = trimedLastUpdated[4]
// 	}

// 	sourceTime := fmt.Sprintf("%v-%v-%v %v:00", year, month, day, hour)
// 	lastUpdate, err := time.Parse(time.DateTime, sourceTime)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	img := setImg(stateAtr)
// 	if img == "cerah" || img == "cerah-berawan" || img == "udara-kabur" {
// 		if lastUpdate.Hour() >= 18 || lastUpdate.Hour() <= 5 {
// 			img += "-malam"
// 		} else {
// 			img += "-siang"
// 		}
// 	}

// 	if strings.Contains(t, "Saat ini") {
// 		yy, mm, dd := time.Now().Date()
// 		h, m, s := time.Now().Clock()

// 		sd := ""
// 		if dd < 10 {
// 			sd += "0" + strconv.Itoa(dd)
// 		}
// 		sh := ""
// 		if h < 10 {
// 			sh += "0" + strconv.Itoa(h)
// 		}

// 		sm := ""
// 		if m < 10 {
// 			sm += "0" + strconv.Itoa(m)
// 		}

// 		ss := ""
// 		if s < 10 {
// 			ss += "0" + strconv.Itoa(s)
// 		}

// 		t = fmt.Sprintf("%v-%v-%v %v:%v:%v", yy, changeMonth(mm.String()[:3]), sd, h, m, s)
// 	} else {
// 		day := strings.Split(t, " ")
// 		hour := strings.Split(hours, " ")
// 		splitHour := strings.Split(hour[0], ".")

// 		if len(day) > 0 {
// 			t = fmt.Sprintf("%v-%v-%v %v:%v:00", time.Now().Year(), changeMonth(day[1]), day[0], splitHour[0], splitHour[1])
// 		}

// 	}

// 	setTime, err := time.Parse(time.DateTime, t)
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	weather := Weather{
// 		Time:          setTime,
// 		Image:         img + ".svg",
// 		Temperature:   temperature,
// 		State:         stateAtr,
// 		Humidity:      humidity,
// 		WindSpeed:     windSpeed,
// 		WindDirection: windDirection,
// 		Visibility:    visibility,
// 		LastUpdate:    lastUpdate,
// 	}

// 	return weather
// }

// type Pasut struct {
// 	Name string `json:"name"`
// 	Data []int  `json:"data"`
// }

// func main() {
// 	// mainDomain := "https://www.bmkg.go.id/"
// 	// weather := "cuaca/prakiraan-cuaca/33.21.04.2011"

// 	c := colly.NewCollector()
// 	pasut := "https://pasut.maritimsemarang.com/?tanggal=2025-03-05"
// 	var err error
// 	pst := []Pasut{}

// 	c.OnHTML("script", func(h *colly.HTMLElement) {
// 		content := h.Text

// 		for {
// 			if len(content) == 0 {
// 				break
// 			}

// 			startIDX := strings.Index(content, `{"name":`)
// 			if startIDX == -1 {
// 				return
// 			}

// 			data := content[startIDX:]
// 			content = data
// 			endIDX := strings.Index(data, "]}")

// 			if endIDX == -1 {
// 				return
// 			}
// 			data = data[:endIDX+1]

// 			content = content[len(data):]

// 			var p Pasut
// 			s := strings.Split(data, `:`)
// 			if len(s) > 0 {
// 				keys := strings.Split(s[1], ",")
// 				key := keys[0][1 : len(keys[0])-1]

// 				values := strings.Split(s[len(s)-1], ",")
// 				values[0] = values[0][1:]
// 				values[len(values)-1] = values[len(values)-1][:len(values[len(values)-1])-1]

// 				pasutData := make([]int, len(values))
// 				for i := range values {
// 					v := 0
// 					if values[i] == "null" {
// 						continue
// 					}

// 					v, err = strconv.Atoi(values[i])
// 					if err != nil {
// 						log.Panic(err)
// 					}

// 					pasutData[i] = v
// 				}
// 				p.Name = key
// 				p.Data = pasutData
// 			}
// 			pst = append(pst, p)
// 		}

// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	c.Visit(pasut)
// 	fmt.Println(pst)
// 	// Instantiate default collector

// 	// fmt.Println(ScrapWeather(c, mainDomain+weather))
// }

// func setImg(s string) string {
// 	s = strings.ToLower(s)
// 	p := strings.Split(s, " ")
// 	if len(p) > 1 {
// 		s = p[0] + "-" + p[1]
// 	}

// 	return s
// }

// func changeMonth(month string) string {
// 	switch month {
// 	case "Jan":
// 		return "01"
// 	case "Feb":
// 		return "02"
// 	case "Mar":
// 		return "03"
// 	case "Apr":
// 		return "04"
// 	case "Mei":
// 		return "05"
// 	case "Jun":
// 		return "06"
// 	case "Jul":
// 		return "07"
// 	case "Agu":
// 		return "08"
// 	case "Sep":
// 		return "09"
// 	case "Okt":
// 		return "10"
// 	case "Nov":
// 		return "11"
// 	case "Des":
// 		return "12"
// 	}

// 	return ""
// }
