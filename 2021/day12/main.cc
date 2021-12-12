#include <iostream>
#include <vector>
#include <fstream>
#include <functional>
#include <map>
#include <cctype>
#include <algorithm>
using namespace std;

struct Cave {
	string name;
	vector<string> adjacent;
	size_t visited = 0;

	Cave() {}

	Cave(string name) {
		this->name = name;
	}

	bool is_small() {
		string name_lower;
		name_lower.resize(this->name.size());
		transform(
			this->name.begin(),
			this->name.end(),
			name_lower.begin(),
			[](char c) -> char { return tolower(c); }
		);

		return this->name == name_lower;
	}
};

size_t part1(map<string, Cave> &caves, string current = "start") {
	bool has_destination = false;
	size_t paths = 0;

	caves[current].visited++;
	for (string adjacent : caves[current].adjacent) {
		if (adjacent == "end") {
			has_destination = true;
			continue;
		}

		if (caves[adjacent].is_small() && caves[adjacent].visited > 0) {
			continue;
		}

		paths += part1(caves, adjacent);
	}

	caves[current].visited = 0;
	if (has_destination) {
		paths++;
	}
	return paths;
}

size_t part2(map<string, Cave> &caves, string current = "start", bool visited_twice = false) {
	bool has_destination = false;
	size_t paths = 0;

	if (++caves[current].visited > 1 && caves[current].is_small()) {
		visited_twice = true;
	}

	for (string adjacent : caves[current].adjacent) {
		if (adjacent == "start") {
			continue;
		} else if (adjacent == "end") {
			has_destination = true;
			continue;
		}

		if (caves[adjacent].is_small() && caves[adjacent].visited > 0 && visited_twice) {
			continue;
		}

		paths += part2(caves, adjacent, visited_twice);
	}

	caves[current].visited--;
	if (has_destination) {
		paths++;
	}
	return paths;
}

int main(void) {
	ifstream file("input.txt");

	vector<pair<string, string>> input;

	string line;
	while (file >> line) {
		size_t dash = line.find("-");
		string from = line.substr(0, dash);
		string to = line.substr(dash+1);

		input.push_back(make_pair(from, to));
	}

	map<string, Cave> caves;
	for (auto connection : input) {
		string first = connection.first, second = connection.second;

		if (caves.count(first) == 0) {
			caves[first] = Cave(first);
		}
		if (caves.count(second) == 0) {
			caves[second] = Cave(second);
		}

		caves[first].adjacent.push_back(second);
		caves[second].adjacent.push_back(first);
	}

	cout << part1(caves) << endl;
	cout << part2(caves) << endl;
}
