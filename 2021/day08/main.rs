use std::fs;
use std::collections::HashMap;

type Display<'a> = (Vec<&'a str>, Vec<&'a str>);

const DIGITS: &[&[u8]] = &[
	&[1, 1, 1, 0, 1, 1, 1],	// 0
	&[0, 0, 1, 0, 0, 1, 0],	// 1
	&[1, 0, 1, 1, 1, 0, 1],	// 2
	&[1, 0, 1, 1, 0, 1, 1],	// 3
	&[0, 1, 1, 1, 0, 1, 0],	// 4
	&[1, 1, 0, 1, 0, 1, 1],	// 5
	&[1, 1, 0, 1, 1, 1, 1],	// 6
	&[1, 0, 1, 0, 0, 1, 0],	// 7
	&[1, 1, 1, 1, 1, 1, 1],	// 8
	&[1, 1, 1, 1, 0, 1, 1]	// 9
];

const UNIQUE: &[usize] = &[1, 4, 7, 8];

const SEGMENTS_FREQ: &[usize] = &[8, 6, 8, 7, 4, 9, 7];

fn part1(input: &Vec<Display>) -> usize {
	let mut unique_lengths = Vec::new();
	for digit in UNIQUE {
		unique_lengths.push(
			DIGITS[*digit]
				.iter()
				.filter(|s| **s == 1u8)
				.count()
		);
	}

	let mut unique = 0usize;
	for display in input {
		unique += display.1
			.iter()
			.filter(|d| unique_lengths.contains(&d.len()))
			.count();
	}

	unique
}

fn part2(input: &Vec<Display>) -> u64 {
	let mut sum = 0u64;

	for display in input {

		// 3 segments can be found with their frequency because it's unique
		let mut frequency: HashMap<char, usize> = HashMap::new();
		for digit in &display.0 {
			for segment in digit.chars() {
				if !frequency.contains_key(&segment) {
					frequency.insert(segment, 0);
				}
				frequency.insert(segment, frequency[&segment]+1);
			}
		}

		let mut associated = vec!['\0'; DIGITS[0].len()];

		// Association of the 3 segments with the matching segments of SEGMENTS_FREQ
		for (segment, count) in frequency.iter() {
			let unique_frequencies: Vec<(usize, &usize)> = SEGMENTS_FREQ
				.iter()
				.enumerate()
				.filter(|f| f.1 == count)
				.collect();

			if unique_frequencies.len() > 1 {
				continue;
			}

			associated[unique_frequencies[0].0] = *segment;
		}

		// Now, with the known segments and the unique digits (starting from the digit 1 which
		// has two segments, one of which is known) it should be possible to find the remaining
		// combinations
		for unique_digit in UNIQUE {
			let unique_segment_amount = DIGITS[*unique_digit]
				.iter()
				.filter(|s| **s == 1u8)
				.count();

			let segments_chars: Vec<char> = display.0
				.iter()
				.find(|d| d.len() == unique_segment_amount)
				.unwrap()
				.chars()
				.collect();

			let first_unknown = segments_chars
				.iter()
				.filter(|s| !associated.contains(s))
				.next()
				.unwrap();

			let mut unknown_segments = Vec::new();
			for (i, segment) in DIGITS[*unique_digit].iter().enumerate() {
				if *segment == 0 || associated[i] != '\0' {
					unknown_segments.push(0);
				} else {
					unknown_segments.push(1);
				}
			}

			let first_unknown_index = unknown_segments
				.iter()
				.enumerate()
				.find(|s| *s.1 == 1)
				.unwrap().0;
			associated[first_unknown_index] = *first_unknown;
		}

		// The final numbers can now be decoded and merged with the help
		// of a string (because I'm lazy)
		let mut value = String::new();
		for final_digit in &display.1 {
			let mut final_digit_segments = vec![0; DIGITS[0].len()];

			for segment_char in final_digit.chars() {
				let segment_index = associated
					.iter()
					.enumerate()
					.find(|a| *a.1 == segment_char)
					.unwrap().0;
				final_digit_segments[segment_index] = 1;
			}

			let matching_digit_to_segment = DIGITS
				.iter()
				.enumerate()
				.find(|d| *d.1 == final_digit_segments.as_slice())
				.unwrap().0;

			value = format!("{}{}", value, matching_digit_to_segment);
		}

		sum += value.parse::<u64>().unwrap();
	}

	sum
}

fn main() {
	let file = fs::read_to_string("input.txt").unwrap();

	let mut lines: Vec<&str> = file.split("\n").collect();
	lines.pop();

	let mut input: Vec<Display> = Vec::new();
	for line in &lines {
		let mut pair = line.split(" | ");
		input.push((
			pair.next().unwrap().split(" ").collect(),
			pair.next().unwrap().split(" ").collect()
		));
	}

	println!("{}", part1(&input));
	println!("{}", part2(&input));
}
