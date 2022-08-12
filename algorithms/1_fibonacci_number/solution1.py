def recursive_fibonacci(up_to: int) -> int:
    if up_to <= 1:
        return up_to

    return recursive_fibonacci(up_to - 1) + recursive_fibonacci(up_to - 2)
