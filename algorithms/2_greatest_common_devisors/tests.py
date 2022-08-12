from unittest import TestCase
from solution1 import naive_gcd


class NaiveGcdTest(TestCase):

    def test_naive_gcd(self):
        result = naive_gcd(12, 6)
        self.assertEquals(result, 6)

    def test_with_both_a_and_b_zero(self):
        self.fail()