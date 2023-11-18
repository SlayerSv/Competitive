function solve() {
  let ans;
  

  return ans;
}


let i = 0;

function processData() {
  const answer = [];
  let result;
  for (i; i < data.length; i = i + 1) {
    result = solve(...formTestData(1));
    if (Array.isArray(result)) result = result.join(" ")
    answer.push(result);
  }
  process.stdout.write(answer.join("\n"));
  process.exit(0);
}

function formTestData(linesCount) {
  if (!linesCount) linesCount = toNums(data[i++]);
  const testData = [];
  while (linesCount--) {
    testData.push(toNums(data[i++]));
  }
  return testData;
}

function toNums(line) {
  line = line.split(" ").map(Number);
  return line.length === 1 ? line[0] : line;
}

const data = [];
const rl = require("readline").createInterface({input: process.stdin});
rl.on("line", line => {data.push(line)});
rl.on("close", processData);
process.on("SIGINT", () => rl.close());