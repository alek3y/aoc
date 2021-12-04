/*
Run with `deno run --allow-read main.ts`
*/

const contents = await Deno.readTextFile("input.txt");
const lines = contents.split("\n");

const order = lines[0].split(",").map(Number);

class Cell {
	value: number;
	selected: boolean = false;

	constructor(value: number) {
		this.value = value;
	}
}

class Board {
	cells: Cell[] = [];
	width: number = 0;
	last: number = 0;

	constructor(width: number) {
		this.width = width;
	}

	select(value: number) {
		this.cells.forEach((cell) => {
			if (cell.value == value) {
				cell.selected = true;
				this.last = value;
			}
		});
	}

	score(): number {
		let sum = 0;
		this.cells.forEach((cell) => {
			if (!cell.selected) {
				sum += cell.value;
			}
		});
		return sum * this.last;
	}

	winner(): boolean {
		let selected = 0;
		for (let i = 0; i < this.cells.length; i++) {
			selected += this.cells[i].selected ? 1 : 0;

			if (i != 0 && ((i+1) % this.width == 0)) {
				if (selected == this.width) {
					return true;
				}
				selected = 0;
			}
		}

		selected = 0;
		let height = this.cells.length / this.width;
		for (let i = 0; i < this.width; i++) {
			for (let j = 0; j < height; j++) {
				selected += this.cells[this.width*j + i].selected ? 1 : 0;

				if (j == height-1) {
					if (selected == height) {
						return true;
					}
					selected = 0;
				}
			}
		}

		return false;
	}

	clone(): Board {
		let cloned = new Board(this.width);
		this.cells.forEach((cell) => {
			cloned.cells.push(new Cell(cell.value));
		});
		return cloned;
	}
}

let boards: Board[] = [];

let cells: number[] = [];
for (let i = 2; i < lines.length; i++) {
	if (lines[i] === "") {
		boards[boards.length-1].cells = cells.map((value, index, array) => new Cell(value));
		cells = [];
		continue;
	}

	let numbers = lines[i]
		.split(" ")
		.filter((char) => char !== "")
		.map(Number);

	if (cells.length == 0) {
		boards.push(new Board(numbers.length));
	}

	cells = cells.concat(numbers);
}

function part1(order: number[], boards: Board[]): number {
	for (let i = 0; i < order.length; i++) {
		let lucky_number = order[i];
		for (let b = 0; b < boards.length; b++) {
			boards[b].select(lucky_number);

			if (boards[b].winner()) {
				return boards[b].score();
			}
		}
	}

	return 0;
}

function part2(order: number[], boards: Board[]): number {
	let winners: number[] = [];
	for (let i = 0; i < order.length; i++) {
		let lucky_number = order[i];
		for (let b = 0; b < boards.length; b++) {
			boards[b].select(lucky_number);

			if (boards[b].winner() && !winners.includes(b)) {
				winners.push(b);

				if (winners.length == boards.length) {
					return boards[b].score();
				}
			}
		}
	}

	return 0;
}

console.log(part1(order, boards.map((board) => board.clone())));
console.log(part2(order, boards.map((board) => board.clone())));
