use std::fs;
use std::collections::{HashSet, HashMap};

fn part1(map: &[&str], start: (usize, usize)) -> usize {
	let mut splits = HashSet::new();
	let mut beams = vec![start];
	let mut generated = HashSet::new();
	while let Some((x, y)) = beams.pop() {
		for y_spread in y..map.len() {
			if &map[y_spread][x..x+1] == "^" {
				let beam = (x-1, y_spread);
				if x > 0 && generated.insert(beam) {
					beams.push(beam);
				}

				let beam = (x+1, y_spread);
				if x < map[y_spread].len()-1 && generated.insert(beam) {
					beams.push(beam);
				}

				splits.insert((x, y_spread));
				break;
			}
		}
	}
	splits.len()
}

fn part2(map: &[&str], start: (usize, usize), memo: &mut HashMap<(usize, usize), usize>) -> usize {
	if memo.contains_key(&start) {
		return memo[&start];
	}

	let mut timelines = 0;

	let (x, y) = start;
	for y_spread in y..map.len() {
		if &map[y_spread][x..x+1] == "^" {
			if x > 0 {
				timelines += part2(map, (x-1, y_spread), memo);
			}
			if x < map[y_spread].len()-1 {
				timelines += part2(map, (x+1, y_spread), memo);
			}
			break;
		}
	}

	if timelines == 0 {
		timelines += 1;	// Only count timelines that haven't met a splitter
	}

	memo.insert(start, timelines);
	timelines
}


fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let map: Vec<_> = input
		.trim()
		.split('\n')
		.map(str::trim)
		.collect();

	let mut start = (0, 0);
	'outer:
	for (y, row) in map.iter().enumerate() {
		for (x, symbol) in row.chars().enumerate() {
			if symbol == 'S' {
				start = (x, y);
				break 'outer;
			}
		}
	}

	println!("{}", part1(&map, start));

	let mut memo = HashMap::new();
	println!("{}", part2(&map, start, &mut memo));
}
