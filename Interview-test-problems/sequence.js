const readline = require("readline");

const rl = readline.createInterface({
	input: process.stdin
});

let counter = 0;
let maxSequence = 0;
let start = 0;

rl.on("line", line => {
	if (line === "1" && start) {
    	counter++;
      maxSequence = Math.max(maxSequence, counter);
    } else {
    	counter = 0;
    }
    start++;
}).on("close", () => {
	process.stdout.write(maxSequence.toString());
});