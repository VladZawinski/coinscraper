package main
import (
	"fmt"
	"encoding/json"
	"github.com/gocolly/colly"
)

type Coin struct {
	Name string
	Price string
}

func main() {
	c := colly.NewCollector()

	c.OnHTML(".lcw-table tr", func(e *colly.HTMLElement) {
		name := e.ChildText(".filter-item-name")
		price := e.ChildText(".main-price")

		eachCoin := Coin {
			Name: name,
			Price: price,
		}
        
		result, err := json.Marshal(eachCoin)
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		} else {
			fmt.Println(string(result))
		}
    })
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})
	
    c.Visit("https://www.livecoinwatch.com")
}