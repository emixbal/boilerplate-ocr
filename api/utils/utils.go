package utils

import (
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/otiai10/gosseract"
)

/*
 * @desc Make multiline string to array string
 *
 * @param  {string} s - string with multiple linne
 *
 * @return {array of string, error}
 */
func BreakMultiLine(s string) ([]string, error) {
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return nil, err
	}

	s = regex.ReplaceAllString(s, "\n")

	regexs, err := regexp.Compile("\n")
	if err != nil {
		return nil, err
	}
	s = regexs.ReplaceAllString(s, ", ")

	arr := strings.Split(s, ", ")

	tmp := arr[len(arr)-1]
	tmp = tmp[:len(tmp)-6]

	arr[len(arr)-1] = tmp
	return arr, nil
}

/*
 * @desc Save image from url to local
 *
 * @param  {string} url - url image
 *
 * @return {string, error}
 */
func SaveImage(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", nil
	}
	defer response.Body.Close()

	path := "./tmp/coba.jpg"
	file, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}
	return path, nil
}

/*
 * @desc Extract image to string
 *
 * @param  {string} path - path image to extract
 *
 * @return {string
 */
func Extract(path string) string {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetImage(path)
	txt, _ := client.Text()
	return txt
}
