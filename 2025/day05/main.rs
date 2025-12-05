use std::fs;

fn part1(ranges: &[(isize, isize)], ingredients: &[isize]) -> usize {
	ingredients
		.iter()
		.fold(0, |acc, ingredient| {
			acc + ranges
				.iter()
				.find(|(start, end)| start <= ingredient && ingredient <= end)
				.is_some() as usize
		})
}

fn part2(ranges: &mut [(isize, isize)]) -> isize {
	ranges.sort_by(|(start_0, end_0), (start_1, end_1)| {	// Sort by start first and end second
		if start_0 == start_1 {
			end_0.cmp(end_1)
		} else {
			start_0.cmp(start_1)
		}
	});

	let mut deduped = Vec::new();
	let (mut start, mut end) = ranges[0];
	for range in ranges.iter() {
		if range.0 <= end && end < range.1 {
			end = range.1;
		} else if range.0 > end {
			deduped.push((start, end));
			start = range.0;
			end = range.1;
		}
	}
	deduped.push((start, end));	// Flush last range built

	deduped
		.iter()
		.map(|(start, end)| end - start + 1)	// Number of ingredients in the range
		.fold(0, |acc, count| acc + count)
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let (ranges, ingredients) = input
		.trim()
		.split_once("\n\n")
		.unwrap();

	let mut ranges: Vec<(isize, isize)> = ranges
		.split('\n')
		.filter_map(|line| line.trim().split_once('-'))
		.filter_map(|(start, end)| Some((start.parse().ok()?, end.parse().ok()?)))
		.collect();

	let ingredients: Vec<isize> = ingredients
		.split('\n')
		.filter_map(|line| line.trim().parse().ok())
		.collect();

	println!("{}", part1(&ranges, &ingredients));
	println!("{}", part2(&mut ranges));
}
