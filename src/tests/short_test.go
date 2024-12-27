package genshort_test

import (
	"fmt"
	"short-url-go/genshort"
	"testing"
)

func TestGenShort(t *testing.T) {
	size := 1000000
	arrShortsURLs := make([]string, size)
	var shortURL string
	for i := 0; i < size; i++ {
		shortURL, _ = genshort.ShortURL("")
		arrShortsURLs[i] = shortURL
	}
	fmt.Println("hello", arrShortsURLs)
	for j, k := range arrShortsURLs[1:] {
		if arrShortsURLs[0] == k {
			t.Errorf("values is equal, first token=%s, duplacated token=%s in %d index", arrShortsURLs[0], k, j)
		}
	}

}
