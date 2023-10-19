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

const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin
});

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
      data = [];
      lineNumber = 0;
    }
  }
  lineNumber++;
});

function printAnswer(answer) {
  if (Array.isArray(answer)) {
    console.log(answer.join(" "));
  }
  else {
    console.log(answer.toString());
  }
}