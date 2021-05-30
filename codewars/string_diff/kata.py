import string
from functools import cmp_to_key


def filter_lower(e):
    """
    Filter only lower letters.
    """
    return e.lower() == e and e in string.ascii_letters


def count_char(s):
    """
    Count occurrences of letters in a sorted list of characters.

    :param s: sorted list of character
    :return: list of tuple representing the character and its frequency in s ([a, 4], [d, 3]...)
    """
    current = None
    count = 0
    res = []
    for c in s:
        if not current:
            current = c
        if c == current:
            count += 1
        else:
            res.append([current, count])
            current = c
            count = 1
    res.append([current, count])
    return res


def combine_counted_list(count1, count2, ind1, ind2):
    """
    Take two counted list from count_char function and combine them in a single list
    keeping the greatest occurrence of each char. If char appears more in count1 ind1 will
    be used as reference, if char appears more in count2 ind2 will be used as reference else
    = will be used as reference

    :param count1: counted list of char
    :param count2: counted list of char
    :param ind1: reference name for count1
    :param ind2: reference name for count2
    :return: A list of combined value indicating the origin of the greatest occurrence, the letter concerned and its frequency
    """
    if not count1 and not count2:
        return []

    if not count2:
        return [[ind1, count1[0][0], count1[0][1]]] + combine_counted_list(count1[1:], count2, ind1, ind2)

    if not count1:
        return [[ind2, count2[0][0], count2[0][1]]] + combine_counted_list(count1, count2[1:], ind1, ind2)

    if count1[0][0] < count2[0][0]:
        return [[ind1, count1[0][0], count1[0][1]]] + combine_counted_list(count1[1:], count2, ind1, ind2)

    if count1[0][0] > count2[0][0]:
        return [[ind2, count2[0][0], count2[0][1]]] + combine_counted_list(count1, count2[1:], ind1, ind2)

    if count1[0][1] > count2[0][1]:
        return [[ind1, count1[0][0], count1[0][1]]] + combine_counted_list(count1[1:], count2[1:], ind1, ind2)

    if count1[0][1] < count2[0][1]:
        return [[ind2, count2[0][0], count2[0][1]]] + combine_counted_list(count1[1:], count2[1:], ind1, ind2)

    return [["=", count1[0][0], count1[0][1]]] + combine_counted_list(count1[1:], count2[1:], ind1, ind2)


def sort_combined_counted(a, b):
    """
    Sort a combined list of occurrence from combine_counted_list
    Orders priority are:
    - a frequency is greater than b
    - a reference is lower than b
    - a char is lower than b

    :return: -1 if a is lower than b using orders priority, 1 else (required for cmp_to_key function)
    """
    if a[2] != b[2]:
        if a[2] > b[2]:
            return -1
        return 1
    if a[0] != b[0]:
        if a[0] < b[0]:
            return -1
        return 1
    if a[1] < b[1]:
        return -1
    return 1


def mix(s1, s2):
    # your code
    s1List = sorted(filter(filter_lower, s1))
    s2List = sorted(filter(filter_lower, s2))

    s1Count = count_char(s1List)
    s2Count = count_char(s2List)
    combined = combine_counted_list(s1Count, s2Count, "1", "2")

    return '/'.join([f"{e[0]}:{e[2] * e[1]}" for e in sorted(filter(lambda x: x[2] > 1, combined), key=cmp_to_key(sort_combined_counted))])
