#include <iostream>
#include <list>
#include <map>
#include <fstream>
using namespace std;

size_t part1(list<char> polymer, map<pair<char, char>, char> &mappings) {
	for (size_t s = 0; s < 10; s++) {
		char previous = *polymer.begin();
		for (auto it = ++polymer.begin(); it != polymer.end(); it++) {
			auto combination = make_pair(previous, *it);
			polymer.insert(it, mappings[combination]);
			previous = *it;
		}
	}

	map<char, size_t> frequency;
	for (auto it = polymer.begin(); it != polymer.end(); it++) {
		if (frequency.count(*it) == 0) {
			frequency[*it] = 0;
		}

		frequency[*it]++;
	}

	size_t most_common = 0, least_common = frequency.begin()->second;
	for (auto pair : frequency) {
		if (pair.second > most_common) {
			most_common = pair.second;
		} else if (pair.second < least_common) {
			least_common = pair.second;
		}
	}

	return most_common - least_common;
}

size_t part2(list<char> polymer, map<pair<char, char>, char> &mappings) {
	map<char, size_t> frequency;

	char previous = *polymer.begin();
	frequency[previous] = 1;	// left of pairs are not counted

	map<pair<char, char>, size_t> pairs;
	for (auto it = ++polymer.begin(); it != polymer.end(); it++) {
		auto pair = make_pair(previous, *it);
		if (pairs.count(pair) == 0) {
			pairs[pair] = 0;
		}

		pairs[pair]++;
		previous = *it;
	}

	for (size_t s = 0; s < 40; s++) {
		auto pairs_previous = pairs;
		for (auto pair : pairs_previous) {
			if (pair.second == 0) {
				continue;
			}

			pairs[pair.first] -= pair.second;

			auto left = make_pair(pair.first.first, mappings[pair.first]);
			auto right = make_pair(mappings[pair.first], pair.first.second);
			if (pairs.count(left) == 0) {
				pairs[left] = 0;
			}
			if (pairs.count(right) == 0) {
				pairs[right] = 0;
			}

			pairs[left] += pair.second;
			pairs[right] += pair.second;
		}
	}

	for (auto pair : pairs) {
		if (frequency.count(pair.first.second) == 0) {
			frequency[pair.first.second] = 0;
		}

		frequency[pair.first.second] += pair.second;
	}

	size_t most_common = 0, least_common = frequency.begin()->second;
	for (auto pair : frequency) {
		if (pair.second > most_common) {
			most_common = pair.second;
		} else if (pair.second < least_common) {
			least_common = pair.second;
		}
	}

	return most_common - least_common;
}

int main(void) {
	fstream file("input.txt");

	string initial;
	file >> initial;
	file.ignore(2);

	list<char> polymer;
	for (char letter : initial) {
		polymer.push_back(letter);
	}

	map<pair<char, char>, char> mappings;
	string line;
	while (getline(file, line)) {
		auto pair = make_pair(line[0], line[1]);
		mappings[pair] = line[line.size()-1];
	}

	cout << part1(polymer, mappings) << endl;
	cout << part2(polymer, mappings) << endl;
}
