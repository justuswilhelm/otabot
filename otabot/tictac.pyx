# cython: linetrace=True
"""TicTacToe."""
from operator import itemgetter
import copy
from .base cimport Game


cdef public enum Colors:
    O = 0
    X = 1
    E = 2


colors = {
    X: 'X',
    E: 'E',
    O: 'O',
}


cdef class Board(Game):
    cdef int[9] board

    def __cinit__(self):
        self.board = [
            E, E, E,
            E, E, E,
            E, E, E,
        ]

    cpdef void set_board(self, board):
        self.board = board

    cpdef bint column(self, int color):
        b = self.board
        return (
            b[0] == b[3] == b[6] == color or
            b[1] == b[4] == b[7] == color or
            b[2] == b[5] == b[8] == color
        )

    cpdef bint row(self, int color):
        b = self.board
        return (
            b[0] == b[1] == b[2] == color or
            b[3] == b[4] == b[5] == color or
            b[6] == b[7] == b[8] == color
        )

    cpdef bint diagonal(self, int color):
        b = self.board
        return (
            b[0] == b[4] == b[8] == color or
            b[2] == b[4] == b[7] == color
        )

    cpdef bint win(self, int color):
        return self.diagonal(color) or self.row(color) or self.column(color)

    cpdef bint over(self):
        return any([
            self.win(X),
            self.win(O),
            len(tuple(self.iter_moves())) == 0
        ])

    cpdef int get(self, tuple position):
        i = position[0] * 3 + position[1]
        return self.board[i]

    cpdef void set(self, tuple position, int value):
        if value != E:
            assert self.get(position) == E
        i = position[0] * 3 + position[1]
        self.board[i] = value

    cpdef void unset(self, tuple position):
        assert self.get(position) != E
        self.set(position, E)

    def iter_moves(self):
        for col in range(3):
            for row in range(3):
                if self.get((row, col)) == E:
                    yield row, col

    def __str__(self):
        return "{}{}{}\n{}{}{}\n{}{}{}".format(
            *(
                colors[s] for s in
                self.board
            )
        )
