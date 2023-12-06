import re

# answer: 85010461
# 81085275

# f = open('example', 'r')
f = open('input', 'r')

engine = f.read().splitlines()

total = 0

for idx, line in enumerate(engine):
    # Find gear
    matches = re.finditer('\*', line)

    for m in matches:
        if m.group() != '':
            x, y = m.span()

            parts = []
            # check line above
            if idx > 0:
                xx = x - 1
                if xx < 0:
                    xx = 0
                yy = y + 1
                if yy > len(line):
                    yy = len(line)

                # find adjacent numbers
                s = re.search('\d', engine[idx - 1][xx:yy])

                if s != None:
                    xx = xx + s.start()
                    yy = xx + 1


                    while xx > 0:
                        if not engine[idx - 1][xx - 1].isdigit():
                            break
                        xx -= 1

                    while yy < len(engine[idx - 1]):
                        if not engine[idx - 1][yy].isdigit():
                            break

                        yy += 1

                    parts.append(int(engine[idx-1][xx:yy]))

            # check same line
            xx = x - 1
            if xx >= 0:
                s = re.search('\d', line[xx])
                print(s)
                if s != None:
                    while xx > 0:
                        if not line[xx - 1].isdigit():
                            break
                        xx -= 1
                    print(line[xx:y - 1])
                    parts.append(int(line[xx:y - 1]))

            yy = y
            if yy < len(line):
                s = re.search('\d', line[yy])
                print(s)
                if s != None:
                    while yy < len(line):
                        if not line[yy].isdigit():
                            break
                        yy += 1
                    print(line[x + 1:yy])
                    parts.append(int(line[x + 1:yy]))

            # check line below
            if idx + 1 < len(engine):
                xx = x - 1
                if xx < 0:
                    xx = 0
                yy = y + 1
                if yy > len(line):
                    yy = len(line)

                # find adjacent symbols
                s = re.search('\d', engine[idx + 1][xx:yy])

                if s != None:
                    xx = xx + s.start()
                    yy = xx + 1
                    while xx > 0:
                        if not engine[idx + 1][xx - 1].isdigit():
                            break
                        xx -= 1

                    while yy < len(engine[idx + 1]):
                        if not engine[idx + 1][yy].isdigit():
                            break
                        yy += 1

                    parts.append(int(engine[idx+1][xx:yy]))

            if len(parts) == 2:
                total += parts[0] * parts[1]


print(total)
