<?php

class Point {
	public $x;
	public $y;

	public function __construct($x, $y) {
		$this->x = $x;
		$this->y = $y;
	}

	public static function parse($point_str) {
		$coordinates = explode(",", $point_str);
		return new Point((int) $coordinates[0], (int) $coordinates[1]);
	}
}

class Line {
	public $start;
	public $end;

	public function __construct($start, $end) {
		$this->start = $start;
		$this->end = $end;
	}

	public static function parse($line_str) {
		$points = explode(" -> ", $line_str);
		return new Line(Point::parse($points[0]), Point::parse($points[1]));
	}
}

function clone_deep($object) {
	return unserialize(serialize($object));
}

function draw_lines(&$grid, $offset, $lines, $allow_diagonals) {
	foreach ($lines as $line) {
		$start = $line->start;
		$end = $line->end;

		if ($start->y == $end->y) {
			$min = min($start->x, $end->x);
			$max = max($start->x, $end->x);
			for ($i = $min - $offset->x; $i <= $max - $offset->x; $i++) {
				$grid[$start->y - $offset->y][$i] += 1;
			}
		} elseif ($start->x == $end->x) {
			$min = min($start->y, $end->y);
			$max = max($start->y, $end->y);
			for ($i = $min - $offset->y; $i <= $max - $offset->y; $i++) {
				$grid[$i][$start->x - $offset->x] += 1;
			}
		} elseif ($allow_diagonals) {
			$y_step = $line->start->y > $line->end->y ? -1 : 1;
			$x_step = $line->start->x > $line->end->x ? -1 : 1;

			// End coordinate is inclusive
			$end->x += $x_step;
			$end->y += $y_step;

			$current = clone_deep($start);
			while ($current->x != $end->x && $current->y != $end->y) {
				$grid[$current->y - $offset->y][$current->x - $offset->x] += 1;
				$current->x += $x_step;
				$current->y += $y_step;
			}
		}
	}
}

function count_intersections($grid) {
	$intersections = 0;
	for ($y = 0; $y < count($grid); $y++) {
		for ($x = 0; $x < count($grid[$y]); $x++) {
			if ($grid[$y][$x] > 1) {
				$intersections++;
			}
		}
	}
	return $intersections;
}

function print_grid($grid) {
	for ($y = 0; $y < count($grid); $y++) {
		for ($x = 0; $x < count($grid[$y]); $x++) {
			$cell = $grid[$y][$x];
			if ($cell == 0) {
				echo ".";
			} else {
				echo $cell;
			}
		}
		echo "\n";
	}
}

function part1($grid, $offset, $input) {
	draw_lines($grid, $offset, $input, false);
	return count_intersections($grid);
}

function part2($grid, $offset, $input) {
	draw_lines($grid, $offset, $input, true);
	return count_intersections($grid);
}

$file = file_get_contents("input.txt");

$input = array();
foreach (explode("\n", $file) as $line) {
	if (strlen($line) > 0) {
		array_push($input, Line::parse($line));
	}
}

$xs = array();
$ys = array();
foreach ($input as $line) {
	array_push($xs, $line->start->x);
	array_push($ys, $line->start->y);

	array_push($xs, $line->end->x);
	array_push($ys, $line->end->y);
}

$diagonal = new Line(
	new Point(min($xs), min($ys)),
	new Point(max($xs), max($ys)),
);

$grid = array();
$height = $diagonal->end->y - $diagonal->start->y + 1;
$width = $diagonal->end->x - $diagonal->start->x + 1;
for ($i = 0; $i < $height; $i++) {
	array_push($grid, array_fill(0, $width, 0));
}

echo part1($grid, $diagonal->start, $input) . "\n";
echo part2($grid, $diagonal->start, $input) . "\n";
?>
