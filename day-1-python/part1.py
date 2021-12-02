f = open("input.txt", "r")
total = 0
prev = ""
for i, curr in enumerate(f):
    if i == 0:
        continue
    if curr > prev:
        total += 1
    prev = curr

print(total)
