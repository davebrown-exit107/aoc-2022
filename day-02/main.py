import os

def score_game(input_lines):
    return 0


def new_score_game(input_lines):
    return 0

if __name__ == '__main__':
    with open('input.txt') as in_fh:
        file_lines = [line.strip() for line in in_fh.readlines()]

    score = score_game(file_lines)
    print(f"Score: {score}")
    new_score = new_score_game(file_lines)
    print(f"New Score: {new_score}")
