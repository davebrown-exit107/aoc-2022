import os

def top_calories(input_lines):
    return 0

def top_three_calories(input_lines):
    return 0

if __name__ == '__main__':
    with open('input.txt') as in_fh:
        file_lines = [lines.strip() for line in in_fh.readlines()]

    top_cals = top_calories(file_lines)
    top_three_cals = top_three_calories(file_lines)
    print(f"Top calories held: {top_cals}")
    print(f"Top three calories held: {top_cals}")
