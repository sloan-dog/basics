#!/bin/python3.6

def merge(arr, helper, low, mid, high):
	for i in range(low, high + 1):
		helper[i] = arr[i]

	helper_left = low
	helper_right = mid + 1
	current = low

	while helper_left <= mid and helper_right <= high:
		# print(helper[helper_left], helper[helper_right])
		if helper[helper_left] <= helper[helper_right]:
			arr[current] = helper[helper_left]
			helper_left += 1
		else:
			arr[current] = helper[helper_right]
			helper_right += 1
		current += 1

	left_remaining = mid - helper_left
	for i in range(left_remaining + 1):
		arr[current + i] = helper[helper_left + i]

def merge_sort(arr, helper, low, high):
	if low < high:
		mid = int( (low + high) / 2)
		merge_sort(arr, helper, low, mid)
		merge_sort(arr, helper, mid + 1, high)
		merge(arr, helper, low, mid, high)

def do_merge(arr):
	l = len(arr)
	helper = [0] * l
	low = 0
	high = l - 1

	merge_sort(arr, helper, low, high)
	return arr


if __name__ == "__main__":
	l = input().strip()
	arr = [int(x) for x in input().strip().split(" ")]

	print(arr)
	do_merge(arr)
	print(arr)

