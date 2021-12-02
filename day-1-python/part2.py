f = open("input.txt", "r")
w = []
total = 0
prev = 0
count = 0
for x in f:
    w.append(int(x))
    if sum(w) > prev and prev != 0:
        total += 1

    print(w)

    prev = sum(w)
    count += 1

    if len(w) == 3:
        w.pop(0)

# if there is `n` lines in the file there is `n - 2` sliding windows
print(total - 2)
