function solve(input) {
  let answer;
  
  return answer;
}
  
const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin
});
  
let lineNumber = -1;
  
rl.on("line", line => {
  if (lineNumber === -1) {
    // skip
  } else if (lineNumber % 2 === 0){
    
  } else {
    const answer = solve(line);
    console.log(answer.toString());
  }
  lineNumber++;
});