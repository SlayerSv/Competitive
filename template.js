function solve(input) {
  let ans;
  
  return ans;
}
  
const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin
});
  
let lineNumber = -1;
let meta = 0;
let lastData = -2;
let data = [];

  
rl.on("line", line => {
  if (lineNumber === -1) {

  } else if (lineNumber === meta) {
      if (lastData < 0) {
        meta = lineNumber + 2;
      } else {
        lastData = lineNumber + line;
        meta = lastData + 1;
      }
      
  } else {
    if (lastData < 0) {
      const [a, b, c] = line.split(" ").map(Number);
      const answer = solve(a, b, c);
      console.log(answer);
    } else {
      data.push(line);
      if (lineNumber === lastData) {
        const answer = solve(data);
        console.log(answer.toString());
        data = [];
      }
    }
  }
  lineNumber++;
});
