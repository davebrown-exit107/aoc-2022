import os

def top_calories(input_lines):
    top_calories = 0
    cur_calories = 0
    for line in input_lines:
        if line == "":
            if cur_calories > top_calories:
                top_calories = cur_calories
            cur_calories = 0
        else:
            cur_calories += int(line)
    if cur_calories > top_calories:
        top_calories = cur_calories
    return top_calories


def top_three_calories(input_lines):
    top_calories = []
    cur_calories = 0
    for line in input_lines:
        if line == "":
            top_calories.append(cur_calories)
            cur_calories = 0
        else:
            cur_calories += int(line)
    top_calories.append(cur_calories)
    top_calories.sort(reverse=True)
    return sum(top_calories[0:3])

if __name__ == '__main__':
    with open('input.txt') as in_fh:
        file_lines = [line.strip() for line in in_fh.readlines()]

    top_cals = top_calories(file_lines)
    top_three_cals = top_three_calories(file_lines)
    print(f"Top calories held: {top_cals}")
    print(f"Top three calories held: {top_three_cals}")
