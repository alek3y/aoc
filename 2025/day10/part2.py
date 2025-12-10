from sympy import *

# # Example
# A = [[3], [1, 3], [2], [2, 3], [0, 2], [0, 1]]
# b = [3, 5, 4, 7]
# Solution: [1, 3, 0, 3, 1, 2]
#
# ## Before reduced row echelon form
# [0 0 0 0 1 1] [3]
# [0 1 0 0 0 1] [5]
# [0 0 1 1 1 0] [4]
# [1 1 0 1 0 0] [7]
# -------------
# [1 3 0 3 1 2]
#
# ## After reduced row echelon form
# [1 0 0 1 0 -1] [2]
# [0 1 0 0 0  1] [5]
# [0 0 1 1 0 -1] [1]
# [0 0 0 0 1  1] [3]
# --------------
# [1 3 0 3 1  2]
def least_presses(buttons, target):
	A = []	# A of the system of equations Ax = b
	for counter in range(len(target)):
		A.append([])
		for button in buttons:
			counter_depends_on_this = counter in button
			A[-1].append(1 if counter_depends_on_this else 0)

	A = Matrix([
		row + [target[i]] for (i, row) in enumerate(A)
	])

	A, pivots = A.rref()
	A = A.tolist()

	pivots = set(pivots)
	equations = {}
	for (i, pivot) in enumerate(pivots):
		constant = A[i][-1]
		equations[pivot] = [(None, constant)]	# Already originally on the right hand side of =
		for var in range(pivot+1, len(buttons)):
			multiplier = A[i][var]
			if multiplier != 0:
				equations[pivot].append((var, -multiplier))	# Moving the member to the other side of = changes the sign

	# Takes in a set of equations in the form (dependency, factor) and
	# reduces it incrementally into one value; gets stuck if none of
	# the equations can be reduced (i.e. cyclic dependency or missing
	# dependency)... You can see I'm tired by how shit this is :D
	def collapse(equations):
		collapsed = {}
		uncollapsed = [var for var in equations]
		while len(uncollapsed) > 0:
			var = uncollapsed.pop(0)
			equation = equations[var]

			for i in range(len(equation)):
				(dependency, multiplier) = equation[i]
				if dependency != None and dependency in collapsed:
					equation = equation[:i] + [(None, collapsed[dependency]*multiplier)] + equation[i+1:]

			merged_constant = 0
			merged_equation = []
			for (dependency, multiplier) in equation:
				if dependency == None:
					merged_constant += multiplier
				else:
					merged_equation.append((dependency, multiplier))

			equations[var] = [(None, merged_constant)] + merged_equation

			if len(merged_equation) == 0:
				if merged_constant < 0 or merged_constant//1 != merged_constant:	# A button can only be pressed a positive integer amount of times
					return None
				collapsed[var] = merged_constant	# Collapse if it's only made of a constant
			else:
				uncollapsed.append(var)
		return collapsed

	free = set(range(len(buttons))).difference(pivots)	# Anything that's not a pivot is a free variable that can be assigned anything
	if len(free) == 0:	# Then there is one solution
		return sum(collapse(equations).values())
	else:
		max_free_value = {}
		for var in free:
			button = buttons[var]
			max_free_value[var] = min([target[counter] for counter in button])

		order = list(free)
		chosen_values = {var: 0 for var in order}

		min_found = None
		reached_end = False
		while not reached_end:	# Try all values in the range of the free variables
			equations_and_chosen = equations.copy()
			equations_and_chosen.update({var: [(None, chosen_value)] for (var, chosen_value) in chosen_values.items()})

			collapsed = collapse(equations_and_chosen)
			if collapsed != None:
				current_sum = sum(collapsed.values())
				if min_found == None or current_sum < min_found:
					min_found = current_sum

			for i in range(0, len(order))[::-1]:
				if chosen_values[order[i]] < max_free_value[order[i]]:
					chosen_values[order[i]] += 1
					break
				elif i > 0:
					chosen_values[order[i]] = 0
				else:
					reached_end = True
					break
		return min_found

machines = []
with open("input.txt", "r") as file:
	for line in file.read().strip().split("\n"):
		buttons = []
		for button in line.split(" ")[1:-1]:
			buttons.append([int(number) for number in button[1:-1].split(",")])
		joltages = [int(number) for number in line.split(" ")[-1][1:-1].split(",")]
		machines.append((buttons, joltages))

all_min = 0
for (i, (buttons, joltages)) in enumerate(machines):
	print(f" {i}", end="\r")	# Some progress indicator...
	all_min += least_presses(buttons, joltages)
print(all_min)
