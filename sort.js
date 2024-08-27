function bubbleSort(arr) {
  let n = arr.length;
  // Traverse through all array elements
  for (let i = 0; i < n - 1; i++) {
    // Last i elements are already sorted
    for (let j = 0; j < n - i - 1; j++) {
      // Swap if the element found is greater than the next element
      if (arr[j] > arr[j + 1]) {
        let temp = arr[j];
        arr[j] = arr[j + 1];
        arr[j + 1] = temp;
      }
    }
  }
}

// Function to print the array
function printArray(arr) {
  console.log(arr.join(" "));
}

// Example usage
let arr = [64, 34, 25, 12, 22, 11, 90];

console.log("Original array:");
printArray(arr);

bubbleSort(arr);

console.log("Sorted array:");
printArray(arr);
