import sys
import itertools

# Function taken from: https://github.com/alan-turing-institute/advent-of-code-2024/blob/main/day-07/python_radka-j/day07.py
def can_be_true(vals, tot, operators):
    for comb in itertools.product(operators, repeat=len(vals) - 1):
        res = vals[0]
        for i, n in enumerate(vals[1:]):
            if comb[i] == "+":
                res += n
            elif comb[i] == "*":
                res *= n
            elif comb[i] == "||":
                res = int(str(res) + str(n))
        if res == tot:
            return True
    return False

def main():
    try:
        file_content = sys.stdin.read()
    except Exception as e:
        print(f"Error reading input: {e}")
        return

    lines = file_content.splitlines()

    part1_lines = []
    part2_lines = []

    for line in lines:
        line = line.strip()

        if ":" in line:
            left, right = line.split(":", 1)
            left_int = int(left)
            right_numbers_int = list(map(int, right.strip().split()))

            # Check for part 1: we try combinations of "+" and "*"
            if can_be_true(right_numbers_int, left_int, operators=["+", "*"]):
                part1_lines.append(left_int)
            # Check for part 2: try combinations of "+", "*", and "||"
            elif can_be_true(right_numbers_int, left_int, operators=["+", "*", "||"]):
                part2_lines.append(left_int)

    print("part 1: ", sum(part1_lines))
    print("part 2: ", sum(part1_lines) + sum(part2_lines))

if __name__ == "__main__":
    main()
