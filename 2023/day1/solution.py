import re
import functools

f = open("input", "r")

total = 0

for line in f:
    chars = re.findall("\d", line)
    num = "".join([chars[0], chars[-1]])
    total += int(num)

print(total)
