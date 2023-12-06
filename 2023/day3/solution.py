import re

f = open('input', 'r')

engine = f.read().splitlines()

total = 0

for idx, line in enumerate(engine):
    matches = re.finditer('\d*', line)

    for m in matches:
        if m.group() != '':
            x, y = m.span()

            # check line above
            if idx > 0:
                xx = x - 1
                if xx < 0:
                    xx = 0
                yy = y + 1
                if yy > len(line):
                    yy = len(line)

                # find adjacent symbols
                s = re.search('[^\d\.]', engine[idx - 1][xx:yy])

                if s != None:
                    total += int(m.group())
                    # continue

            # check same line
            xx = x - 1
            if xx >= 0:
                s = re.search('[^\d\.]', line[xx])
                if s != None:
                    total += int(m.group())
                    # continue

            yy = y
            if yy < len(line):
                s = re.search('[^\d\.]', line[yy])
                if s != None:
                    total += int(m.group())
                    # continue

            # check line below
            if idx + 1 < len(engine):
                xx = x - 1
                if xx < 0:
                    xx = 0
                yy = y + 1
                if yy > len(line):
                    yy = len(line)

                # find adjacent symbols
                s = re.search('[^\d\.]', engine[idx + 1][xx:yy])

                if s != None:
                    total += int(m.group())
                    # continue


print(total)
