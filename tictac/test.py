from unittest import TestCase

from game import Board, mcts, X, O, E


class MctsTest(TestCase):

    def test_get(self):
        b = Board()
        self.assertEqual(
            b.get(0, 0),
            E,
        )
        b.set(0, 0, O)
        self.assertEqual(
            b.get(0, 0),
            O,
        )
        b.set(2, 1, X)
        self.assertEqual(
            b.get(2, 1),
            X,
        )

    def test_set(self):
        b = Board()
        b.set(0, 0, X)
        self.assertEqual(
            b.get(0, 0),
            X,
        )
        b.unset(0, 0)
        self.assertEqual(
            b.get(0, 0),
            E,
        )

    def test_draw(self):
        b = Board()
        b.board = [
            O, X, E,
            O, X, O,
            X, O, X,
        ]
        expected = (0, 2), 0
        self.assertEqual(
            mcts(b, X, True, X),
            expected,
        )
        expected = (0, 2), 0
        self.assertEqual(
            mcts(b, X, True, O),
            expected,
        )
        expected = (0, 2), 0
        self.assertEqual(
            mcts(b, O, True, O),
            expected,
        )
        expected = (0, 2), 0
        self.assertEqual(
            mcts(b, O, True, X),
            expected,
        )

    def test_last(self):
        b = Board()
        b.board = [
            O, O, E,
            O, X, O,
            X, O, X,
        ]
        expected = (0, 2), 0
        self.assertEqual(
            mcts(b, X, True, X),
            expected,
        )
        expected = (0, 2), -1
        self.assertEqual(
            mcts(b, X, True, O),
            expected
        )

    def test_two_empty(self):
        b = Board()
        b.board = [
            X, O, O,
            X, X, O,
            E, O, E,
        ]
        expected = (2, 0), 1.0
        self.assertEqual(
            mcts(b, X, True, O),
            expected
        )
        expected = (2, 0), -1.0
        self.assertEqual(
            mcts(b, O, True, X),
            expected
        )
        expected = (2, 2), 1.0
        self.assertEqual(
            mcts(b, O, True, O),
            expected
        )

    def test_done(self):
        b = Board()
        b.board = [
            X, O, O,
            X, X, O,
            E, O, O,
        ]
        expected = None, 1.0
        self.assertEqual(
            mcts(b, O, True, O),
            expected
        )
        expected = None, -1
        self.assertEqual(
            mcts(b, X, True, X),
            expected
        )
