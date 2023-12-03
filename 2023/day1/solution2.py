import re
import functools

f = open("input", "r")

mydict = {"one": "1",
          "two": "2",
          "three": "3",
          "four": "4",
          "five": "5",
          "six": "6",
          "seven": "7",
          "eight": "8",
          "nine": "9"
          }

regex = "one|two|three|four|five|six|seven|eight|nine|\d"

total = 0

for line in f:
    chars = []
    while len(line):
        m = re.search(regex, line)
        if m:
            chars.append(m.group())
            line = line[m.start() + 1:]
        else:
            break

    x = mydict.get(chars[0], chars[0])
    y = mydict.get(chars[-1], chars[-1])
    num = "".join([x, y])
    total += int(num)

print(total)
