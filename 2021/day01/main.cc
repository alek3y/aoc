#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int part1(vector<int> &input) {
	int count = 0, previous = 0;

	for (int number : input) {
		if (number > previous) {
			count++;
		}
		previous = number;
	}

	return count - 1;
}

int part2(vector<int> &input) {
	vector<int> windows;
	for (size_t i = 0; i < input.size() - 2; i++) {
		windows.push_back(input[i] + input[i+1] + input[i+2]);
	}

	return part1(windows);
}

int main(void) {
	ifstream file("input.txt");

	vector<int> input;
	for (string line; file >> line; ) {
		input.push_back(stoi(line));
	}

	cout << part1(input) << endl;
	cout << part2(input) << endl;
}
