use std::fs;

fn parse_instruction(instruction: &str) -> (&str, i64) {
	let mut pair = instruction.split(" ");
	let command = pair.next().unwrap();
	let argument: i64 = pair.next()
		.and_then(|s| s.parse().ok())
		.unwrap();

	(command, argument)
}

fn part1(input: &Vec<&str>) -> i64 {
	let mut position = 0;
	let mut depth = 0;

	for line in input {
		let (direction, steps) = parse_instruction(line);

		match direction {
			"forward" => position += steps,
			"up" => depth -= steps,
			"down" => depth += steps,
			_ => ()
		}
	}

	position * depth
}

fn part2(input: &Vec<&str>) -> i64 {
	let mut position = 0;
	let mut depth = 0;
	let mut aim = 0;

	for line in input {
		let (direction, steps) = parse_instruction(line);

		match direction {
			"forward" => {
				position += steps;
				depth += aim * steps;
			},
			"up" => aim -= steps,
			"down" => aim += steps,
			_ => ()
		}
	}

	position * depth
}

fn main() {
	let contents = fs::read_to_string("input.txt").unwrap();

	let mut input = Vec::new();
	for line in contents.lines() {
		input.push(line);
	}

	println!("{}", part1(&input));
	println!("{}", part2(&input));
}
