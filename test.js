function solve(a, b, c) {
  let num1;
  let num2;
  let same;

  num1 = a / b;
  num2 = a / c;
  let least = lcm(b, c);
  same = a / least;
  num1 -= same;
  num2 -= same;
  let ans;
  ap1 = (0n + (num1 - 1n)) * num1 / 2n;
  ap2 = (2n + (num2 - 1n)) * num2 / 2n;
  ans = a * num1 - ap1 - ap2;
  return ans;
}
  
const readline = require("readline");
const rl = readline.createInterface({
  input: process.stdin
});
  
let lineNumber = -1;
let meta = -2;
let lastData = -2;
let data = [];

  
rl.on("line", line => {
  if (lineNumber === -1) {

  } else if (lineNumber === meta) {
      if (lastData < 0) {
        meta = lineNumber + 2;
      } else {
        lastData = lineNumber + +line;
        meta = lastData + 1;
      }
      
  } else {
    if (lastData < 0) {
      const [a, b, c] = line.split(" ").map(BigInt);
      const answer = solve(a, b, c);
      console.log(answer.toString());
    } else {
      data.push(line.split(" ").map(Number));
      if (lineNumber === lastData) {
        const answer = solve(data);
        console.log(answer.toString());
        data = [];
      }
    }
  }
  lineNumber++;
});

function gcd(a, b)
{
if (b == 0)
    return a;
return gcd(b, a % b);
}
 
// Function to return LCM of two numbers
function lcm(a, b)
{
    return a / gcd(a, b) * b;
}