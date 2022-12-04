import unittest

from main import top_calories, top_three_calories

class TestCalories(unittest.TestCase):
    def test_top_calories(self):
        sample_input = [ "1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"]

        want = 24000
        got = top_calories(sample_input)

        self.assertEqual(got, want)

    def test_top_three_calories(self):
        sample_input = [ "1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"]

        want = 24000
        got = top_three_calories(sample_input)

        self.assertEqual(got, want)

if __name__ == '__main__':
    unittest.main()
