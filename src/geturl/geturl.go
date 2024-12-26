package geturl

import (
	"bufio"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"os"
	"regexp"
	"strings"
)

const (
	httpStr  string = "http://"
	matchStr string = "^https://*|^http://*"
)

func GetUrl() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	if err != nil {
		return "", err
	}
	_, err = validateUrl(text)
	if err != nil {
		return "", err
	}
	return text, nil
}

func validateUrl(str string) (bool, error) {
	if len(str) < 10 {
		return false, fmt.Errorf("link less 10 chars")
	}
	res, err := regexp.MatchString(matchStr, str)
	fmt.Print("\n", res, "\n")
	if err != nil {
		fmt.Println(err)
		return false, fmt.Errorf("error in regex: %v", err)
	}
	if res {
		if succes, err := getResponse(str); succes {
			fmt.Println(str)
			return true, nil
		} else {
			fmt.Println(err)
			return false, fmt.Errorf("response code not supported\nCode:%v", err)
		}
	} else {
		str = httpStr + str
		fmt.Println(str)
		if succes, err := getResponse(str); succes {
			fmt.Println(str)
			return true, nil
		} else {
			fmt.Println(err)
			return false, fmt.Errorf("response code not supported\nCode:%v", err)
		}
	}
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
