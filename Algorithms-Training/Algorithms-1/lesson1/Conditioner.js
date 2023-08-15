const fs = require('fs')
let fileContent = fs.readFileSync("input.txt", "utf8");
const [temps, type] = fileContent.toString().split('\n');
let [tRoom, tCond] = temps.toString().split(' ');
let result;
if (type === "heat") {
	result = tCond > tRoom ? tCond : tRoom;
} else if (type === "freeze") {
	result = tCond < tRoom ? tCond : tRoom;
} else if (type === "auto") {
	result = tCond;
} else if (type === "fan") {
	result = tRoom;  
}
fs.writeFileSync("output.txt", result.toString())