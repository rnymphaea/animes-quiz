package parser

import (
	"anime/models"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

func parseAnimes(start, end int) []models.Anime {
	animes := make([]models.Anime, 0)
	for i := start; i < end; i++ {
		URL := "https://shikimori.one/animes/"
		URL += strconv.Itoa(i)
		resp, err := http.Get(URL)

		if err != nil {
			log.Fatal("Error: cannot connect to given URL")
		}

		defer resp.Body.Close()

		if err != nil {
			log.Fatal("Error: cannot read data from response body!")
		}

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		title := doc.Find("h1").Text()
		desc := doc.Find("div.b-text_with_paragraphs").Text()
		// rating := doc.Find("div.score-value.score-9").Text()

		if len(desc) > 0 {
			info := strings.Split(doc.Find("div.b-entry-info").Text(), "    ")
			var ep uint16
			var typee string
			var genres string
			for _, i := range info {
				w := strings.Split(i, ": ")
				if strings.Compare(w[0], "Эпизоды") == 0 {
					eps, _ := strconv.Atoi(w[1])
					ep = uint16(eps)
				} else if strings.Compare(w[0], "  Тип") == 0 {
					typee = w[1]
				} else if strings.Compare(w[0], "Жанры") == 0 {
					genres = w[1]
				}
			}
			anime := &models.Anime{
				Title:    title,
				Desc:     desc,
				Episodes: ep,
				Genre:    genres,
				Status:   "",
				Type:     typee,
			}
			animes = append(animes, *anime)
		}
	}
	return animes
}

func GetAnimesFromSite(countPages int) []models.Anime {
	var animes []models.Anime
	alpha := countPages / 25
	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			an := parseAnimes(num*alpha, (num+1)*alpha)
			mu.Unlock()

			animes = append(animes, an...)
		}(i)
	}
	wg.Wait()
	return animes
}
