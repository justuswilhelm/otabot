"""Benchmark MinMax."""
from unittest import TestCase
import cProfile
import pstats


from otabot.tictac import Board, X, O
from otabot.minmax import MinMax


class MinMaxTest(TestCase):

    def setUp(self):
        self.g = Board()
        self.s = MinMax()

    def test_empty(self):
        def fn():
            return self.s.make_move(self.g, O, True, O)
        self.g.set((0, 0), X)
        expected = (1, 1), 0
        cProfile.runctx(
            "fn()", globals(), locals(), "Profile.prof"
        )

        s = pstats.Stats("Profile.prof")
        s.strip_dirs().sort_stats("time").print_stats()
        self.assertEqual(
            fn(),
            expected,
        )
