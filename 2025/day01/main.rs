use std::fs;

fn part1(lines: &[&str]) -> isize {
	lines.iter()
		.map(|line| {
			let clicks: isize = line[1..].parse().unwrap();

			if &line[..1] == "L" {
				-clicks
			} else {
				clicks
			}
		})
		.fold((50, 0), |(sum, zeroes), rotations| {
			let sum = (sum + rotations).rem_euclid(100);
			(sum, zeroes + (sum == 0) as isize)
		}).1
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let lines: Vec<&str> = input.trim()
		.split('\n')
		.map(str::trim)
		.collect();

	println!("{}", part1(&lines));
}
