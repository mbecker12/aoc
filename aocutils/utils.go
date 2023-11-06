package aocutils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"aoc.mb/secret"
)

var BaseAocUrl = "https://adventofcode.com/"

func getAocDailyUrl(year string, day string) string {
	return fmt.Sprintf(BaseAocUrl+"%s/day/%s/", year, day)
}

func getDataFileName(year string, day string) string {
	return fmt.Sprintf("aoc%s/%s/data.txt", year, day)
}

func handleResponse(resp *http.Response, err error) {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalln(resp)
	}
	if err != nil {
		log.Fatalln(err)
	}
}

func DownloadAocInput(year string, day string) []byte {
	fmt.Printf("Fetching input for year %s, day %s\n", year, day)
	fileName := getDataFileName(year, day)
	if fileExists(fileName) {
		fmt.Printf("Reading data from file %s\n", fileName)
		return readDataFromFile(fileName)
	}
	fmt.Println("Downloading data")
	client := &http.Client{}
	url := getAocDailyUrl(year, day) + "input"
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Cookie", fmt.Sprintf("session=%s", secret.Session))
	req.Header.Set("User-Agent", "https://github.com/mbecker12/aoc by marvinbecker@mail.de")

	resp, err := client.Do(req)
	handleResponse(resp, err)
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Input fetched successfully")
	fmt.Println("")
	return body
}

func fileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	if err == nil {
		return true // File exists
	} else if os.IsNotExist(err) {
		return false // File does not exist
	} else {
		fmt.Printf("Error checking file: %v\n", err)
		return false // Error occurred
	}
}

func WriteDaysDataToFile(year string, day string, data []byte) {
	fileName := getDataFileName(year, day)
	WriteDataToFile(fileName, data)
}

func WriteDataToFile(fileName string, data []byte) {
	// Open the file for writing. If the file doesn't exist, it will be created.
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}

func readDataFromFile(fileName string) []byte {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil
	}
	return data
}

func SplitByteInput(input []byte, delimiter string) []string {
	inputStr := string(input)
	return SplitStringInput(inputStr, delimiter)
}

func SplitStringInput(input string, delimiter string) []string {
	groups := strings.Split(input, delimiter)

	return groups
}
