function solve(input) {
  let answer;
  
  return answer;
}
  
const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin
});
  
let lineNumber = 0;
// let meta = 0;
// let data = [];
// let lastData = 1;
  
rl.on("line", line => {
  if (lineNumber === -1) {
  } else if (lineNumber % 2 === 0){

    // lastData = +line;
    // meta = lastData + 1;
  } else {
    const answer = solve(line);
    console.log(answer.toString());
    // data.push(line);
    // if (lineNumber === lastData) {
    //   const answer = solve(data);
    //   console.log(answer.toString());
    //   data = [];
    // }
  }
  lineNumber++;
});