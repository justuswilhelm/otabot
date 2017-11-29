#!/usr/bin/env python3
"""."""
from operator import itemgetter
import copy
import logging


logging.basicConfig(level=logging.DEBUG)


X = False
O = True
E = 2

player = X
pc = O


colors = {
    X: 'X',
    E: 'E',
    O: 'O',
}


class Board:

    def __init__(self):
        self.board = [
            E, E, E,
            E, E, E,
            E, E, E,
        ]

    def column(self, color):
        b = self.board
        for i in range(3):
            if b[i + 0] == b[i + 3] == b[i + 6] == color:
                return True
        return False

    def row(self, color):
        b = self.board
        for i in range(3):
            if b[i * 3 + 0] == b[i * 3 + 1] == b[i * 3 + 2] == color:
                return True
        return False

    def diagonal(self, color):
        b = self.board
        return (
            b[0] == b[4] == b[8] == color or
            b[2] == b[4] == b[7] == color
        )

    def win(self, color):
        return self.diagonal(color) or self.row(color) or self.column(color)

    def over(self):
        return self.win(X) or self.win(O) or len(tuple(self.iter_moves())) == 0

    def get(self, row, column):
        i = row * 3 + column
        return self.board[i]

    def set(self, row, column, value):
        if value != E:
            assert self.get(row, column) == E
        i = row * 3 + column
        self.board[i] = value

    def unset(self, row, column):
        assert self.get(row, column) != E
        self.set(row, column, E)

    def copy(self):
        board = self.__class__()
        board.board = copy.deepcopy(self.board)
        return board

    def iter_moves(self):
        for col in range(3):
            for row in range(3):
                if self.get(row, col) == E:
                    yield row, col

    def __str__(self):
        return "{}{}{}\n{}{}{}\n{}{}{}".format(
            *(
                colors[s] for s in
                self.board
            )
        )


def mcts(board, current_player, maxim, color, depth=0):
    def _iter():
        opposite = not color
        not_maxim = not maxim
        for row, column in board.iter_moves():
            board.set(row, column, color)
            yield (
                (row, column),
                mcts(
                    board,
                    current_player,
                    not_maxim,
                    opposite,
                    depth + 1,
                )[1],
            )
            board.unset(row, column)
    if board.over():
        if board.win(current_player):
            return None, 1
        if board.win(not current_player):
            return None, -1
        return None, 0
    if maxim:
        return max(_iter(), key=itemgetter(1))
    return min(_iter(), key=itemgetter(1))


def main():
    """."""
    print("You are {}".format(colors[player]))
    b = Board()
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
                b.set(row, column, player)
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
        move, score = mcts(b, pc, True, pc)
        if move is None:
            print("Draw!")
            return
        (p_row, p_col) = move
        print("Computer move: {} {} ({} score)".format(p_row, p_col, score))
        print()
        b.set(p_row, p_col, pc)
        print(b)
        print()
        if b.win(pc):
            print("You lose!")
            return


if __name__ == "__main__":
    main()
