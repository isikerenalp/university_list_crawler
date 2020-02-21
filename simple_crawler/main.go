package main

import (
	"fmt"
	"net/http"
	"os"
  "io/ioutil"
	"strings"
	"strconv"
	"regexp"
)

func main() {
	// 1000 - 2000 Arası Devlet Üniversitesi && 2000-3000 Arası Özel Üniversite && 3000-4000 KKTC
	for counter := 1000; counter < 4000; counter++ {

		url_counter := strconv.Itoa(counter)
	  baseURL := "https://yokatlas.yok.gov.tr/onlisans-univ.php?u=" + url_counter // Önlisans Üniversite-Fakülte-Bölüm verisi
		// baseURL := "https://yokatlas.yok.gov.tr/lisans-univ.php?u=" + url_counter  // Lisans Üniversite-Fakülte-Bölüm verisi
	  response, err := http.Get(baseURL)
		checkErr(err)
		defer response.Body.Close()
		body, _ := ioutil.ReadAll(response.Body)
		x := string(body)

		// Üniversite Adı
		i := strings.Index(x, "'ndeki")
		ii := strings.Index(x, "<title>")
		var universityName string
		if i > -1 && ii > -1 {
			chars := x[ii+7:i]
			if chars == "" {
				continue
			} else {
				universityName = chars
			}
		} else {
			fmt.Println("Not found")
			fmt.Println(x)
		}

		// Bölüm ve Fakülte Bilgisi
		re := regexp.MustCompile("<div style=\"overflow: hidden; text-overflow: ellipsis; white-space: nowrap;width:80%\">(.|\n)*?</font>")
		unvara := re.FindAllString(x, -1)
		if unvara == nil {
			fmt.Println("Yok.")
		} else {
			for _, unvara1 := range unvara {
				unis := string(unvara1)
				result := strings.ReplaceAll(unis, "<div style=\"overflow: hidden; text-overflow: ellipsis; white-space: nowrap;width:80%\">", "")
				result2 := strings.ReplaceAll(result, "</div>", "")
				result3 := strings.ReplaceAll(result2, "<font color=\"#FFFFFF\">", "")
				result4 := strings.ReplaceAll(result3, "</font>", "")
				result5 := strings.ReplaceAll(result4, "<small>", "")
				result6 := strings.ReplaceAll(result5, "(", "")
				result7 := strings.ReplaceAll(result6, ")", "")
				s := strings.Split(result7, "\n")
				department := s[0]
				faculty := s[1]
				fmt.Printf("{\n\"ID\" : " + url_counter + ",\n" + "\"universityName\" : " + "\"" + universityName + "\"" + ",\n" )
				fmt.Printf("\"faculty\" : " + "\"" + faculty + "\"" + ",\n" + "\"department\" : " + "\"" + department + "\"" + "\n},\n")
			}
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
