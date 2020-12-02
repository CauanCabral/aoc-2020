file = open("input", "r", encoding="utf-8")

## Rule: 'symbol' should be present on 'value' n-times, where min <= n <= max
def isValid(value, symbol, min, max):
    count = value.count(symbol)

    return count >= min and count <= max

valids = 0
for line in file:
    tokens = line.split()
    [countMin, countMax] = tokens[0].split('-')
    symbol = tokens[1][:-1]
    password = tokens[2]

    if isValid(value=password, symbol=symbol, min=int(countMin), max=int(countMax)):
        valids += 1

print("There are ", valids, " valids password")