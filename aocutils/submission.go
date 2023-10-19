package aocutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"time"

	"aoc.mb/secret"
)

func getSubmissionFileName(year string, day string, level int) string {
	now := time.Now().Unix()
	return fmt.Sprintf("aoc%s/%s/submission-%d-%d.txt", year, day, level, now)
}

func waitAfterLatestSubmission(year string, day string, level int, waitTime int) bool {
	var filenames []string

	fmt.Println("Checking for latest submission")
	files, err := os.ReadDir(fmt.Sprintf("aoc%s/%s/", year, day))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(("Checking submission names"))
	for _, file := range files {
		if strings.Contains(file.Name(), fmt.Sprintf("submission-%d", level)) {
			filenames = append(filenames, file.Name())
		}
	}

	if len(filenames) == 0 {
		fmt.Println("No matching submission files found")
		return false
	}

	sort.Strings(filenames)
	latestSubmissionName := filenames[len(filenames)-1]
	tmp := strings.Split(latestSubmissionName, "-")
	latestSubmissionTime, _ := strconv.Atoi(strings.Split(tmp[len(tmp)-1], ".")[0])

	timeDelta := int(time.Now().Unix()) - latestSubmissionTime
	fmt.Printf("Waiting for %d seconds\n", waitTime-timeDelta)
	if (timeDelta) < waitTime {
		fmt.Printf("Last submission was less than %d seconds ago\n", waitTime)
		return true
	}
	return false

}

func SubmitAocResult(year string, day string, level int, answer int) {
	waitTime := 300
	if waitAfterLatestSubmission(year, day, level, waitTime) {
		return
	}

	fileName := getSubmissionFileName(year, day, level)
	fmt.Printf("Save submission to file %s\n", fileName)
	_answer := []byte(fmt.Sprint(answer))
	WriteDataToFile(fileName, _answer)

	fmt.Printf("Submitting answer for year %s, day %s, level %d to\n%s\n\n", year, day, level, getAocDailyUrl(year, day))
	cmd := exec.Command(
		"python3",
		"aocutils/post.py",
		fmt.Sprintf("--answer=%d", answer),
		fmt.Sprintf("--level=%d", level),
		fmt.Sprintf("--session=%s", secret.Session),
		fmt.Sprintf("--year=%s", year),
		fmt.Sprintf("--day=%s", day),
	)

	var outbuf, errbuf strings.Builder
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	stdout := outbuf.String()
	stderr := errbuf.String()

	if strings.Contains(stdout, "seem to be solving the right level") {
		fmt.Println("You don't seem to be solving the right level")
	} else if strings.Contains(stdout, "You gave an answer too recently") {
		fmt.Println("You gave an answer too recently")
	} else if strings.Contains(stdout, "not the right answer") {
		fmt.Println("That's not the right answer")
		if strings.Contains(stdout, "your answer is too high") {
			fmt.Println("Your answer is too high")
		} else if strings.Contains(stdout, "your answer is too low") {
			fmt.Println("Your answer is too low")
		}
	} else if strings.Contains(stdout, "the right answer") {
		fmt.Println("That's the right answer")
	} else {
		fmt.Println("Unknown response:")
		fmt.Printf("Stdout: %s\n", stdout)
		fmt.Printf("Stderr: %s\n", stderr)
	}
}

// In theory, this should work
// but it always seems to get redirected.
// Hence, the implemtation in python that is
// invoked here
//
//lint:ignore U1000 Ignore unused function, keep it as an example for a solution attempt
func submitAocResultInGo(year string, day string, level int, answer int) {
	client := &http.Client{}
	url := getAocDailyUrl(year, day) + "answer"

	/* defining payload bytes string*/
	bodyMap := map[string]string{"level": fmt.Sprint(level), "answer": fmt.Sprint(answer)}
	bodyBytes, _ := json.Marshal(bodyMap)

	req, err := http.NewRequest("GET", url, bytes.NewReader(bodyBytes))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Cookie", fmt.Sprintf("session=%s", secret.Session))
	req.Header.Add("User-Agent", "https://github.com/mbecker12/aoc by marvinbecker@mail.de")
	fmt.Println(req.Method)
	fmt.Println(req.URL)
	fmt.Println(req.Header)

	resp, err := client.Do(req)
	handleResponse(resp, err)
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	a := string(body)
	fmt.Println(a)
}
