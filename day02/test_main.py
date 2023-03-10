import unittest

from main import score_game, new_score_game

class TestScoring(unittest.TestCase):
    def test_score_game(self):
        sample_input = ["A Y", "B X", "C Z"]

        want = 15
        got = score_game(sample_input)

        self.assertEqual(got, want)

    def test_new_score_game(self):
        sample_input = ["A Y", "B X", "C Z"]

        want = 12
        got = new_score_game(sample_input)

        self.assertEqual(got, want)

if __name__ == '__main__':
    unittest.main()
