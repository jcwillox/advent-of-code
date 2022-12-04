# Advent of Code ![christmas-tree](https://api.iconify.design/twemoji/christmas-tree.svg?height=28) [![cookiecutter](https://api.iconify.design/icomoon-free/pacman.svg?color=%23d4aa00&height=26)](https://github.com/cookiecutter/cookiecutter)

This repository is a cookiecutter template but it also contains my solutions which are not included when use this template. This is thanks to the way cookiecutter works, in which the top-level directory is not included in the template.

## Usage

First you need to install cookiecutter, see the [installation](https://cookiecutter.readthedocs.io/en/latest/installation.html) docs for more info.

The easiest way is usually `pip` or `pipx`:
```bash
# using pip
pip install cookiecutter
# using pipx
pipx install cookiecutter
```

**Using the template**

Running the following command will create a folder named `advent-<year>` with a folder for each language you choose.

```bash
cookiecutter gh:jcwillox/advent-of-code
```

**Adding more days and languages**

The command above only creates 1 year, 1 language, 1 day. Naturally you might want more than that. All you need to do is run the same command again with `-f -s` in the same directory and it will add the additional files without overwriting your changes.

```bash
cookiecutter gh:jcwillox/advent-of-code -f -s
```
