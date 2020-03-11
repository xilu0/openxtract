package main

import (
	"log"
	"net/http"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"flag"
)
var topic string
var page int
func init()  {
	flag.StringVar(&topic,"t","kubernetes","topic, default kubernetes")
	flag.StringVar(&topic,"topic","kubernetes","topic, default kubernetes")
	flag.IntVar(&page,"p",1,"page, default 1")
	flag.IntVar(&page,"page",1,"page, default 1")
}

type repository struct {
	Name string `json:"name"`
	Star string `json:"star"`
	Watch string `json:"watch"`
	Fork string `json:"fork"`
	Commits int32 `json:"commits"`
	Releases int32 `json:"releases"`
	Contributors int32 `json:"contributors"`
	License string `json:"license"`
	Description string `json:"description"`
	Topics []string `json:"topics"`
	Readme string `json:"readme"`
}

func Scrape(url string) {
	// Request the HTML page.
	res, err := http.Get(url)
	//res, err := http.Get("https://github.com/topics/sdk")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".application-main article h1.text-normal").Each(func(i int, s *goquery.Selection) {
			//fmt.Println(strings.Replace(strings.Replace(s.Text(), " ","", -1), "\n", "", -1))
			project := strings.Replace(strings.Replace(s.Text(), " ","", -1), "\n", "", -1)
			if v,_ := regexp.Match("^[0-9A-Za-z]+/[0-9A-Za-z]+$", []byte(project)); v == true {
				fmt.Println(project)
				GetDetail(project)
			}
	})

}

func GetDetail(project string) {
	res, err := http.Get("https://github.com/"+project)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	star := doc.Find(".js-social-count").Text()
	//var r1 repository
	fmt.Println("star:", FilterString(star))

	commits := doc.Find(".commits a span").Text()
	//var r1 repository
	fmt.Println("commits:", FilterString(commits))

	describe := doc.Find(".repository-content .text-gray-dark").Text()
	//var r1 repository
	fmt.Println("describe:", strings.Replace(describe,"\n", "", -1))

	//contributors := doc.Find("[data-hovercard-type='contributors']").Text()
	////var r1 repository
	//fmt.Println("contributors:", FilterString(contributors))

}
func FilterString(s string) string {
	return strings.Replace(strings.Replace(s, " ","", -1), "\n", "", -1)
	
}

// /html/body/div[4]/main/div[2]/div[2]/div/div[1]/article[1]/div[1]/div/div[1]/h1/a[2]
//*[@id="js-repo-pjax-container"]/div[1]/div/ul/li[3]/div/form[2]/a

func main() {
	flag.Parse()
	for i := 1; i <= page; i++ {
		url := "https://github.com/topics/" +topic+ "?page=" + strconv.Itoa(i)
		Scrape(url)
		fmt.Println(url)
	}
}
