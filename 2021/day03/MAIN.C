/*
Language: C
Compiler: OpenWatcom C/C++ (16-bit)
Environment: FreeDOS
*/

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Diags {
	size_t gamma, epsilon;
};

struct Rating {
	size_t oxygen, co2;
};

size_t fsize(FILE *file) {
	size_t pos, size;

	pos = ftell(file);
	fseek(file, 0, SEEK_END);
	size = ftell(file);

	fseek(file, pos, SEEK_SET);
	return size;
}

short most_common(size_t bit, char **input, size_t length) {
	size_t ones = 0, i;
	for (i = 0; i < length; i++) {
		ones += input[i][bit] == '1';
	}
	return ones > length/2;
}

struct Diags part1(char **input, size_t length) {
	struct Diags diags = {0};
	size_t mask = 0;

	size_t i;
	size_t width = strlen(input[0]);
	for (i = 0; i < width; i++) {
		diags.gamma += most_common(i, input, length) << (width - i-1);
	}

	mask = 0;
	for (i = 0; i < width; i++) {
		mask += 1 << i;
	}

	diags.epsilon = ~diags.gamma & mask;
	return diags;
}

short most_common_filtered(size_t bit, char **input, size_t *indexes, size_t length) {
	size_t half_length = length/2;
	size_t ones = 0, i;
	for(i = 0; i < length; i++) {
		ones += input[indexes[i]][bit] == '1';
	}

	if (ones == half_length && (length % 2 == 0)) {
		return -1;
	}
	return ones > half_length;
}

size_t rating_compute(short common_equals, short most_common, char **input, size_t *indexes, size_t length) {
	size_t width = strlen(input[0]);
	size_t rating, replace_index;
	size_t i, bit;
	short common;

	for (bit = 0; bit < width && length > 1; bit++) {
		common = most_common_filtered(bit, input, indexes, length);
		if (common == -1) {
			common = common_equals;
		} else if (!most_common) {
			common = !common;
		}

		replace_index = 0;
		for (i = 0; i < length; i++) {
			if (input[indexes[i]][bit] - 0x30 == common) {
				indexes[replace_index++] = indexes[i];
			}
		}
		length = replace_index;
	}

	rating = 0;
	for (i = 0; i < width; i++) {
		rating += (input[indexes[0]][i] - 0x30) << (width - i-1);
	}

	return rating;
}

struct Rating part2(char **input, size_t length) {
	struct Rating rating = {0};
	size_t *indexes = malloc(length * sizeof(*indexes));
	short common_equal, common;
	size_t i, j;

	for (i = 0; i < length; i++) {
		indexes[i] = i;
	}
	rating.oxygen = rating_compute(1, 1, input, indexes, length);
	
	for (i = 0; i < length; i++) {
		indexes[i] = i;
	}
	rating.co2 = rating_compute(0, 0, input, indexes, length);

	return rating;
}

int main(void) {
	FILE *file;
	size_t file_size, newlines, line_size;
	char *contents, **input;
	struct Diags diags;
	struct Rating rating;

	size_t i;
	char *j;

	file = fopen("input.txt", "r");
	file_size = fsize(file);
	contents = malloc(file_size + 1);

	fread(contents, file_size, 1, file);
	contents[file_size] = '\0';

	for (
		j = contents, newlines = 0;
		j = strchr(j, '\n');
		j++, newlines++
	);

	*strchr(contents, '\n') = '\0';
	line_size = strlen(contents);
	free(contents);
	fseek(file, 0, SEEK_SET);

	input = malloc(newlines * sizeof(contents));
	for (i = 0; i < newlines; i++) {
		input[i] = malloc(line_size + 1);
		fscanf(file, "%s", input[i]);
	}

	diags = part1(input, newlines);
	printf("%u*%u\n", diags.gamma, diags.epsilon);

	rating = part2(input, newlines);
	printf("%u*%u\n", rating.oxygen, rating.co2);

	return 0;
}
