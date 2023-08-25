const readline = require("readline");
const rl = readline.createInterface({
    input: process.stdin
});
let numbers;

rl.on("line", line => {
    maxBrackets = Number(line);
    function bracketsGenerator(maxBrackets, open, close, currBrackets) {
		if (open < maxBrackets) {
    		bracketsGenerator(maxBrackets, open + 1, close, currBrackets + "(");
    	}
    	if (close < open) {
    		bracketsGenerator(maxBrackets, open, close + 1, currBrackets + ")");
    	}
    	if (maxBrackets === close) {
    		process.stdout.write(currBrackets + "\n");
   		}
	}
	bracketsGenerator(maxBrackets, 1, 0, "(");
});