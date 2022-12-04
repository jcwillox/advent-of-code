#!/usr/bin/env python
import shutil

if __name__ == "__main__":
    languages = ["go", "python", "typescript"]
    selected_language = "{{ cookiecutter.language }}"

    for language in languages:
        if language != selected_language:
            shutil.rmtree(f"advent-{language}", ignore_errors=True)
