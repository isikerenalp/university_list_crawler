# Yök Atlas Crawler

YÖK Lisans Atlası (2016-2017-2018-2019) ve YÖK Önlisans Atlası (2016-2017-2018-2019) verilerinin `Üniversite Adı - Fakülte - Bölüm` şeklinde alınıp json olarak tutulması için hazırlanmıştır. Bu kapsamda iki adet `GO` kodu yazılmış olup birinde [goquery](https://github.com/PuerkitoBio/goquery) paketi ile diğerinde ise standart `GO` kütüphaneleri kullanılmıştır.   

# `1-) Goquery Crawler`

    go get -u github.com/PuerkitoBio/goquery
    git clone https://github.com/isikerenalp/university_list_crawler.git
    cd university_list_crawler/goquery_crawler
    go run main.go > list.json

# `2-) Simple Crawler`

    git clone https://github.com/isikerenalp/university_list_crawler.git
    cd university_list_crawler/simple_crawler
    go run main.go > list.json
