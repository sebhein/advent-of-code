def solve():
    hits = 0
    index = 50
    with open("input.txt") as fin:
        for line in fin.readlines():
            turn = int(line[1:])
            assert turn < 1000

            while turn > 0:
                if turn < 100:
                    nt = turn
                    turn = 0
                else:
                    nt = 100
                    turn -= 100
                    hits += 1
                    continue

                if line[0] == "L":
                    next = index - nt
                    if next < 0:
                        if index > 0:
                            hits += 1
                        index = 100 + next
                    else:
                        index = next
                    print("L", index, turn, nt, hits)
                else:
                    next = index + nt
                    if next > 99:
                        index = next - 100
                        if index > 0:
                            hits += 1
                    else:
                        index = next
                    print("R", index, turn, nt, hits)

            assert index >= 0
            assert index < 100

            if index == 0:
                hits += 1

    print(f"Hit 0 {hits} times")

if __name__ == "__main__":
    solve()
