import sys
file = open("input.txt", "r").read().replace("\n", "")

area = [
	[int(n) for n in i.replace(",", "").split("=")[1].split("..")]
	for i in file.split(" ")[-2:]
]

class Point:
	def __init__(self, x, y):
		self.x = x
		self.y = y

class Area:
	def __init__(self, left_right, bottom_up):
		self.begin = Point(left_right[0], bottom_up[0])
		self.width = left_right[1] - left_right[0] + 1
		self.height = bottom_up[1] - bottom_up[0] + 1

	def within(self, point):
		return (
			(point.x >= self.begin.x and point.x < self.begin.x + self.width) and
			(point.y >= self.begin.y and point.y < self.begin.y + self.height)
		)

class Probe:
	def __init__(self, velocity):
		self.position = Point(0, 0)
		self.velocity = Point(velocity[0], velocity[1])

	def step(self):
		self.position.x += self.velocity.x
		self.position.y += self.velocity.y

		if self.velocity.x > 0:
			self.velocity.x -= 1
		elif self.velocity.x < 0:
			self.velocity.x += 1

		self.velocity.y -= 1

def try_velocities(region):
	max_y = 0
	velocities_count = 0
	for y in range(-abs(region.begin.y), abs(region.begin.y)):
		for x in range(0, region.begin.x + region.width):
			probe = Probe((x, y))

			max_y_for_probe = 0
			while (
				probe.position.x < region.begin.x + region.width and
				probe.position.y >= region.begin.y
			):
				if region.within(probe.position):
					if max_y_for_probe > max_y:
						max_y = max_y_for_probe
					velocities_count += 1
					break

				if probe.position.y > max_y_for_probe:
					max_y_for_probe = probe.position.y

				probe.step()

	return (max_y, velocities_count)

def part1(region):
	return try_velocities(region)[0]

def part2(region):
	return try_velocities(region)[1]

region = Area(area[0], area[1])
print(part1(region))
print(part2(region))
