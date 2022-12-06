from glob import glob
from types import ModuleType
from typing import cast

from pytest import Metafunc


def load_cases(day: str, part: str):
    def load_file(path: str):
        with open(path) as file:
            return file.read().strip()

    return list(
        zip(
            map(load_file, glob(f"../samples/{day}/*.input.txt")),
            map(load_file, glob(f"../samples/{day}/*.{part}.output.txt")),
        )
    )


def pytest_generate_tests(metafunc: Metafunc):
    cases = load_cases(
        day=cast(ModuleType, metafunc.module).__spec__.name.removesuffix("_test"),
        part=metafunc.definition.name.removeprefix("test_part"),
    )
    ids = map(lambda x: f"sample{x+1}", range(len(cases)))
    metafunc.parametrize("source,output", cases, ids=ids)
