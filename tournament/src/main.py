import sys
from typing import List


def main():
    print("Hello from tournament!")


def start_tournament(input: list[int]) -> List[int]:
    it = iter(input)

    pairs = zip(it, it)

    result: List[int] = []
    for pair in pairs:
        if pair[0] < pair[1]:
            result.append(pair[0])
        else:
            result.append(pair[1])
    return result

def check_if_even(number_of_players: int) -> None:
    if number_of_players % 2 != 0:
            sys.exit("Not even number of players")



def draw_generator(number_of_players: int) -> List[int]:
    check_if_even(number_of_players)

    def helper(players: List[int]) -> List[int]:
        if len(players) == 1:
            return players
        half = len(players) // 2
        left = helper(players[:half])
        right = helper(players[half:])
        result: List[int] = []
        for l, r in zip(left, right):
            result.append(l)
            result.append(r)
        return result

    players = [i for i in range(1, number_of_players + 1)]
    return helper(players)


if __name__ == "__main__":
    input: List[int] = draw_generator(8)

    start_tournament(input)

    print(draw_generator(8))

