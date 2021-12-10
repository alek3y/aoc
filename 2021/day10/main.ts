/*
Run with `deno run --allow-read main.ts`
*/

const contents = await Deno.readTextFile("input.txt");
const lines = contents.split("\n");
lines.pop();

const MATCHING: any = {
	"(": [")", 3, 1],
	"[": ["]", 57, 2],
	"{": ["}", 1197, 3],
	"<": [">", 25137, 4]
}

function parser(line: string): number[] {
	let openings = Object.keys(MATCHING);

	let stack = [];
	for (let i = 0; i < line.length; i++) {
		let character = line.charAt(i);

		if (openings.includes(character)) {
			stack.push(i);
		} else {
			let opening = stack.pop();
			if (opening && MATCHING[line.charAt(opening)][0] !== character) {
				throw new Error(character);
			}
		}
	}

	return stack;
}

function part1(lines: string[]): number {
	let values = Object.values(MATCHING);
	let closings = values.map((value: any) => (value[0]));
	let scores = values.map((value: any) => (value[1]));

	let illegal_count = Array(closings.length).fill(0);
	for (let i = 0; i < lines.length; i++) {
		let line = lines[i];

		try {
			parser(line)
		} catch (error) {
			illegal_count[closings.indexOf(error.message)] += 1;
		}
	}

	let score = 0;
	for (let i = 0; i < closings.length; i++) {
		score += illegal_count[i] * scores[i];
	}

	return score;
}

function part2(lines: string[]): number {
	let values = Object.values(MATCHING);
	let openings = Object.keys(MATCHING);
	let scores = values.map((value: any) => (value[2]));

	let lines_score = []
	for (let i = 0; i < lines.length; i++) {
		let line = lines[i];

		let score = 0;
		try {
			let stack = parser(line);
			let opening;
			while ((opening = stack.pop()) !== undefined) {
				score = score*5 + scores[openings.indexOf(line[opening])];
			}
		} catch (error) {}

		if (score != 0) {
			lines_score.push(score);
		}
	}

	let sorted = lines_score.sort((n1, n2) => n1 - n2);
	return sorted[Math.trunc(sorted.length/2)];
}

console.log(part1(lines));
console.log(part2(lines));
