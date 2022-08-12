def fibonacci_dynamic_programming(up_to: int) -> int:
    if up_to <= 1:
        return up_to

    first_element = 0
    second_element = 1

    for i in range(1, up_to):
        temp = first_element + second_element

        first_element = second_element
        second_element = temp

    return second_element
