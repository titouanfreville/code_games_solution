import unittest

from word_frequency.frequency import top_3_words


class Test(unittest.TestCase):
    @staticmethod
    def assert_equals(v1, v2):
        assert v1 == v2


Test.assert_equals(top_3_words("a a a  b  c c  d d d d  e e e e e"), ["e", "d", "a"])
Test.assert_equals(top_3_words("e e e e DDD ddd DdD: ddd ddd aa aA Aa, bb cc cC e e e"), ["e", "ddd", "aa"])
Test.assert_equals(top_3_words("  //wont won't won't "), ["won't", "wont"])
Test.assert_equals(top_3_words("  , e   .. "), ["e"])
Test.assert_equals(top_3_words("  ...  "), [])
Test.assert_equals(top_3_words("  '  "), [])
Test.assert_equals(top_3_words("  '''  "), [])
Test.assert_equals(top_3_words("""In a village of La Mancha, the name of which I have no desire to call to
mind, there lived not long since one of those gentlemen that keep a lance
in the lance-rack, an old buckler, a lean hack, and a greyhound for
coursing. An olla of rather more beef than mutton, a salad on most
nights, scraps on Saturdays, lentils on Fridays, and a pigeon or so extra
on Sundays, made away with three-quarters of his income."""), ["a", "of", "on"])
