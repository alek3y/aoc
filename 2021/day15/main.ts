/*
Pathfinding algorithm: https://www.youtube.com/watch?v=wtdtkJgcYUM

Run with `deno run --allow-read main.ts`
*/

const contents = await Deno.readTextFile("input.txt");
const lines = contents.split("\n");
lines.pop();

let risks = lines.map((line) => line.split("").map(Number));

class Node {
	distance = Infinity;
	visited = false;
	y: number;
	x: number;

	constructor(y: number, x: number) {
		this.y = y;
		this.x = x;
	}
}

function part1(risks: number[][]): number {
	let graph = [];
	for (let y = 0; y < risks.length; y++) {
		graph.push(new Array(risks[y].length));
		for (let x = 0; x < risks[y].length; x++) {
			graph[y][x] = new Node(y, x);
		}
	}

	let origin = graph[0][0];
	origin.distance = 0;

	let queue: Node[] = [origin];
	while (true) {
		let node = queue.shift();
		if (node === undefined) {
			break;
		}

		if (node.visited) {
			continue;
		}
		node.visited = true;

		let adjacent = [
			graph[node.y-1]?.[node.x],
			graph[node.y+1]?.[node.x],
			graph[node.y][node.x+1],
			graph[node.y][node.x-1]
		].filter((node) => node !== undefined);

		for (let i = 0; i < adjacent.length; i++) {
			let distance = node.distance + risks[adjacent[i].y][adjacent[i].x];
			if (distance < adjacent[i].distance) {
				adjacent[i].distance = distance;
			}

			queue.push(adjacent[i]);
		}
	}

	let last_row = graph[graph.length-1];
	return last_row[last_row.length-1].distance;
}

console.log(part1(risks));
