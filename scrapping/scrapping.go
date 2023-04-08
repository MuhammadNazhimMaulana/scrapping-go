package scrapping

import (
	"encoding/csv"
	"fmt"
	"log"

	"os"

	"github.com/gocolly/colly"
)

type Words struct {
	url   string
	image string
}

func Test() {
	// Initialization
	var words []Words

	c := colly.NewCollector()

	// Called Before Visiting the Web
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Mengunjungi", r.URL)
	})

	// Called Whenever error is occured
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Oops: ", err)
	})

	// First Example
	// c.OnHTML("div.mdk-body-paragraph", func(e *colly.HTMLElement) {
	// 	fmt.Printf("Product Name: %s \n", e.Text)
	// })

	// c.Visit("https://www.merdeka.com/trending/80-kata-kata-mutiara-tentang-hidup-sehat-cocok-untuk-sambut-hari-kesehatan-nasional-kln.html")

	// Second Example
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		word := Words{}

		// scraping the data of interest
		word.url = e.ChildAttr("a", "href")
		word.image = e.ChildAttr("img", "src")
		// word.image = e.ChildText("h2 .wp-block-heading")

		// adding the product instance with scraped data to the list of products
		words = append(words, word)
		fmt.Printf("Hasil: %s \n", word.image)
	})

	c.Visit("https://scrapeme.live/shop/")

	// opening the CSV file
	file, err := os.Create("results.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	// // initializing a file writer
	writer := csv.NewWriter(file)

	// writing the CSV headers
	headers := []string{
		"url",
		"image",
	}
	writer.Write(headers)

	// writing each Pokemon product as a CSV row
	for _, word := range words {
		// converting a PokemonProduct to an array of strings
		record := []string{
			word.url,
			word.image,
		}

		// adding a CSV record to the output file
		writer.Write(record)
	}
	defer writer.Flush()
}
