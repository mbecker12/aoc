import requests
import argparse

def parse_args():
    parser = argparse.ArgumentParser(
        description="""Post a result to adventofcode.com"""
    )
    parser.add_argument(
        "--answer",
        type=int,
        help="Answer to the riddle",
    )
    parser.add_argument(
        "--level",
        type=int,
        help="Determines level 1 or 2 of the daily challenge",
        default=2
    )
    parser.add_argument(
        "--session",
        type=str,
        help="User-related session cookie",
    )
    parser.add_argument(
        "--url",
        type=str,
        help="Advent of code base url",
        default="https://adventofcode.com/"
    )
    parser.add_argument(
        "--year",
        type=str,
        help="Advent of code year",
    )
    parser.add_argument(
        "--day",
        type=str,
        help="Advent of code day",
    )

    return parser.parse_args()

args = parse_args()
url = f"{args.url}{args.year}/day/{args.day}/answer"
data = {"level": args.level, "answer": args.answer}
headers = {"Cookie": f"session={args.session}", "User-Agent": "https://github.com/mbecker12/aoc by marvinbecker@mail.de"}
resp = requests.post(url, data=data, headers=headers)

resp.raise_for_status()

print(f"{resp.content=}")
print(f"{resp.status_code=}")
