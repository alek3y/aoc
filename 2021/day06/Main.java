/*
Language: Java 1.5.0
Environment: MacOS X 10.4 "Tiger"
*/

import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;
import java.util.ArrayList;
import java.util.HashMap;
import java.math.BigInteger;

public class Main {
	public static int TIMER = 7;

	public static long part1(ArrayList<Integer> input) {
		ArrayList<Integer> fishes = new ArrayList<Integer>(input);

		for (int d = 0; d < 80; d++) {
			int born = 0;

			for (int i = 0; i < fishes.size(); i++) {
				if (fishes.get(i) == 0) {
					fishes.set(i, TIMER-1);
					born++;
				} else {
					fishes.set(i, fishes.get(i)-1);
				}
			}

			for (int i = 0; i < born; i++) {
				fishes.add(TIMER-1 + 2);
			}
		}
		return fishes.size();
	}

	public static BigInteger part2(ArrayList<Integer> input) {
		HashMap<Integer, BigInteger> next_day = new HashMap<Integer, BigInteger>();

		for (int i = 0; i <= TIMER-1 + 2; i++) {
			next_day.put(i, BigInteger.valueOf(0));
		}

		for (int i = 0; i < input.size(); i++) {
			int key = input.get(i);
			next_day.put(key, next_day.get(key).add(BigInteger.valueOf(1)));
		}

		for (int d = 0 ; d < 256; d++) {
			HashMap<Integer, BigInteger> last_day = new HashMap<Integer, BigInteger>(next_day);

			for (int key : last_day.keySet()) {
				BigInteger value = last_day.get(key);

				if (value.compareTo(BigInteger.valueOf(0)) == 1) {
					next_day.put(key, next_day.get(key).subtract(value));
					if (key == 0) {
						next_day.put(TIMER-1, next_day.get(TIMER-1).add(value));
						next_day.put(TIMER-1 + 2, next_day.get(TIMER-1 + 2).add(value));
					} else {
						next_day.put(key-1, next_day.get(key-1).add(value));
					}
				}
			}
		}

		BigInteger size = BigInteger.valueOf(0);
		for (int key : next_day.keySet()) {
			size = size.add(next_day.get(key));
		}

		return size;
	}

	public static void main(String[] args) throws FileNotFoundException {
		File file = new File("input.txt");
		Scanner reader = new Scanner(file);
		String[] input = reader.nextLine().split(",");

		ArrayList<Integer> fishes = new ArrayList<Integer>();
		for (String fish : input) {
			fishes.add(Integer.valueOf(fish));
		}

		System.out.println(part1(fishes));
		System.out.println(part2(fishes));
	}
}
