use std::fs;

fn part1(banks: &Vec<Vec<isize>>) -> isize {
	banks
		.iter()
		.fold(0, |acc, bank| {
			let first = bank.iter()
				.rev().skip(1).rev()
				.max()
				.unwrap();
			let second = bank.iter()
				.skip_while(|&x| x < first).skip(1)
				.max()
				.unwrap();
			acc + (first * 10 + second)
		})
}

fn part2(banks: &Vec<Vec<isize>>) -> isize {
	let n_batteries = 12;
	banks
		.iter()
		.fold(0, |acc, bank| {
			let mut joltage = 0;
			let mut batteries_search = 0;
			for batteries_left in (0..n_batteries).rev() {
				let mut battery_index = batteries_search;
				for i in batteries_search..(bank.len() - batteries_left) {
					if bank[i] > bank[battery_index] {
						battery_index = i;
					}
				}
				batteries_search = battery_index+1;
				joltage += bank[battery_index] * 10isize.pow(batteries_left as u32);
			}
			acc + joltage
		})
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let banks: Vec<Vec<isize>> = input.trim()
		.split('\n')
		.map(|line| line
			.trim()
			.chars()
			.filter_map(|digit| digit.to_digit(10))
			.map(|digit| digit as isize)
			.collect())
		.collect();

	println!("{}", part1(&banks));
	println!("{}", part2(&banks));
}
