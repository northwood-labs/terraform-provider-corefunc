#!/usr/bin/env python3
# Adapted from https://github.com/pre-commit/pre-commit-hooks/blob/main/pre_commit_hooks/trailing_whitespace_fixer.py (MIT)
from __future__ import annotations

import os
from isbinary import is_binary_file
from typing import IO


def fix_file(filename: str) -> bool:
    with open(filename, mode='rb') as file_processed:
        lines = file_processed.readlines()

    newlines = [_process_line(line) for line in lines]

    if newlines != lines:
        with open(filename, mode='wb') as file_processed:
            for line in newlines:
                file_processed.write(line)

        return True
    else:
        return False


def _process_line(line: bytes) -> bytes:
    if line[-2:] == b'\r\n':
        eol = b'\r\n'
        line = line[:-2]
    elif line[-1:] == b'\n':
        eol = b'\n'
        line = line[:-1]
    else:
        eol = b''
    # # preserve trailing two-space for non-blank lines in markdown files
    # if is_markdown and (not line.isspace()) and line.endswith(b'  '):
    #     return line[:-2].rstrip(chars) + b'  ' + eol
    return line.rstrip() + eol


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
            ret_for_file = fix_file(filename)

            if ret_for_file:
                print(f'Fixing {filename}')

            retv |= ret_for_file

    return retv


if __name__ == '__main__':
    main()
