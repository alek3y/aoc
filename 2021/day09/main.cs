using System;
using System.IO;
using System.Linq;
using System.Collections.Generic;

class Day09 {
	static bool isLowest(int y, int x, int[,] input) {
		bool isLowest = true;
		int num = input[y,x];

		if (y > 0) {
			isLowest = isLowest && num < input[y-1,x];
		}
		if (y < input.GetLength(0)-1) {
			isLowest = isLowest && num < input[y+1,x];
		}
		if (x > 0) {
			isLowest = isLowest && num < input[y,x-1];
		}
		if (x < input.GetLength(1)-1) {
			isLowest = isLowest && num < input[y,x+1];
		}

		return isLowest;
	}

	static int part1(int[,] input) {
		int risk = 0;

		for (int y = 0; y < input.GetLength(0); y++) {
			for (int x = 0; x < input.GetLength(1); x++) {
				if (isLowest(y, x, input)) {
					risk += 1 + input[y,x];
				}
			}
		}

		return risk;
	}

	static Tuple<int, int> basin(int y, int x, int[,] input) {
		if (isLowest(y, x, input)) {
			return new Tuple<int, int>(y, x);
		}

		Tuple<int, int> smallest = Tuple.Create(y, x);
		if (y > 0 && input[y-1,x] < input[smallest.Item1,smallest.Item2]) {
			smallest = Tuple.Create(y-1, x);
		}
		if (y < input.GetLength(0)-1 && input[y+1,x] < input[smallest.Item1,smallest.Item2]) {
			smallest = Tuple.Create(y+1, x);
		}
		if (x > 0 && input[y,x-1] < input[smallest.Item1,smallest.Item2]) {
			smallest = Tuple.Create(y, x-1);
		}
		if (x < input.GetLength(1)-1 && input[y,x+1] < input[smallest.Item1,smallest.Item2]) {
			smallest = Tuple.Create(y, x+1);
		}

		return basin(smallest.Item1, smallest.Item2, input);
	}

	static int part2(int[,] input) {
		Dictionary<Tuple<int, int>, int> areas = new Dictionary<Tuple<int, int>, int>();

		for (int y = 0; y < input.GetLength(0); y++) {
			for (int x = 0; x < input.GetLength(1); x++) {
				if (input[y,x] != 9) {
					Tuple<int, int> currentBasin = basin(y, x, input);
					if (! areas.ContainsKey(currentBasin)) {
						areas[currentBasin] = 0;
					}
					areas[currentBasin] += 1;
				}
			}
		}

		int mul = 1;
		int[] values = areas.Values.ToArray();
		Array.Sort(values);
		for (int i = values.Length-3; i < values.Length; i++) {
			mul *= values[i];
		}

		return mul;
	}

	static void Main(string[] args) {
		string[] lines = File.ReadAllText("input.txt").Split("\n");
		int[,] input = new int[lines.Length-1, lines[0].Length];

		for (int i = 0; i < lines.Length; i++) {
			for (int j = 0; j < lines[i].Length; j++) {
				input[i,j] = lines[i][j] - '0';
			}
		}

		Console.WriteLine(part1(input));
		Console.WriteLine(part2(input));
	}
}
