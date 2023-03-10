import os

def score_game(input_lines):
    outcome_matrix = [[3,6,0],[0,3,6],[6,0,3]]
    move_matrix = [1,2,3]
    total_score = 0
    for game in input_lines:
        them, us = game.split(" ")
        if them == "A":
            if us == 'X':
                total_score += outcome_matrix[0][0] + move_matrix[0]
            if us == 'Y':
                total_score += outcome_matrix[0][1] + move_matrix[1]
            if us == 'Z':
                total_score += outcome_matrix[0][2] + move_matrix[2]
        if them == "B":
            if us == 'X':
                total_score += outcome_matrix[1][0] + move_matrix[0]
            if us == 'Y':
                total_score += outcome_matrix[1][1] + move_matrix[1]
            if us == 'Z':
                total_score += outcome_matrix[1][2] + move_matrix[2]
        if them == "C":
            if us == 'X':
                total_score += outcome_matrix[2][0] + move_matrix[0]
            if us == 'Y':
                total_score += outcome_matrix[2][1] + move_matrix[1]
            if us == 'Z':
                total_score += outcome_matrix[2][2] + move_matrix[2]
    return total_score


def new_score_game(input_lines):
    move_matrix = [[3,1,2],[1,2,3],[2,3,1]]
    outcome_matrix = [0,3,6]
    total_score = 0
    for game in input_lines:
        them, us = game.split(" ")
        if them == "A":
            if us == 'X':
                total_score += move_matrix[0][0] + outcome_matrix[0]
            if us == 'Y':
                total_score += move_matrix[0][1] + outcome_matrix[1]
            if us == 'Z':
                total_score += move_matrix[0][2] + outcome_matrix[2]
        if them == "B":
            if us == 'X':
                total_score += move_matrix[1][0] + outcome_matrix[0]
            if us == 'Y':
                total_score += move_matrix[1][1] + outcome_matrix[1]
            if us == 'Z':
                total_score += move_matrix[1][2] + outcome_matrix[2]
        if them == "C":
            if us == 'X':
                total_score += move_matrix[2][0] + outcome_matrix[0]
            if us == 'Y':
                total_score += move_matrix[2][1] + outcome_matrix[1]
            if us == 'Z':
                total_score += move_matrix[2][2] + outcome_matrix[2]
    return total_score

if __name__ == '__main__':
    with open('input.txt') as in_fh:
        file_lines = [line.strip() for line in in_fh.readlines()]

    score = score_game(file_lines)
    print(f"Score: {score}")
    new_score = new_score_game(file_lines)
    print(f"New Score: {new_score}")
