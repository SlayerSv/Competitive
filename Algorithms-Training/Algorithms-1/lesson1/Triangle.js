const fs = require('fs')
let fileContent = fs.readFileSync("input.txt", "utf8");
let [sideA, sideB, sideC] = fileContent.toString().split('\n');
sideA = +sideA;
sideB = +sideB;
sideC = +sideC;
let result;

if ((sideA >= sideB + sideC) || (sideB >= sideA + sideC) || (sideC >= sideB + sideA)) {
	result = "NO";
} else if (sideA <= 0 || sideB <= 0 || sideC <= 0) {
	result = "NO";
} else {
	result = "YES";
}
fs.writeFileSync("output.txt", result.toString())