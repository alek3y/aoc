#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <alloca.h>

size_t fsize(FILE *file) {
	size_t pos = ftell(file);
	fseek(file, 0, SEEK_END);
	size_t size = ftell(file);

	fseek(file, pos, SEEK_SET);
	return size;
}

typedef struct {
	int *items;
	size_t width, height;
} Map;

int *map_at(size_t y, size_t x, Map map) {
	return &map.items[map.width*y + x];
}

void map_print(Map map) {
	for (size_t y = 0; y < map.height; y++) {
		for (size_t x = 0; x < map.width; x++) {
			printf("%c", *map_at(y, x, map) + '0');
		}
		printf("\n");
	}
}

size_t flash(size_t y, size_t x, Map map) {
	size_t flashes = 1;

	*map_at(y, x, map) = 0;
	for (ssize_t rel_y = -1; rel_y < 2; rel_y++) {
		for (ssize_t rel_x = -1; rel_x < 2; rel_x++) {
			ssize_t abs_y = y + rel_y, abs_x = x + rel_x;

			if (abs_x < 0 || abs_y < 0) {
				continue;
			} else if (abs_x >= map.width || abs_y >= map.height) {
				continue;
			}

			if (*map_at(abs_y, abs_x, map) > 0) {
				(*map_at(abs_y, abs_x, map))++;
			}

			if (*map_at(abs_y, abs_x, map) > 9) {
				flashes += flash(abs_y, abs_x, map);
			}
		}
	}

	return flashes;
}

size_t part1(Map map) {
	size_t flashes = 0;

	for (size_t s = 0; s < 100; s++) {
		for (size_t y = 0; y < map.height; y++) {
			for (size_t x = 0; x < map.width; x++) {
				(*map_at(y, x, map))++;
			}
		}

		for (size_t y = 0; y < map.height; y++) {
			for (size_t x = 0; x < map.width; x++) {
				if (*map_at(y, x, map) > 9) {
					flashes += flash(y, x, map);
				}
			}
		}
	}

	return flashes;
}

size_t part2(Map map) {
	size_t steps = 0;

	while (1) {
		for (size_t y = 0; y < map.height; y++) {
			for (size_t x = 0; x < map.width; x++) {
				(*map_at(y, x, map))++;
			}
		}

		size_t flashes = 0;
		for (size_t y = 0; y < map.height; y++) {
			for (size_t x = 0; x < map.width; x++) {
				if (*map_at(y, x, map) > 9) {
					flashes += flash(y, x, map);
				}
			}
		}

		if (flashes == map.height * map.width) {
			break;
		}
		steps++;
	}

	return steps+1;
}

int main(void) {
	FILE *file = fopen("input.txt", "r");

	Map map = {
		.items = NULL,
		.width = 0,
		.height = 0
	};

	int *numbers = malloc(fsize(file)*sizeof(int));

	char *line = NULL;
	while (getline(&line, alloca(sizeof(size_t)), file) != -1) {
		if (map.width == 0) {
			map.width = strlen(line)-1;
		}

		for (size_t i = 0; i < strlen(line)-1; i++) {
			numbers[map.height*map.width + i] = line[i] - '0';
		}
		map.height++;

		free(line);
		line = NULL;
	}

	size_t numbers_size = map.height * map.width * sizeof(int);
	map.items = malloc(numbers_size);

	memcpy(map.items, numbers, numbers_size);
	printf("%lu\n", part1(map));

	memcpy(map.items, numbers, numbers_size);
	printf("%lu\n", part2(map));

	free(numbers);
	free(map.items);
	fclose(file);
}
