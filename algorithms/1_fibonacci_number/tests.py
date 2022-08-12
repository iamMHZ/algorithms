from unittest import TestCase
from solution1 import recursive_fibonacci


class RecursiveFibonacciTest(TestCase):

    def test_with_zero(self):
        result = recursive_fibonacci(0)
        self.assertEquals(result, 0)

    def test_with_one(self):
        result = recursive_fibonacci(1)
        self.assertEquals(result, 1)

    def test_with_negative_number(self):
        self.fail()

    def test_with_positive_number(self):
        result = recursive_fibonacci(3)
        self.assertEquals(result, 2)
