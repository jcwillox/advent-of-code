import { describe } from "vitest";
import * as day from "./day{{ cookiecutter.day }}";
import { runBench } from "./utils/tests";

describe("part1", () => {
  runBench(1, 1, day.part1);
});

describe("part2", () => {
  runBench(1, 2, day.part2);
});
