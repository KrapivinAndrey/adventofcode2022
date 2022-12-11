from tqdm import tqdm


class Monkey:
    def __init__(
        self, items: str, operation: str, divider: int, index_true, index_false
    ):
        self.items = [int(x.strip()) for x in items.split(",")]
        self.divider = divider
        self.true_monkey = index_true
        self.false_monkey = index_false
        self.operation = operation
        self.inspected = 0

    def play(self):
        throw = []
        for old in self.items:
            result = eval(self.operation) % 9699690
            if result % self.divider == 0:
                throw.append((self.true_monkey, result))
            else:
                throw.append((self.false_monkey, result))
            self.inspected += 1
        self.items.clear()

        return throw

    def __repr__(self):
        return " ".join([str(x) for x in self.items])


def main():

    monkeys = []
    monkey0 = Monkey("85, 79, 63, 72", "old * 17", 2, 2, 6)
    monkey1 = Monkey("53, 94, 65, 81, 93, 73, 57, 92", "old * old", 7, 0, 2)
    monkey2 = Monkey("62, 63", "old + 7", 13, 7, 6)
    monkey3 = Monkey("57, 92, 56", "old + 4", 5, 4, 5)
    monkey4 = Monkey("67", "old + 5", 3, 1, 5)
    monkey5 = Monkey("85, 56, 66, 72, 57, 99", "old + 6", 19, 1, 0)
    monkey6 = Monkey("86, 65, 98, 97, 69", "old * 13", 11, 3, 7)
    monkey7 = Monkey("87, 68, 92, 66, 91, 50, 68", "old + 2", 17, 4, 3)

    monkeys.append(monkey0)
    monkeys.append(monkey1)
    monkeys.append(monkey2)
    monkeys.append(monkey3)
    monkeys.append(monkey4)
    monkeys.append(monkey5)
    monkeys.append(monkey6)
    monkeys.append(monkey7)

    for i in tqdm(range(10000)):
        for monkey in monkeys:
            throw = monkey.play()
            for num, res in throw:
                monkeys[num].items.append(res)

    for monkey in monkeys:
        print(monkey.inspected)


if __name__ == "__main__":
    main()
