def solve_p1():
    sum = 0
    with open("input.txt") as fin:
        for line in fin.readlines():
            nums = [int(s) for s in line.strip()]
            first = max(nums[:len(nums) - 1])
            index = nums.index(first)
            second = max(nums[index + 1:])
            # print(first, second)
            sum += first * 10 + second

    print(f"Sum of joltages: {sum}")

def solve_p2():
    joltage_len = 12
    sum = 0
    with open("input.txt") as fin:
        for line in fin.readlines():
            nums = [int(s) for s in line.strip()]
            start_idx = 0
            for iter in range(1, joltage_len + 1):
                lo, hi = start_idx, len(nums) - joltage_len + iter
                num = max(nums[lo:hi])
                start_idx = lo + nums[lo:hi].index(num) + 1
                sum += num * (10**(joltage_len - iter))

    print(f"Sum of joltages: {sum}")

if __name__ == "__main__":
    solve_p1()
    solve_p2()
