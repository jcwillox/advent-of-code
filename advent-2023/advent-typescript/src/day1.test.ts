import { describe } from "vitest";
import * as day from "./day1";
import { runTests } from "./utils/tests";

describe("part1", () => {
  runTests(1, 1, day.part1);
});

describe("part2", () => {
  runTests(1, 2, day.part2);
});
