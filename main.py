#!/usr/bin/env python3
import functools
import readline
import random

from decimal import Decimal
from enum import Enum, auto


class RandomPolicy:
    """."""

    def __init__(self, game):
        """."""
        self.game = game

    def genmove(self, color):
        """."""


class FieldState(Enum):
    """."""

    EMPTY = auto()
    BLACK = auto()
    WHITE = auto()


class Game:
    """Store complete go game state."""

    BOARDSIZE = 9
    KOMI = Decimal('7.5')
    COLUMNS = "ABCDEFGHJKLMNOPQ"
    STARS = {
        9: [
            (3, 3),
            (3, 7),
            (7, 3),
            (5, 5),
            (7, 7),
        ]
    }
    POLICY_CLS = RandomPolicy

    def __init__(self, boardsize=None, komi=None):
        """Initialize with board and komi."""
        self.komi: Decimal = komi or self.KOMI
        self.create_board(boardsize or self.BOARDSIZE)
        self.policy = self.POLICY_CLS(self)

    def create_board(self, size):
        """Create a board with a given size."""
        self.board = [
            [FieldState.EMPTY for _ in range(size)]
            for _ in range(size)
        ]

    def clear_board(self):
        """Clear the board."""
        self.create_board(self.boardsize)

    def parse_move(self, move):
        """."""
        column, row = move
        return self.COLUMNS.find(column), int(row) - 1

    def to_move(self, move):
        """."""
        return "A1"

    def play(self, state, move):
        """Play a move."""
        column, row = move
        self.board[row][column] = state

    @property
    def boardsize(self):
        """Return boardsize."""
        return len(self.board)

    def is_star(self, row, column):
        """."""
        stars = self.STARS[self.boardsize]
        return (row, column) in stars

    def genmove(self, color):
        """."""
        self.policy.genmove(color)

    def showboard(self):
        """Show board."""
        def _iter_cols(index, row):
            yield "{:2d}".format(index)
            for col_index, col in enumerate(row, start=1):
                if col == FieldState.EMPTY:
                    if self.is_star(index, col_index):
                        yield '+'
                    else:
                        yield '.'
                elif col == FieldState.WHITE:
                    yield 'O'
                elif col == FieldState.BLACK:
                    yield 'X'
            yield "{}".format(index)

        def _iter_rows():
            cols = "   {}".format(" ".join(self.COLUMNS[:self.boardsize]))
            yield cols
            indices = reversed(range(1, self.boardsize + 1))
            for index, row in zip(indices, self.board):
                yield " ".join(_iter_cols(index, row))
            yield cols
        return "\n{}".format("\n".join(_iter_rows()))


class Shell():
    """Provide shell interface to game state."""

    def __init__(self):
        self.game = Game()
        self.commands = {
            name[3:]: getattr(self, name)
            for name in dir(self) if name.startswith('do_')
        }

    def answer(self, line, success=True):
        prefix_id = self.id or ""
        if success:
            prefix = "={} ".format(prefix_id)
        else:
            prefix = "?{} ".format(prefix_id)
        print("{}{}".format(prefix, line), end="\n\n")

    def call_command(self, name, args):
        """Call a command. Handle errors."""
        try:
            cmd = self.commands[name]
        except KeyError:
            self.answer("Unknown command", success=False)
        else:
            try:
                result, success = cmd(*args)
                result = result or ""
                self.answer(result, success=success)
            except TypeError as e:
                self.answer(
                    "Error when calling function: {}".format(e),
                    success=False,
                )

    def _loop(self):
        while True:
            i = input().strip().split()
            # Format: ""
            if not i:
                self.answer("", success=True)
                self.id = None
            # Format: "id name [args]"
            try:
                id_raw, cmd_name, *args = i
                self.id = int(id_raw)
            # Format: "name [args]"
            except ValueError:
                cmd_name, *args = i
                self.id = None
            self.call_command(cmd_name, args)

    def cmdloop(self):
        readline.parse_and_bind("tab: complete")
        readline.set_completer(self.complete)
        try:
            self._loop()
        except EOFError:
            pass

    @functools.lru_cache()
    def complete_prefix(self, prefix):
        """Find completion for a prefix."""
        return tuple(
            k for k in self.commands.keys()
            if k.startswith(prefix)
        )

    def complete(self, text, state):
        """Return completer for readline."""
        return self.complete_prefix(text)[state]

    def do_version(self):
        """Return the version of this bot."""
        return "1.0.0", True

    def do_name(self):
        """Return the name of this name."""
        return "Otabot", True

    def do_protocol_version(self):
        """Return the GTP version used."""
        return "2", True

    def do_known_command(self, cmd):
        """Return whether a command is known."""
        if cmd in self.commands:
            return "true", True
        else:
            return "false", True

    def do_list_commands(self):
        """List all known commands."""
        result = "\n".join(self.commands.keys())
        return result, True

    def do_boardsize(self, size):
        """Set boardsize."""
        self.game.create_board(int(size))
        return None, True

    def do_clear_board(self):
        """Set boardsize."""
        self.game.clear_board()
        return None, True

    def do_komi(self, komi):
        """Set komi."""
        self.game.komi = Decimal(komi)
        return None, True

    def do_play(self, player, move):
        """Play a move."""
        state = {
            'W': FieldState.WHITE,
            'B': FieldState.BLACK,
        }[player]
        move = self.game.parse_move(move)
        self.game.play(state, move)
        return None, True

    def do_showboard(self):
        return self.game.showboard(), True


if __name__ == '__main__':
    Shell().cmdloop()
