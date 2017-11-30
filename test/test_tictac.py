"""Test boards."""
from unittest import TestCase

from otabot.tictac import Board, X, O, E


class BoardTest(TestCase):

    def setUp(self):
        self.g = Board()

    def test_get(self):
        a = 0, 0
        self.assertEqual(
            self.g.get(a),
            E,
        )
        self.g.set(a, O)
        self.assertEqual(
            self.g.get(a),
            O,
        )
        b = 2, 1
        self.g.set(b, X)
        self.assertEqual(
            self.g.get(b),
            X,
        )

    def test_set(self):
        self.g.set(0, 0, X)
        self.assertEqual(
            self.g.get(0, 0),
            X,
        )
        self.g.unset(0, 0)
        self.assertEqual(
            self.g.get(0, 0),
            E,
        )
