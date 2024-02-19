package scraper

import (
	"github.com/gocolly/colly"
	"website/internal/models"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)


func Scrape(source string)[]models.Omnibus {
	var OmnibusList []models.Omnibus
	// startTime := time.Now()



	c:= colly.NewCollector(
		colly.AllowedDomains("www.instocktrades.com", "instocktrades.com", "amazon.sg","www.amazon.sg"),
	)

	c.OnHTML("div[class=title]", func(e *colly.HTMLElement) {
		name := e.ChildText("a")
		if strings.Contains(name, "Omnibus") && strings.Contains(name, "HC") {
			// fmt.Println(name)
			link := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
			// fmt.Println(link)
			c.Visit(link)
			
		} 

	})


	c.OnHTML("div[class=productcontent]", func(e *colly.HTMLElement) {
		// Create an instance of Omnibus
		Omnibus := models.Omnibus{}
		Omnibus.ISTUrl = e.Request.URL.String()
		Omnibus.Name = e.ChildText("h1")
		selection := e.DOM
		info := selection.Find("div.prodinfo")
		infoNodes := info.Children().Nodes

		for i := 0; i < len(infoNodes); i++ {
			infoNode := selection.FindNodes(infoNodes[i]).Text()
			values := strings.Split(infoNode, ":")
			if len(values) > 1 {
				label := strings.TrimSpace(values[0])
				value := strings.TrimSpace(values[1])

				switch label {
				case "Publisher":
					Omnibus.Publisher = value
				case "Page Count":
					Omnibus.PageCount, _ = strconv.Atoi(value)
				case "UPC":
					Omnibus.UPC = value
				}
			}
		}

		Omnibus.DateCreated = time.Now().Format("2006-01-02")
		Omnibus.Sale = false

		pricing := selection.Find("div.pricing")
		pricingNodes := pricing.Children().Nodes

		for i := 0; i < len(pricingNodes); i++ {
			pricingNode := selection.FindNodes(pricingNodes[i]).Text()
			if strings.Contains(pricingNode, ":") {
				values := strings.Split(pricingNode, ":")
				if len(values) > 1 {
					label := strings.TrimSpace(values[0])
					value := strings.TrimSpace(values[1])
					v := strings.Replace(value, "$", "", -1)


					switch label {
					case "Was":
						price,_ := strconv.ParseFloat(v, 64)
						Omnibus.Price = float32(price)
					case "IST Price":
						current, _ := strconv.ParseFloat(v, 64)
						Omnibus.Current = float32(current)

					}
				} 
			} else {
				re := regexp.MustCompile(`(\d+)`)
				matches := re.FindAllString(pricingNode, -1)
				Omnibus.Saving, _ = strconv.Atoi(matches[0])

			}
		}

		// fmt.Println(Omnibus)
		OmnibusList = append(OmnibusList, Omnibus)
	})

	c.OnHTML("a.btn.hotaction", func(e *colly.HTMLElement) {
		link := e.Request.AbsoluteURL(e.Attr("href"))

		c.Visit(link)
			

	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(source)

	// endTime := time.Now()

	// elapsedTime := endTime.Sub(startTime)

	// Print the elapsed time
	// fmt.Printf("Program execution time: %s\n", elapsedTime)
	// OmnibusJson, err := json.Marshal(OmnibusList)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	
	return OmnibusList
}
	