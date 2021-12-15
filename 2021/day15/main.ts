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

	let finite: Set<Node> = new Set();
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
		finite.delete(node);

		let adjacent = [
			graph[node.y-1]?.[node.x],
			graph[node.y+1]?.[node.x],
			graph[node.y][node.x+1],
			graph[node.y][node.x-1]
		].filter((node) => node !== undefined);

		for (let i = 0; i < adjacent.length; i++) {
			if (adjacent[i].visited) {
				continue;
			}

			let distance = node.distance + risks[adjacent[i].y][adjacent[i].x];
			if (distance < adjacent[i].distance) {
				adjacent[i].distance = distance;
				finite.add(adjacent[i]);
			}
		}

		let shortest = Infinity;
		let shortest_node: Node;
		for (let item of finite) {
			if (item.distance < shortest && !item.visited) {
				shortest = item.distance;
				shortest_node = item;
			}
		}

		queue.push(shortest_node!);
	}

	let last_row = graph[graph.length-1];
	return last_row[last_row.length-1].distance;
}

function part2(risks: number[][]): number {
	let entire_cave: number[][] = [];

	for (let i = 0; i < 5; i++) {
		for (let y = 0; y < risks.length; y++) {
			entire_cave.push([]);
			for (let j = 0; j < 5; j++) {
				for (let x = 0; x < risks[y].length; x++) {
					let new_risk = (risks[y][x] + i+j - 1) % 9 + 1;

					entire_cave[entire_cave.length-1].push(new_risk);
				}
			}
		}
	}

	return part1(entire_cave);
}

console.log(part1(risks));
console.log(part2(risks));
