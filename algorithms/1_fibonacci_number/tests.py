from unittest import TestCase
from solution1 import recursive_fibonacci
from solution2 import fibonacci_dynamic_programming


class RecursiveFibonacciTest(TestCase):

    def test_with_zero(self):
        result = recursive_fibonacci(0)
        self.assertEquals(result, 0)

    def test_with_one(self):
        result = recursive_fibonacci(1)
        self.assertEquals(result, 1)

    def test_with_positive_number(self):
        result = recursive_fibonacci(3)
        self.assertEquals(result, 2)

    def test_for_large_number(self):
        result = recursive_fibonacci(20)
        self.assertEquals(result, 6765)

    def test_with_negative_number(self):
        self.fail()


class DynamicProgrammingFibonacciTest(TestCase):
    def test_with_zero(self):
        result = fibonacci_dynamic_programming(0)
        self.assertEquals(result, 0)

    def test_with_one(self):
        result = fibonacci_dynamic_programming(1)
        self.assertEquals(result, 1)

    def test_with_positive_number(self):
        result = fibonacci_dynamic_programming(3)
        self.assertEquals(result, 2)

    def test_for_large_number(self):
        result = fibonacci_dynamic_programming(20)
        self.assertEquals(result, 6765)
