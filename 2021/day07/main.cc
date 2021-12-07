#include <iostream>
#include <fstream>
#include <vector>
#include <array>
#include <sstream>
#include <map>
#include <algorithm>
using namespace std;

size_t part1(vector<int> input) {
	map<int, size_t> fuels;

	for (int num : input) {
		if (fuels.count(num) > 0) {
			continue;
		}

		size_t sum = 0;
		for (int other : input) {
			sum += abs(num - other);
		}
		fuels[num] = sum;
	}

	int smallest_key = input[0];
	for (pair<int,size_t> pair : fuels) {
		if (pair.second < fuels[smallest_key]) {
			smallest_key = pair.first;
		}
	}

	return fuels[smallest_key];
}

size_t increasing_sum(size_t n) {
	if (n <= 0) {
		return 0;
	}
	return n + increasing_sum(n-1);
}

size_t part2(vector<int> input) {
	int max = *max_element(input.begin(), input.end());
	int min = *min_element(input.begin(), input.end());

	vector<size_t> fuels(max-min+1, min);

	for (int i = min; i <= max; i++) {
		size_t sum = 0;
		for (int other : input) {
			sum += increasing_sum(abs(i - other));
		}
		fuels[i-min] = sum;
	}

	int smallest_key = 0;
	for (int i = 0; i < fuels.size(); i++) {
		if (fuels[i] < fuels[smallest_key]) {
			smallest_key = i;
		}
	}

	return fuels[smallest_key];
}

int main(void) {
	ifstream input("input.txt");
	string line;
	input >> line;

	vector<int> crabs;
	stringstream stream(line);
	for (string crab; getline(stream, crab, ','); ) {
		crabs.push_back(atoi(crab.c_str()));
	}

	cout << part1(crabs) << endl;
	cout << part2(crabs) << endl;
}
