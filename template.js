function solve() {
  let ans;
  

  return ans;
}

function prepareData(data) {
  data = [...data[0].split(" ").map(Number)];
  

  return data;
}

let lineNumber = 0;
let meta = -2;
let lastData = -2;
let data = [];

const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin
});

rl.on("line", line => {
  if (lineNumber === -1) {

  } else if (lineNumber === meta) {
      if (lastData < 0) {
        meta = lineNumber + 2;
      } else {
        lastData = lineNumber + Number(line);
        meta = lastData + 1;
      }
      
  } else {
    data.push(line);
    if (lastData < 0 || lineNumber === lastData) {
      const preparedData = prepareData(data);
      const answer = solve(...preparedData);
      printAnswer(answer);
      data = [];
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