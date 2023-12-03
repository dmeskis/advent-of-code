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

    greens = []
    reds = []
    blues = []
    for t in p[1]:
        if t[1] == 'green':
            greens.append(int(t[0]))
        elif t[1] == 'blue':
            blues.append(int(t[0]))
        elif t[1] == 'red':
            reds.append(int(t[0]))

    g = max(greens)
    r = max(reds)
    b = max(blues)

    total += g * r * b

print(total)

