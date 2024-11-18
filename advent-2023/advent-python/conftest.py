from glob import glob
from types import ModuleType
from typing import cast
from pytest import Metafunc
from os.path import exists


def load_sample(path: str):
    if exists(path):
        with open(path) as file:
            return file.read().strip()
    return ""


def load_cases(day: str, part: str):
    cases: list[tuple[str, str]] = []
    outputs = map(load_sample, glob(f"../samples/{day}/p{part}.*.output.txt"))

    for idx, output in enumerate(outputs):
        input_sample = load_sample(f"../samples/{day}/p{part}.s{idx+1}.input.txt")
        if not input_sample:
            input_sample = load_sample(f"../samples/{day}/s{idx+1}.input.txt")
        cases.append((input_sample, output))

    return cases


def pytest_generate_tests(metafunc: Metafunc):
    cases = load_cases(
        day=cast(ModuleType, metafunc.module).__spec__.name.removesuffix("_test"),
        part=metafunc.definition.name.removeprefix("test_part"),
    )
    ids = map(lambda x: f"sample{x+1}", range(len(cases)))
    metafunc.parametrize("source,output", cases, ids=ids)
