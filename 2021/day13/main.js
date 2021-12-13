/*
Run with `deno run --allow-read main.js`
*/

const contents = await Deno.readTextFile("input.txt");
const sections = contents.split("\n\n");

let dots = [];
let lines_dots = sections[0].split("\n");
for (let i = 0; i < lines_dots.length; i++) {
	dots.push(
		lines_dots[i].split(",").map(Number)
	);
}

let folds = [];
let lines_folds = sections[1].split("\n");
lines_folds.pop();
for (let i = 0; i < lines_folds.length; i++) {
	let words = lines_folds[i].split(" ");
	let instruction = words[words.length-1].split("=");

	let coords = [0, 0];
	if (instruction[0] == "y") {
		coords[1] = Number(instruction[1]);
	} else if (instruction[0] == "x") {
		coords[0] = Number(instruction[1]);
	}
	folds.push(coords);
}

let smallest = [...dots[0]];
let biggest = [...smallest];
for (let i = 0; i < dots.length; i++) {
	if (dots[i][0] < smallest[0]) {
		smallest[0] = dots[i][0];
	} else if (dots[i][0] > biggest[0]) {
		biggest[0] = dots[i][0];
	}

	if (dots[i][1] < smallest[1]) {
		smallest[1] = dots[i][1];
	} else if (dots[i][1] > biggest[1]) {
		biggest[1] = dots[i][1];
	}
}

let paper = [];
for (let i = 0; i <= biggest[1] - smallest[1]; i++) {
	let line = Array(biggest[0] - smallest[0] + 1);
	line.fill(0);
	paper.push(line);
}

for (let i = 0; i < dots.length; i++) {
	paper[dots[i][1]][dots[i][0]]++;
}

function fold(paper, coords) {
	let folded = JSON.parse(JSON.stringify(paper));
	if (coords[1] > 0) {
		folded = folded.slice(0, coords[1]);
	} else if (coords[0] > 0) {
		for (let y = 0; y < folded.length; y++) {
			folded[y] = folded[y].slice(0, coords[0]);
		}
	}

	for (let y = coords[1] != 0 ? coords[1]+1 : 0; y < paper.length; y++) {
		for (let x = coords[0] != 0 ? coords[0]+1 : 0; x < paper[y].length; x++) {
			if (paper[y][x] > 0) {
				let mirrored = [
					Math.abs(coords[1] - (y - coords[1])),
					Math.abs(coords[0] - (x - coords[0]))
				];

				folded[mirrored[0]][mirrored[1]] += paper[y][x];
			}
		}
	}

	return folded;
}

function part1(paper, folds) {
	let folded = fold(paper, folds[0]);
	let count = 0;
	for (let y = 0; y < folded.length; y++) {
		for (let x = 0; x < folded[y].length; x++) {
			if (folded[y][x] > 0) {
				count++;
			}
		}
	}
	return count;
}

function part2(paper, folds) {
	let folded = paper;
	for (let i = 0; i < folds.length; i++) {
		folded = fold(folded, folds[i]);
	}

	for (let y = 0; y < folded.length; y++) {
		let line = "";
		for (let x = 0; x < folded[y].length; x++) {
			line += folded[y][x] > 0 ? "#" : " ";
		}
		console.log(line);
	}
}

console.log(part1(paper, folds));
part2(paper, folds);
