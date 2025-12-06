use std::fs;
use std::iter::Sum;

fn reduce(numbers: &Vec<Vec<usize>>, ops: &[(usize, &str)]) -> usize {
	let mut total = 0;
	for (i, &(_, op)) in ops.iter().enumerate() {
		total += match op {
			"+" => usize::sum(numbers[i].iter()),
			"*" => numbers[i].iter().fold(1, |acc, number| acc*number),
			_ => unreachable!()
		};
	}
	total
}

fn part1(lines: &[&str], ops: &[(usize, &str)]) -> usize {
	let mut numbers = Vec::new();
	for (i, &(start, _)) in ops.iter().enumerate() {
		let end = ops
			.get(i+1)
			.map_or(lines[0].len(), |&(end, _)| end);	// Every untrimmed line has the same length

		let mut column = Vec::new();
		for line in lines {
			let number: usize = line[start..end]
				.trim()
				.parse()
				.unwrap();

			column.push(number);
		}
		numbers.push(column);
	}

	reduce(&numbers, ops)
}

fn part2(lines: &[&str], ops: &[(usize, &str)]) -> usize {
	let mut numbers = Vec::new();
	for (i, &(start, _)) in ops.iter().enumerate() {
		let end = ops
			.get(i+1)
			.map_or(lines[0].len(), |&(end, _)| end);

		let mut column = Vec::new();
		for j in start..end {
			let mut number = None;

			for line in lines {
				let symbol = &line[j..j+1];
				if symbol == " " {
					continue;
				}

				let digit = symbol
					.parse::<usize>()
					.unwrap();

				number = number
					.or(Some(0))
					.map(|number| number * 10 + digit);
			}

			if number.is_some() {
				column.push(number.unwrap());
			}
		}
		numbers.push(column);
	}

	reduce(&numbers, ops)
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let lines: Vec<_> = input
		.trim()
		.split('\n')
		.collect();

	let ops: Vec<_> = lines[lines.len()-1]
		.match_indices(|symbol| "+*".contains(symbol))
		.collect();

	let lines = &lines[..lines.len()-1];

	println!("{}", part1(&lines, &ops));
	println!("{}", part2(&lines, &ops));
}
