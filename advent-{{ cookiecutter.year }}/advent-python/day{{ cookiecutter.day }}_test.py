import day{{ cookiecutter.day }}


def test_part1(source, output):
    assert str(day{{ cookiecutter.day }}.part1(source)) == output


def test_part2(source, output):
    assert str(day{{ cookiecutter.day }}.part2(source)) == output
