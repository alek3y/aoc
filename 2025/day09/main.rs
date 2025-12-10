use std::fs;
use std::cmp;
use std::collections::HashSet;

#[derive(Clone, Copy, Eq, PartialEq, Hash, Debug)]
struct Point {
	x: isize, y: isize
}
type Edge = (Point, Point);

fn part1(coords: &[Point]) -> isize {
	let mut biggest_area = 1;
	for from in coords.iter() {
		for to in coords.iter() {
			let area = (to.y - from.y + 1).abs() * (to.x - from.x + 1).abs();
			biggest_area = cmp::max(biggest_area, area);
		}
	}
	biggest_area
}

fn inside(edges: (&Vec<Edge>, &HashSet<Edge>), point: &Point) -> bool {
	let (col_edges, row_edges) = edges;

	let possible_edges = col_edges
		.iter().rev()
		.skip_while(|(a, _)| point.x < a.x);	// Skips edges with x value that is over point.x

	let mut edges = 0;
	let mut last_corner_met: Option<Edge> = None;
	for &(a, b) in possible_edges {
		if !((a.y <= point.y && point.y <= b.y) || (b.y <= point.y && point.y <= a.y)) {
			continue;
		} else if a.x == point.x {
			return true;
		}

		let corner_met = if a.y == point.y {
			Some((a, b))
		} else if b.y == point.y {
			Some((b, a))
		} else {
			None
		};

		// A column wise edge should only be counted if point is not the second
		// occurring corner of the row wise edge, otherwise the direction of the
		// start and end column wise edges which are adjacent to the two corners
		// must go in the same direction (i.e. have both an obtuse or acute angle).
		let not_in_row_or_both_obtuse_corners = corner_met.zip(last_corner_met).is_none_or(
			|((corner, other), (last_corner, last_other))| {
				(
					!row_edges.contains(&(corner, last_corner))
					&& !row_edges.contains(&(last_corner, corner))
				) || (
					(other.y < point.y) == (last_other.y < point.y)
				)
			}
		);

		if not_in_row_or_both_obtuse_corners {
			edges += 1;
		}
		last_corner_met = corner_met;
	}
	edges % 2 != 0
}

fn part2(coords: &[Point], precision: usize) -> Option<isize> {
	let mut col_edges = Vec::new();
	let mut row_edges = HashSet::new();
	for i in 0..coords.len() {
		let a = coords[i];
		let b = coords[(i+1) % coords.len()];

		if a.x == b.x {
			col_edges.push((a, b));
		} else if a.y == b.y {
			row_edges.insert((a, b));
		}
	}
	col_edges.sort_by_key(|edge| edge.0.x);

	let mut rectangles = Vec::new();
	for (i, a) in coords.iter().enumerate() {
		for b in coords.iter().skip(i+1) {
			let area = ((b.y - a.y).abs() + 1) * ((b.x - a.x).abs() + 1);
			rectangles.push((a, b, area));
		}
	}
	rectangles.sort_by_key(|&(_, _, area)| area);

	'outer:
	for &(a, b, area) in rectangles.iter().rev() {
		if !inside((&col_edges, &row_edges), &(Point {x: a.x, y: b.y}))
			|| !inside((&col_edges, &row_edges), &(Point {x: b.x, y: a.y})) {	// Check opposite corners
			continue;
		}

		// Perimeter of the rectangle
		let y_range = cmp::min(a.y, b.y)+1..cmp::max(a.y, b.y);
		let x_range = cmp::min(a.x, b.x)+1..cmp::max(a.x, b.x);

		for i in 0..precision {	// Interlace coordinates so that small gaps are checked later
			for y in y_range.clone().skip(i).step_by(precision) {	// Left edge
				if !inside((&col_edges, &row_edges), &(Point {x: a.x, y})) {
					continue 'outer;
				}
			}
			for y in y_range.clone().skip(i).step_by(precision) {	// Right edge
				if !inside((&col_edges, &row_edges), &(Point {x: b.x, y})) {
					continue 'outer;
				}
			}
			for x in x_range.clone().skip(i).step_by(precision) {	// Top edge
				if !inside((&col_edges, &row_edges), &(Point {x, y: a.y})) {
					continue 'outer;
				}
			}
			for x in x_range.clone().skip(i).step_by(precision) {	// Bottom edge
				if !inside((&col_edges, &row_edges), &(Point {x, y: b.y})) {
					continue 'outer;
				}
			}
		}
		return Some(area);
	}
	None
}

fn main() {
	let input = fs::read_to_string("input.txt").unwrap();

	let coords: Vec<Point> = input
		.trim()
		.split('\n')
		.filter_map(|line| line.trim().split_once(','))
		.filter_map(|(x, y)| {
			Some(Point {
				x: x.parse().ok()?,
				y: y.parse().ok()?,
			})
		})
		.collect();

	println!("{}", part1(&coords));
	println!("{}", part2(&coords, 10).unwrap());
}
