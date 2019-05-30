package main


import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)


type UrlShortener struct {
    storage map[string]string
    service_url string
    key_length int
}


func randomString(length int) string {
    rand.Seed(time.Now().UnixNano())
    chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
    var builder strings.Builder
    for i := 0; i < length; i++ {
        builder.WriteRune(chars[rand.Intn(len(chars))])
    }
    return builder.String()
}


func (u *UrlShortener) Shorten(url string) string {
    for {
        key := randomString(u.key_length)
        _, ok := u.storage[key]
        if !ok {
            u.storage[key] = url
            return u.service_url + "/" + key
        }
        time.Sleep(2 * time.Millisecond)
    }
}


func (u *UrlShortener) Resolve(url string) string {
    url_parts := strings.Split(url, "/")
    key := url_parts[len(url_parts)-1]
    return u.storage[key]
}


func main() {
    shortener := UrlShortener{make(map[string]string), "https://stupidurl.com", 4}

    url1 := "https://yandex.ru/search/?lr=213&clid=2186621&text=what%20color%20is%20red"
    url2 := "https://www.youtube.com/watch?v=3H0b5qUfDug"
    url3 := "http://conjugator.reverso.net/conjugation-english-verb-not%20to%20duck.html"

    url1_sh := shortener.Shorten(url1)
    url2_sh := shortener.Shorten(url2)
    url3_sh := shortener.Shorten(url3)

    random_url := "https://foobar.com/gre2332"

    fmt.Println("Original url 1:", url1)
    fmt.Println("Original url 2:", url2)
    fmt.Println("Original url 3:", url3, "\n")

    fmt.Println("Shortened url 1:", url1_sh)
    fmt.Println("Shortened url 2:", url2_sh)
    fmt.Println("Shortened url 3:", url3_sh, "\n")

    fmt.Println("Resolved url 1:", shortener.Resolve(url1_sh))
    fmt.Println("Resolved url 2:", shortener.Resolve(url2_sh))
    fmt.Println("Resolved url 3:", shortener.Resolve(url3_sh), "\n")
    fmt.Println("Trying to resolve random url", random_url, ": '" + shortener.Resolve(random_url) + "'")

}
