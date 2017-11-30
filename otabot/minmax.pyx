# cython: linetrace=True
# distutils: define_macros=CYTHON_TRACE_NOGIL=1
"""Minmax implementation."""
from operator import itemgetter


from .base cimport Game


cdef class MinMax:
    def make_move(self, Game game, int current_player, bint maxim, int color):
        def _iter():
            opposite = not color
            not_maxim = not maxim
            for position in game.iter_moves():
                game.set(position, color)
                yield (
                    position,
                    self.make_move(
                        game,
                        current_player,
                        not_maxim,
                        opposite,
                    )[1],
                )
                game.unset(position)
        if game.over():
            if game.win(current_player):
                return None, 1
            if game.win(not current_player):
                return None, -1
            return None, 0
        if maxim:
            return max(_iter(), key=itemgetter(1))
        return min(_iter(), key=itemgetter(1))
