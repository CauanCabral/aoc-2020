file = open("input", "r", encoding="utf-8")

## Rule: 'symbol' should be present on 'value' only on posX or posY, not both
def isValid(value, symbol, posX, posY):
    return (value[posX - 1] == symbol or value[posY - 1] == symbol) and value[posX - 1] != value[posY - 1]

valids = 0
for line in file:
    tokens = line.split()
    [posX, posY] = tokens[0].split('-')
    symbol = tokens[1][:-1]
    password = tokens[2]

    if isValid(value=password, symbol=symbol, posX=int(posX), posY=int(posY)):
        valids += 1

print("There are ", valids, " valids password")