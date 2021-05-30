import re


class WordTree(object):
    def __init__(self):
        self.value = None
        self.weight = 0
        self.length = 0
        self.leftNode = None
        self.rightNode = None

    def add_value(self, val):
        if self.value is None:  # add first element
            self.value = val
            self.weight = 1
            return

        if self.value < val:
            if self.rightNode is not None:
                self.rightNode.__add_value(val)
            else:
                newRight = WordTree()
                newRight.value = val
                newRight.weight = 1
                self.rightNode = newRight
            return

        if self.value > val:
            if self.leftNode is not None:
                self.leftNode.__add_value(val)
            else:
                newLeft = WordTree()
                newLeft.value = val
                newLeft.weight = 1
                self.leftNode = newLeft
            return

        self.weight += 1
        return

    def is_leaf(self):
        return self.leftNode is None and self.rightNode is None

    def __three_max_node(self):
        if self.value is None:
            return []

        if self.is_leaf():
            return [self]

        leftMax = []
        rightMax = []

        if self.leftNode:
            leftMax = self.leftNode.__three_max_node()
        if self.rightNode:
            rightMax = self.rightNode.__three_max_node()

        maxList = leftMax + rightMax
        maxList.append(self)

        return sorted(maxList, reverse=True, key=lambda x: x.weight)[0:3]

    def three_max(self):
        return [e.value for e in self.__three_max_node()]

    def print(self):  # debug
        if self.is_leaf():
            return f"<{self.value}, {self.weight}>"

        if self.leftNode is not None and self.rightNode is not None:
            return f"<{self.value}, {self.weight}> - ({self.leftNode.print()}) - ({self.rightNode.print()})"
        elif self.leftNode is not None:
            return f"<{self.value}, {self.weight}> - ({self.leftNode.print()}) - None"

        return f"<{self.value}, {self.weight}> - None - ({self.rightNode.print()})"


def top_3_words(text):
    tree = WordTree()

    for val in text.split(" "):
        found = re.findall(r"([a-zA-Z]+('[a-zA-Z]+)?)", val)
        if len(found) > 0:
            tree.add_value(found[0][0].lower())

    return tree.three_max()
