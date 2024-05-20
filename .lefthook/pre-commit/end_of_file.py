#!/usr/bin/env python3
# Adapted from https://github.com/pre-commit/pre-commit-hooks/blob/main/pre_commit_hooks/end_of_file_fixer.py (MIT)
from __future__ import annotations

import os
from isbinary import is_binary_file
from typing import IO
from typing import Sequence


def fix_file(file_obj: IO[bytes]) -> int:
    # Test for newline at end of file
    # Empty files will throw IOError here
    try:
        file_obj.seek(-1, os.SEEK_END)
    except OSError:
        return 0

    last_character = file_obj.read(1)

    # last_character will be '' for an empty file
    if last_character not in {b'\n', b'\r'} and last_character != b'':
        # Needs this seek for windows, otherwise IOError
        file_obj.seek(0, os.SEEK_END)
        file_obj.write(b'\n')
        return 1

    while last_character in {b'\n', b'\r'}:
        # Deal with the beginning of the file
        if file_obj.tell() == 1:
            # If we've reached the beginning of the file and it is all
            # linebreaks then we can make this file empty
            file_obj.seek(0)
            file_obj.truncate()
            return 1

        # Go back two bytes and read a character
        file_obj.seek(-2, os.SEEK_CUR)
        last_character = file_obj.read(1)

    # Our current position is at the end of the file just before any amount of
    # newlines.  If we find extraneous newlines, then backtrack and trim them.
    position = file_obj.tell()
    remaining = file_obj.read()

    for sequence in (b'\n', b'\r\n', b'\r'):
        if remaining == sequence:
            return 0
        elif remaining.startswith(sequence):
            file_obj.seek(position + len(sequence))
            file_obj.truncate()
            return 1

    return 0


def main() -> int:
    exclude = {'./.git'}
    retv = 0

    for root, dirs, files in os.walk('.'):
        dirs[:] = [d for d in dirs if os.path.join(root, d) not in exclude]
        for filename in files:
            if is_binary_file(os.path.join(root, filename)):
                continue

            # Read as binary so we can read byte-by-byte
            filename = os.path.join(root, filename)
            with open(filename, 'rb+') as file_obj:
                ret_for_file = fix_file(file_obj)

                if ret_for_file:
                    print(f'Fixing {filename}')

                retv |= ret_for_file

    return retv


if __name__ == '__main__':
    main()
