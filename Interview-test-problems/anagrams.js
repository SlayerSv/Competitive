const readline = require("readline");
const rl = readline.createInterface({
	input: process.stdin
});

let words = [];

rl.on("line", line => {
	words.push(line);
}).on("close", () => {
	const [word1, word2] = words;
    if (word1.length !== word2.length) {process.stdout.write("0"); return;}
	const map1 = new Map();
    const map2 = new Map();
    for (let i = 0; i < word1.length; i++) {
    	map1.set(word1[i], map1.get(word1[i]) + 1 || 1);
    }
    for (let i = 0; i < word2.length; i++) {
    	map2.set(word2[i], map2.get(word2[i]) + 1 || 1);
    }
    for (const [letter, count] of map1) {
		if (map2.get(letter) !== count) {process.stdout.write("0"); return;}
    }
    process.stdout.write("1");
});