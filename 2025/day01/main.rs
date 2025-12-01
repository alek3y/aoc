use std::fs;

fn part1(rotations: &[isize]) -> isize {
	rotations
		.iter()
		.fold((50, 0), |(dial, zeroes), rotations| {
			let dial = (dial + rotations).rem_euclid(100);
			(dial, zeroes + (dial == 0) as isize)
		}).1
}

fn part2(rotations: &[isize]) -> isize {
	rotations
		.iter()
		.fold((50, 0), |(dial, zeroes), rotations| {
			let sum = dial + rotations;
			(
				sum.rem_euclid(100),
				zeroes + sum.abs() / 100
					+ (sum.signum() != dial.signum() && dial != 0) as isize	// Also count sign changes (except when dial is 0)
			)
		}).1
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let rotations: Vec<isize> = input.trim()
		.split('\n')
		.map(str::trim)
		.map(|line| {
			let clicks: isize = line[1..].parse().unwrap();
			if &line[..1] == "L" {
				-clicks
			} else {
				clicks
			}
		})
		.collect();

	println!("{}", part1(&rotations));
	println!("{}", part2(&rotations));
}
