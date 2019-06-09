package main


import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)


type UrlShortener struct {
    storage map[string]string
    serviceUrl string
    keyLength int
}


func randomString(length int) string {
    chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    var builder strings.Builder
    for i := 0; i < length; i++ {
        builder.WriteRune(chars[rand.Intn(len(chars))])
    }
    return builder.String()
}


func (self *UrlShortener) Shorten(url string) string {
    for {
        key := randomString(self.keyLength)
        _, ok := self.storage[key]
        if !ok {
            self.storage[key] = url
            return self.serviceUrl + "/" + key
        }
    }
}


func (self *UrlShortener) Resolve(url string) string {
    urlParts := strings.Split(url, "/")
    key := urlParts[len(urlParts)-1]
    return self.storage[key]
}


func init() {
    rand.Seed(time.Now().UnixNano())
}


func main() {
    shortener := UrlShortener{make(map[string]string), "https://stupidurl.com", 4}

    url1 := "https://yandex.ru/search/?lr=213&clid=2186621&text=what%20color%20is%20red"
    url2 := "https://www.youtube.com/watch?v=3H0b5qUfDug"
    url3 := "http://conjugator.reverso.net/conjugation-english-verb-not%20to%20duck.html"

    urlShort1 := shortener.Shorten(url1)
    urlShort2 := shortener.Shorten(url2)
    urlShort3 := shortener.Shorten(url3)

    randomUrl := "https://foobar.com/gre2332"

    fmt.Println("Original url 1:", url1)
    fmt.Println("Original url 2:", url2)
    fmt.Println("Original url 3:", url3)
    fmt.Println()
    fmt.Println("Shortened url 1:", urlShort1)
    fmt.Println("Shortened url 2:", urlShort2)
    fmt.Println("Shortened url 3:", urlShort3)
    fmt.Println()
    fmt.Println("Resolved url 1:", shortener.Resolve(urlShort1))
    fmt.Println("Resolved url 2:", shortener.Resolve(urlShort2))
    fmt.Println("Resolved url 3:", shortener.Resolve(urlShort3))
    fmt.Println()
    fmt.Println("Trying to resolve random url", randomUrl, ": '" + shortener.Resolve(randomUrl) + "'")
}
