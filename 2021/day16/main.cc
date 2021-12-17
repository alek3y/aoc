#include <iostream>
#include <cstdint>
#include <fstream>
#include <vector>
#include <cassert>
using namespace std;

int64_t pack(size_t from, size_t to, vector<uint8_t> &sequence) {
	assert(to - from <= sizeof(long)*8);

	int64_t result = 0;
	for (size_t i = from; i < to; i++) {
		result |= sequence[i] << (to-i - 1);
	}
	return result;
}

struct Packet {
	int version, type;
	vector<uint8_t> data;
	vector<Packet> sub;

	size_t length;

	static Packet parse(size_t from, size_t to, vector<uint8_t> &sequence) {
		Packet parsed;

		parsed.version = pack(from, from+3, sequence);
		parsed.type = pack(from+3, from+6, sequence);

		if (parsed.type == 4) {
			int is_last = 0;
			size_t last = 0;
			for (size_t i = from+6; i < to && !is_last; i += 5) {
				is_last = sequence[i] == 0;
				last = i;

				for (size_t b = 1; b <= 4 && i+b < to; b++) {
					parsed.data.push_back(sequence[i+b]);
				}
			}
			parsed.length = last+5 - from;
		} else {
			int length_type = pack(from+6, from+7, sequence);

			if (length_type == 0) {
				size_t sub_bits_len = pack(from+7, from+7+15, sequence);
				size_t sub_end = sub_bits_len + from+7+15;

				for (size_t sub_index = from+7+15; sub_index < sub_end; ) {
					Packet current = Packet::parse(sub_index, sub_end, sequence);

					parsed.sub.push_back(current);
					sub_index += current.length;
				}

				parsed.length = sub_end - from;
			} else if (length_type == 1) {
				size_t sub_count = pack(from+7, from+7+11, sequence);

				size_t sub_index = from+7+11;
				for (size_t i = 0; i < sub_count && sub_index < to; i++) {
					Packet current = Packet::parse(sub_index, to, sequence);

					parsed.sub.push_back(current);
					sub_index += current.length;
				}

				parsed.length = sub_index - from;
			}
		}

		return parsed;
	}

	size_t value() {
		switch (this->type) {
			case 4:
				return pack(0, this->data.size(), this->data);
			case 0:
				{
					size_t sum = 0;
					for (Packet &sub : this->sub) {
						sum += sub.value();
					}
					return sum;
				}
			case 1:
				{
					size_t mul = 1;
					for (Packet &sub : this->sub) {
						size_t val = sub.value();
						mul *= val;
					}
					return mul;
				}
			case 2:
				{
					size_t min = this->sub[0].value();
					for (size_t i = 1; i < this->sub.size(); i++) {
						size_t value = this->sub[i].value();
						if (value < min) {
							min = value;
						}
					}
					return min;
				}
			case 3:
				{
					size_t max = 0;
					for (Packet &sub : this->sub) {
						size_t value = sub.value();
						if (value > max) {
							max = value;
						}
					}

					return max;
				}
			case 5:
				{
					return this->sub[0].value() > this->sub[1].value();
				}
			case 6:
				{
					return this->sub[0].value() < this->sub[1].value();
				}
			case 7:
				{
					return this->sub[0].value() == this->sub[1].value();
				}
		}

		return 0;
	}
};

size_t sum_versions(Packet &packet) {
	size_t sum = packet.version;
	for (Packet &sub : packet.sub) {
		sum += sum_versions(sub);
	}
	return sum;
}

size_t part1(vector<uint8_t> &sequence) {
	Packet parsed = Packet::parse(0, sequence.size(), sequence);
	return sum_versions(parsed);
}

size_t part2(vector<uint8_t> &sequence) {
	Packet parsed = Packet::parse(0, sequence.size(), sequence);
	return parsed.value();
}

int main(void) {
	ifstream file("input.txt");
	string line;
	file >> line;

	vector<uint8_t> sequence;

	for (char letter : line) {
		int value = letter;
		if (letter >= 'A') {
			value = value - 'A' + 10;
		} else {
			value -= '0';
		}

		for (int i = 3; i >= 0; i--) {
			sequence.push_back((value >> i) & 1);
		}
	}

	cout << part1(sequence) << endl;
	cout << part2(sequence) << endl;
}
