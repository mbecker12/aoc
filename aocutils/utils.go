package aocutils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	// "aoc.mb/secret"
)

var PythonBinPath = "/home/marvin/virtualenv/aoc/bin/python"
var BaseAocUrl = "https://adventofcode.com/"

func GetBasePath() string {
	cmdOut, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		fmt.Printf(`Error on getting the go-kit base path: %s - %s\n`, err.Error(), string(cmdOut))
		os.Exit(-1)
		return ""
	}
	path := string(cmdOut)
	path = strings.TrimSpace(path) + "/"
	return path
}

func getAocDailyUrl(year string, day string) string {
	return fmt.Sprintf(BaseAocUrl+"%s/day/%s/", year, day)
}

func GetDataFileName(year string, day string) string {
	path := "aoc%s/%s/data.txt"
	return fmt.Sprintf(path, year, day)
}

func handleResponse(resp *http.Response, err error) {
	if resp == nil || err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Fatalln(resp)
	}
}

func DownloadAocInput(year string, day string) []byte {
	fmt.Printf("Fetching input for year %s, day %s\n", year, day)
	fileName := GetBasePath() + GetDataFileName(year, day)
	if FileExists(fileName) {
		fmt.Printf("Reading data from file %s\n", fileName)
		return ReadDataFromFile(fileName)
	}
	fmt.Println("Downloading data")
	client := &http.Client{}
	url := getAocDailyUrl(year, day) + "input"
	req, _ := http.NewRequest("GET", url, nil)
	sessionToken := GetSessionToken("")

	req.Header.Set("Cookie", fmt.Sprintf("session=%s", sessionToken))
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

func GetSessionToken(filename string) string {
	tokenFile := ""
	if filename == "" {
		tokenFile = ".config/aoc/token"
		homeDir, _ := os.UserHomeDir()
		tokenFile = homeDir + "/" + tokenFile
	}
	sessionToken := ReadDataFromFile(tokenFile)
	sessionTokenString := strings.TrimSpace(string(sessionToken))
	return sessionTokenString
}

func FileExists(fileName string) bool {
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
	fileName := GetDataFileName(year, day)
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

func ReadDataFromFile(fileName string) []byte {
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
