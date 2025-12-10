use std::fs;
use std::cmp::{self, Ordering};
use std::collections::{HashSet, HashMap};

fn least_presses(state: &Vec<bool>, buttons: &HashSet<Vec<usize>>, pressed: &HashSet<Vec<usize>>, memo: &mut HashMap<(Vec<bool>, Vec<Vec<usize>>), usize>) -> usize {
	let mut memo_key = (
		state.clone(),
		pressed.iter()
			.map(|buttons| {
				let mut sorted_buttons = buttons.clone();
				sorted_buttons.sort();
				sorted_buttons
			})
			.collect::<Vec<_>>()
	);
	memo_key.1.sort_by(|a, b| {	// For the pressed buttons to be comparable they need to be normalized equally
		if a.len() < b.len() {
			Ordering::Less
		} else if a.len() > b.len() {
			Ordering::Greater
		} else {
			for (item_a, item_b) in a.iter().zip(b.iter()) {
				if item_a < item_b {
					return Ordering::Less;
				} else if item_a > item_b {
					return Ordering::Greater;
				}
			}
			Ordering::Equal
		}
	});

	if memo.contains_key(&memo_key) {
		return memo[&memo_key];
	} else if state.iter().all(|button| !button) {	// Going from empty to state or state to empty is the same process
		memo.insert(memo_key, pressed.len());
		return pressed.len();
	}

	let mut least = usize::MAX;
	for combo in buttons.difference(&pressed) {
		let mut pressed = pressed.clone();
		pressed.insert(combo.clone());

		let mut state = state.clone();
		for &button in combo {
			state[button] = !state[button];
		}

		let presses = least_presses(&state, buttons, &pressed, memo);
		least = cmp::min(least, presses);
	}

	memo.insert(memo_key, least);
	least
}

fn part1(machines: &[(Vec<bool>, HashSet<Vec<usize>>, Vec<usize>)]) -> usize {
	machines
		.iter()
		.fold(0, |acc, (state, buttons, _)| {
			let mut memo = HashMap::new();
			acc + least_presses(state, buttons, &HashSet::new(), &mut memo)
		})
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let machines: Vec<_> = input
		.trim()
		.split('\n')
		.map(|line| {
			let (state, suffix) = line.trim().split_once(' ').unwrap();
			let (prefix, joltage) = suffix.rsplit_once(' ').unwrap();

			let state: Vec<bool> = state
				.chars()
				.collect::<Vec<_>>()
				.into_iter()
				.skip(1).rev()
				.skip(1).rev()
				.map(|button| button == '#')
				.collect();

			let wiring: HashSet<Vec<usize>> = prefix
				.split(' ')
				.map(|buttons|
					buttons[1..buttons.len()-1]
						.split(',')
						.filter_map(|button| button.parse().ok())
						.collect::<Vec<_>>())
				.collect();

			let joltage: Vec<usize> = joltage[1..joltage.len()-1]
				.split(',')
				.filter_map(|number| number.parse().ok())
				.collect();

			(state, wiring, joltage)
		})
		.collect();

	println!("{}", part1(&machines));
}
