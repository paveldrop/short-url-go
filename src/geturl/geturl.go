package geturl

import (
	"bufio"
	"os"
)

func GetUrl() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	} else {
		return text, nil
	}
}

func validateUrl(str string) string {

}
