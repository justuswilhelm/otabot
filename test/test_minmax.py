"""Test MinMax."""
from unittest import TestCase

from otabot.tictac import Board, X, O, E
from otabot.minmax import MinMax


class MinMaxTest(TestCase):

    def setUp(self):
        self.g = Board()
        self.s = MinMax()

    def test_draw(self):
        self.g.set_board([
            O, X, E,
            O, X, O,
            X, O, X,
        ])
        expected = (0, 2), 0
        self.assertEqual(
            self.s.make_move(self.g, X, True, X),
            expected,
        )
        expected = (0, 2), 0
        self.assertEqual(
            self.s.make_move(self.g, X, True, O),
            expected,
        )
        expected = (0, 2), 0
        self.assertEqual(
            self.s.make_move(self.g, O, True, O),
            expected,
        )
        expected = (0, 2), 0
        self.assertEqual(
            self.s.make_move(self.g, O, True, X),
            expected,
        )

    def test_last(self):
        self.g.set_board([
            O, O, E,
            O, X, O,
            X, O, X,
        ])
        expected = (0, 2), 0
        self.assertEqual(
            self.s.make_move(self.g, X, True, X),
            expected,
        )
        expected = (0, 2), -1
        self.assertEqual(
            self.s.make_move(self.g, X, True, O),
            expected
        )

    def test_two_empty(self):
        self.g.set_board([
            X, O, O,
            X, X, O,
            E, O, E,
        ])
        expected = (2, 0), 1.0
        self.assertEqual(
            self.s.make_move(self.g, X, True, O),
            expected
        )
        expected = (2, 0), -1.0
        self.assertEqual(
            self.s.make_move(self.g, O, True, X),
            expected
        )
        expected = (2, 2), 1.0
        self.assertEqual(
            self.s.make_move(self.g, O, True, O),
            expected
        )

    def test_done(self):
        self.g.set_board([
            X, O, O,
            X, X, O,
            E, O, O,
        ])
        expected = None, 1.0
        self.assertEqual(
            self.s.make_move(self.g, O, True, O),
            expected
        )
        expected = None, -1
        self.assertEqual(
            self.s.make_move(self.g, X, True, X),
            expected
        )
