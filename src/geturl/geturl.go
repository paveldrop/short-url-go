package geturl

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
	"strings"
	"time"
)

const (
	httpStr = "http://"
)

func GetUrl() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	if err != nil {
		return "", err
	}
	validateUrl(text)
	return text, nil
}

func validateUrl(str string) string {
	res, err := regexp.MatchString("^https://*|^http://*", str)
	fmt.Print("\n", res, "\n")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if res {
		if succes, err := getResponse(str); succes {
			fmt.Println(str)
			return str
		} else {
			fmt.Println(err)
			return ""
		}
	} else {
		str = httpStr + str
		fmt.Println(str)
		if succes, err := getResponse(str); succes {
			fmt.Println(str)
			return str
		} else {
			fmt.Println(err)
		}
	}
	time.Sleep(5 * time.Second)

	return str
}

func getResponse(url string) (bool, error) {
	client := &http.Client{}
	jar, err := cookiejar.New(nil)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	client.Jar = jar

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64; rv:133.0) Gecko/20100101 Firefox/133.0")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "en-US,en;q=0.5")
	// req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	// req.Header.Set("Connection", "keep-alive")
	// req.Header.Set("Upgrade-Insecure-Requests", "1")
	// req.Header.Set("Sec-Fetch-Dest", "document")
	// req.Header.Set("Sec-Fetch-Mode", "navigate")
	// req.Header.Set("Sec-Fetch-Site", "none")
	// req.Header.Set("Sec-Fetch-User", "?1")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	defer resp.Body.Close()

	respStatus := resp.StatusCode

	fmt.Println(resp.StatusCode)
	if (respStatus >= 200 && respStatus < 300) || (respStatus >= 400 && respStatus < 500) {
		return true, nil
	} else {
		return false, fmt.Errorf("status code not validate and equal %d", respStatus)
	}
}
