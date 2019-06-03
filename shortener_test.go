package main


import "testing"


func TestShortenAndResolve(t *testing.T) {
    t.Log("Shortening and resolving url")
    shortener := UrlShortener{make(map[string]string), "https://stupidurl.com", 4}
    url := "https://www.youtube.com/watch?v=3H0b5qUfDug"
    urlShort := shortener.Shorten(url)

    if urlResolved := shortener.Resolve(urlShort); urlResolved != url {
        t.Errorf("Expected %s, got %s instead", url, urlResolved)
    }
}


func TestRandomUrl(t *testing.T) {
    t.Log("Trying to resolve random url")
    shortener := UrlShortener{make(map[string]string), "https://stupidurl.com", 4}
    randomUrl := "https://foobar.com/gre2332"

    if urlResolved := shortener.Resolve(randomUrl); urlResolved != "" {
        t.Errorf("Expected empty string, got %s instead", urlResolved)
    }
}
