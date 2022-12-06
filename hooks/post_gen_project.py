#!/usr/bin/env python
import shutil
from glob import glob

if __name__ == "__main__":
    remove_suffix = "{{ cookiecutter.__remove_suffix }}"

    for path in glob(f"*{remove_suffix}"):
        shutil.rmtree(path)
