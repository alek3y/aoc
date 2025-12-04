use std::fs;

fn neighbors(map: &Vec<Vec<char>>, x: isize, y: isize) -> usize {
	let mut neighbors = 0;
	for dy in -1..=1 {
		if y + dy < 0 || y + dy >= map.len() as isize {
			continue;
		}

		for dx in -1..=1 {
			if x + dx < 0 || x + dx >= map[(y + dy) as usize].len() as isize {
				continue;
			}

			if map[(y + dy) as usize][(x + dx) as usize] == '@' {
				neighbors += 1;
			}
		}
	}
	neighbors
}

fn part1(map: &Vec<Vec<char>>) -> usize {
	let mut accessible_rolls = 0;
	for y in 0..map.len() {
		for x in 0..map[y].len() {
			if map[y][x] != '@' {
				continue;
			}

			if neighbors(map, x as isize, y as isize) <= 4 {
				accessible_rolls += 1;
			}
		}
	}
	accessible_rolls
}

fn part2(map: &mut Vec<Vec<char>>) -> usize {
	let mut removed = 0;
	loop {
		let mut refresh = false;
		let state: Vec<_> = map.iter().map(|row| row.clone()).collect();

		for y in 0..state.len() {
			for x in 0..state[y].len() {
				if state[y][x] != '@' {
					continue;
				}

				if neighbors(&state, x as isize, y as isize) <= 4 {
					map[y][x] = '.';
					removed += 1;
					refresh = true;
				}
			}
		}

		if !refresh {
			break;
		}
	}
	removed
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let mut map: Vec<Vec<char>> = input.trim()
		.split('\n')
		.map(|line| line.trim().chars().collect())
		.collect();

	println!("{}", part1(&map));
	println!("{}", part2(&mut map));
}
