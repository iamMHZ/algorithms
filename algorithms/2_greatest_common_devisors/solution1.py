def naive_gcd(a: int, b: int) -> int:
    gcd = 0

    for i in range(1, a + b):
        if a % i == 0 and b % i == 0:
            gcd = i

    return gcd
