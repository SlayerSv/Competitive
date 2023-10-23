function solve() {
  let ans;
  

  return ans;
}

function prepareData(data) {
  data = data.map(line => {
    line = line.split(" ").map(Number);
    if (line.length === 1) return line[0];
    return line;
  });


  return data;
}

let lineNumber = 0;
const SKIP = [0];
let lastData = 1;
let meta = -2;
let data = [];

const rl = require("readline").createInterface({input: process.stdin});
rl.on("line", line => {
  if (SKIP.includes(lineNumber)) {
  } else if (lineNumber === meta) {
      lastData = lineNumber + Number(line);
  } else {
    data.push(line);
    if (lineNumber === lastData) {
      const preparedData = prepareData(data);
      const answer = solve(...preparedData);
      printAnswer(answer);
      data.length = 0;
      lineNumber = 0;
    }
  }
  ++lineNumber;
});

function printAnswer(answer) {
  if (Array.isArray(answer)) {
    answer[answer.length - 1] += "\n";
    process.stdout.write(answer.join(" "));
  }
  else {
    process.stdout.write(answer + "\n");
  }
}