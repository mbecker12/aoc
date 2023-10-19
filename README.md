# aoc

Collection of Advent of Code solutions


## Automation Guidelines

This script/repo/tool does follow the automation guidelines on the /r/adventofcode community wiki https://www.reddit.com/r/adventofcode/wiki/faqs/automation.

Specifically:

- Outbound calls are throttled to every 5 minutes in `aocutils/submission.go`
- Once inputs are downloaded, they are cached locally (`DownloadAocInput()`)
  - If you suspect your input is corrupted, you can manually request a fresh copy using manualDownloadFunction() - To be implemented
- The User-Agent header in `aocutils/post.py` is set to me since I maintain this tool :)
