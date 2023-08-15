const readline = require("readline");

const rl = readline.createInterface({
	input: process.stdin
});

let start = 0;
let prev = "";

rl.on("line", line => {
	if (start) {
    	if (prev !== line) {
       		process.stdout.write(line + "\n");
        	prev = line;
    	}
    }
    start++;
});