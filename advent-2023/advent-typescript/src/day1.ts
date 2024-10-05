export function part1(source: string) {
  return (
    source
      .split("\n")
      // filter out digits
      .map((x) => [...x].filter((c) => c >= "0" && c <= "9"))
      // make line number
      .map((x) => Number(x[0] + x.slice(-1)))
      // sum
      .reduce((a, b) => a + b, 0)
  );
}

export function part2(source: string) {
  const numberWords = {
    one: "1",
    two: "2",
    three: "3",
    four: "4",
    five: "5",
    six: "6",
    seven: "7",
    eight: "8",
    nine: "9",
  };
  return (
    source
      .split("\n")
      // filter out digits
      .map((line) => {
        // find first number word
        let minIdx = line.length;
        let minNum = "";
        for (const [word, num] of Object.entries(numberWords)) {
          const idx = line.indexOf(word);
          if (idx > -1 && idx < minIdx) {
            minIdx = idx;
            minNum = num;
          }
        }
        // find first number
        for (const num of Object.values(numberWords)) {
          const idx = line.indexOf(num);
          if (idx > -1 && idx < minIdx) {
            minIdx = idx;
            minNum = num;
          }
        }
        // find last number word
        let maxIdx = -1;
        let maxNum = "";
        for (const [word, num] of Object.entries(numberWords)) {
          let idx = line.lastIndexOf(word);
          if (idx > -1) {
            idx += word.length - 1;
            // console.log("MaxNum1", idx, word)
            if (idx > maxIdx) {
              maxIdx = idx;
              maxNum = num;
            }
          }
        }
        // find last number
        for (const num of Object.values(numberWords)) {
          const idx = line.lastIndexOf(num);
          // console.log("MaxNum2", idx, num)
          if (idx > -1 && idx > maxIdx) {
            maxIdx = idx;
            maxNum = num;
          }
        }

        return Number(minNum + maxNum);
      })
      // sum
      .reduce((a, b) => a + b, 0)
  );
}
