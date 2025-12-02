use std::fs;
use std::collections::HashSet;

fn part1(ranges: &[(isize, isize)]) -> isize {
	ranges
		.iter()
		.fold(0, |mut acc, (start, end)| {
			let first_half_scale = 10isize.pow((end.ilog10()+1) / 2);
			for half_offset in 0..=((end - start) / first_half_scale + 1) {	// Tries at least once for when the difference is less than first_half_scale
				let first_half_increased = (start + half_offset * first_half_scale).to_string();
				if first_half_increased.len() <= 1 {	// Could never be a repeated string
					continue
				}

				let first_half = &first_half_increased[..first_half_increased.len()/2];
				let repeated: isize = (first_half.to_string() + first_half).parse().unwrap();

				if start <= &repeated && &repeated <= end {
					acc += repeated
				}
			}
			acc
		})
}

fn part2(ranges: &[(isize, isize)]) -> isize {
	ranges
		.iter()
		.fold(0, |mut acc, (start, end)| {
			let mut already_repeated = HashSet::new();
			let first_half_scale = 10isize.pow((end.ilog10()+1) / 2);
			for half_offset in 0..=((end - start) / first_half_scale + 1) {	// To be repeated it needs to increase *from* the first half and upwards
				let first_half_increased = (start + half_offset * first_half_scale).to_string();
				if first_half_increased.len() <= 1 {
					continue;
				}

				for section in 1..=first_half_increased.len()/2 {
					if first_half_increased.len() % section != 0 {
						continue;
					}

					let repeated: isize = first_half_increased[..section]
						.repeat(first_half_increased.len() / section)
						.parse()
						.unwrap();

					if already_repeated.contains(&repeated) {
						continue;
					}

					if start <= &repeated && &repeated <= end {
						acc += repeated;
						already_repeated.insert(repeated);
					}
				}
			}
			acc
		})
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let ranges: Vec<(isize, isize)> = input.trim()
		.split(',')
		.filter_map(|line| line.trim().split_once('-'))
		.filter_map(|(start, end)| Some((start.parse().ok()?, end.parse().ok()?)))
		.collect();

	println!("{}", part1(&ranges));
	println!("{}", part2(&ranges));
}
