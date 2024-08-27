def bubble_sort(arr):
    n = len(arr)
    # Traverse through all array elements
    for i in range(n - 1):
        # Last i elements are already sorted
        for j in range(n - i - 1):
            # Swap if the element found is greater than the next element
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]

# Function to print the array
def print_array(arr):
    for i in arr:
        print(i, end=" ")
    print()

# Example usage
arr = [64, 34, 25, 12, 22, 11, 90]

print("Original array:")
print_array(arr)

bubble_sort(arr)

print("Sorted array:")
print_array(arr)
