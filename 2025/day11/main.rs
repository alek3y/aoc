use std::fs;
use std::collections::HashMap;

fn part1(source: &str, relations: &HashMap<&str, Vec<&str>>) -> usize {
	if source == "out" {
		return 1;
	}

	let mut paths = 0;
	if let Some(children) = relations.get(source) {
		for child in children {
			paths += part1(child, relations);
		}
	}
	paths
}

fn part2<'a>(source: &'a str, requirements: &HashMap<&'a str, bool>, relations: &HashMap<&'a str, Vec<&'a str>>, memo: &mut HashMap<(&'a str, Vec<(&'a str, bool)>), usize>) -> usize {
	let mut memo_key = (source, requirements.clone().into_iter().collect::<Vec<_>>());
	memo_key.1.sort_by_key(|&(source, _)| source);

	if memo.contains_key(&memo_key) {
		return memo[&memo_key];
	} else if source == "out" {
		let valid = requirements.values().all(|&satified| satified) as usize;
		memo.insert(memo_key, valid);
		return valid;
	}

	let mut paths = 0;
	if let Some(children) = relations.get(source) {
		for child in children {
			let mut requirements = requirements.clone();
			if requirements.contains_key(child) {
				requirements.insert(child, true);
			}

			paths += part2(child, &requirements, relations, memo);
		}
	}
	memo.insert(memo_key.clone(), paths);
	paths
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let pairs: Vec<(_, Vec<_>)> = input
		.trim()
		.split('\n')
		.filter_map(|line| line.trim().split_once(':'))
		.map(|(name, rest)| (name, rest.trim().split(' ').collect()))
		.collect();

	let mut relations = HashMap::new();
	for (parent, children) in pairs {
		relations.insert(parent, children);
	}

	println!("{}", part1("you", &relations));

	let mut memo = HashMap::new();
	let requirements = HashMap::from([
		("fft", false),
		("dac", false),
	]);
	println!("{}", part2("svr", &requirements, &relations, &mut memo));
}
