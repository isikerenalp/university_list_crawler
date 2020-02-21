package main

import (
  "fmt"
  "log"
  "strconv"
  "strings"
  "net/http"
  "github.com/PuerkitoBio/goquery"
)

func main() {
  // 1000 - 2000 Arası Devlet Üniversitesi && 2000-3000 Arası Özel Üniversite && 3000-4000 KKTC
	for xyz := 1000; xyz < 4000; xyz++ {

		url_counter := strconv.Itoa(xyz)
    // Request the HTML page.
    res, err := http.Get("https://yokatlas.yok.gov.tr/onlisans-univ.php?u="+url_counter) // Önlisans Üniversite-Fakülte-Bölüm
    // res, err := http.Get("https://yokatlas.yok.gov.tr/lisans-univ.php?u="+url_counter) // Lisans Üniversite-Fakülte-Bölüm
    checkErr(err)
    defer res.Body.Close()
    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    checkErr(err)

    // Find the University name
    doc.Find(".page-header").Each(func(i int, s *goquery.Selection) {
      band := s.Find("h1").Text()
      universityName := strings.ReplaceAll(band, "'ndeki Tüm Lisans Programları  (Alfabetik Sırada)", "")
      // Find the department and faculty
      doc.Find(".panel-title").Each(func(i int, s *goquery.Selection) {
        department := s.Find("div").Text()
        res1 := s.Find("small").Text()
        res2 := strings.ReplaceAll(res1, "(", "")
        faculty := strings.ReplaceAll(res2, ")", "")
        fmt.Printf("{\n\"ID\" : " + url_counter + ",\n" + "\"universityName\" : " + "\"" + universityName + "\"" + ",\n" )
        fmt.Printf("\"faculty\" : " + "\"" + faculty + "\"" + ",\n" + "\"department\" : " + "\"" + department + "\"" + "\n},\n")
      })
    })

  }
}

func checkErr(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

