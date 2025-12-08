use std::fs;
use std::collections::{HashSet, HashMap};
use std::cmp::Ordering;

fn distance(a: &(isize, isize, isize), b: &(isize, isize, isize)) -> f64 {
	(((a.0 - b.0).pow(2)
		+ (a.1 - b.1).pow(2)
		+ (a.2 - b.2).pow(2)) as f64).sqrt()
}

fn connect(coords: &[(isize, isize, isize)], snapshot_sum_after: usize) -> (usize, isize) {
	let mut pairs = Vec::new();
	for i in 0..coords.len() {
		for j in i+1..coords.len() {
			pairs.push((coords[i], coords[j], distance(&coords[i], &coords[j])));
		}
	}
	pairs.sort_by(|(_, _, distance_0), (_, _, distance_1)| {
		if distance_1 - distance_0 < 2.0 * f64::EPSILON {
			Ordering::Equal
		} else if distance_0 < distance_1 {
			Ordering::Less
		} else {
			Ordering::Greater
		}
	});

	let mut circuits = Vec::new();
	let mut groupings = HashMap::new();
	for coord in coords.iter() {
		groupings.insert(coord, circuits.len());
		circuits.push(HashSet::from([coord]));
	}

	let mut snapshot_sum = 0;
	let mut last_two = 0;
	for (i, (from, to, _)) in pairs.iter().enumerate() {
		if groupings[from] == groupings[to] {
			continue;
		}

		let target_group;
		let source_group;
		if circuits[groupings[from]].len() > circuits[groupings[to]].len() {
			target_group = groupings[from];
			source_group = groupings[to];
		} else {
			target_group = groupings[to];
			source_group = groupings[from];
		}

		{
			let source_cloned = circuits[source_group].clone();
			circuits[target_group].extend(source_cloned);
		}

		for source_coord in &circuits[source_group] {
			groupings.insert(source_coord, target_group);
		}
		circuits[source_group].clear();

		if i+1 == snapshot_sum_after {
			let mut snapshot = circuits.clone();
			snapshot.sort_by_key(|circuit| circuit.len());
			snapshot_sum = snapshot
				.iter()
				.rev()
				.take(3)
				.fold(1, |acc, circuit| {
					acc * circuit.len()
				});
		}
		last_two = from.0 * to.0;
	}

	(snapshot_sum, last_two)
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let coords: Vec<_> = input
		.trim()
		.split('\n')
		.map(|line| {
			let split: Vec<_> = line
				.trim()
				.splitn(3, ',')
				.filter_map(|number| number.parse::<isize>().ok())
				.collect();

			(split[0], split[1], split[2])
		})
		.collect();

	let (part1, part2) = connect(&coords, 1000);	// Should stop after 10 pairs with sample.txt
	println!("{}", part1);
	println!("{}", part2);
}
