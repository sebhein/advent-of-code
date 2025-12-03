def solve_p1():
    with open("input.txt") as fin:
        line = fin.read()

    sum = 0

    for rnge in line.split(","):
        start, end = rnge.strip().split("-")
        start_viable = len(start) % 2 == 0
        end_viable = len(end) % 2 == 0
        if not start_viable and not end_viable and (len(start) == len(end)):
            continue

        for id in range(int(start), int(end) + 1):
            str_id = str(id)
            if len(str_id) % 2 != 0:
                continue

            mid = len(str_id) // 2

            if str_id[:mid] == str_id[mid:]:
                sum += id

    print(f"Sum of invalid ids: {sum}")

def solve_p2():
    with open("input.txt") as fin:
        line = fin.read()

    sum = 0

    for rnge in line.split(","):
        start, end = rnge.strip().split("-")
        for id in range(int(start), int(end) + 1):
            str_id = str(id)

            for maybe_stride in reversed(range(1, (len(str_id) // 2) + 1)):
                if len(str_id) % maybe_stride != 0:
                    continue

                first, second = 0, maybe_stride
                invalid = True
                while second <= len(str_id) - maybe_stride:
                    if str_id[first: second] != str_id[second: second + maybe_stride]:
                        invalid = False
                        break
                    first += maybe_stride
                    second += maybe_stride

                if invalid:
                    sum += id
                    break


    print(f"Sum of invalid ids: {sum}")

if __name__ == "__main__":
    solve_p1()
    solve_p2()
