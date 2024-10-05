words = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


def part1(source: str):
    total = 0
    for line in source.splitlines(keepends=False):
        numbers = [x for x in line if x.isnumeric()]
        total += int(numbers[0] + numbers[-1])
    return total


def part2(source: str):
    total = 0
    for line in source.splitlines(keepends=False):
        # first number
        num1 = ""
        for i, c in enumerate(line):
            if c.isnumeric():
                num1 = c
                break
            for word in words:
                if line[i:].startswith(word):
                    num1 += words[word]
                    break
            if len(num1) == 1:
                break
        # last number
        num2 = ""
        for i, c in reversed(list(enumerate(line))):
            if c.isnumeric():
                num2 = c
                break
            for word in words:
                if line[: i + 1].endswith(word):
                    num2 += words[word]
                    break
            if len(num2) == 1:
                break
        total += int(num1 + num2)
    return total
