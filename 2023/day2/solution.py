f = open('input', 'r')

# Bag loaded with..
RED = 12
GREEN = 13
BLUE = 14

def parse_line(line):
    line = line.rstrip() # remove newline
    line = line.split(":")

    id = line[0].split(" ")[-1]

    games = line[1].split(";")

    turns = []
    for game in games:
        g = game.split(",")
        for turn in g:
            t = turn.lstrip().split(' ')
            turns.append(t)

    return id, turns

total = 0

for line in f:
    p = parse_line(line)

    valid = True
    for t in p[1]:
        if t[1] == 'green' and int(t[0]) > GREEN:
            valid = False
            break
        elif t[1] == 'blue' and int(t[0]) > BLUE:
            valid = False
            break
        elif t[1] == 'red' and int(t[0]) > RED:
            valid = False
            break

    if valid:
        total += int(p[0])


print(total)

