#!/usr/bin/env python3
"""."""
import logging

from otabot import tictac
from otabot import minmax
from otabot import base


logging.basicConfig(level=logging.DEBUG)


player = tictac.X
pc = tictac.O


def main():
    """."""
    print("You are {}".format(tictac.colors[player]))
    b = tictac.Board()
    s = minmax.MinMax()
    print(b)
    print()
    while True:
        while True:
            try:
                row, column = (int(i) for i in input("Your move: ").split())
            except ValueError:
                print("Invalid input")
                print()
                continue
            try:
                b.set((row, column), player)
            except AssertionError:
                print("Field occupied")
                print()
            else:
                break
        print(b)
        print()
        if b.win(player):
            print("You win!")
            return
        move, score = s.make_move(b, pc, True, pc)
        if move is None:
            print("Draw!")
            return
        print("Computer move: {} ({} score)".format(move, score))
        print()
        b.set(move, pc)
        print(b)
        print()
        if b.win(pc):
            print("You lose!")
            return


if __name__ == "__main__":
    main()
