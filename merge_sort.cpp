#include <iostream>
using namespace std;

// To execute C++, please define "int main()"

void merge(int arr[], int ref[], int low, int mid, int high) {
  // copy elements from target array into ref array
  for(int i = low; i <= high; i++) {
    ref[i] = arr[i];
  }
  
  /*  Example..
    subarray of array (low -> hi) = [ 1 3 7 4 8 12 ]
    subarray of helper (low -> hi) = [ 1 3 7 4 8 12 ]

    eventually array will be [ 1 3 4 7 8 12 ]
    and somewhere in the middle of the iterations will be

    helper : [ 1 3 7 4 8 12 ]
    array : [ 1 3 4 4 8 12 ]
  */
  
  int ref_left = low;
  int ref_right = mid + 1;
  int current = low;
  
  while (ref_left <= mid && ref_right <= high) {
    if (ref[ref_left] <= ref[ref_right]) {
      arr[current] = ref[ref_left];
      ref_left ++;
    } else {
      arr[current] = ref[ref_right];
      ref_right ++;
    }
    current ++;
  }
  
  /* the right and left side of the subarray are already there
  so if we break out of the loop all we need to do is copy over the rest of the left side of the array.
  if there is any of the left...left 
  */ 
  
  int left_rem = mid - ref_left;
  for (int i = 0; i <= left_rem; i++) {
    arr[current + i] = ref[ref_left + i];
  }
  
}

void merge_sort(int arr[], int ref[], int low, int high) {
  if (low < high) {
    int mid = (low + high) / 2;
    merge_sort(arr, ref, low, mid);
    merge_sort(arr, ref, mid + 1, high);
    merge(arr, ref, low, mid, high);
  }
}

void merge_sort(int arr[], int arr_length) {
  int low = 0;
  int high = arr_length - 1;
  int helper[arr_length];
  
  merge_sort(arr, helper, low, high);
}
                     
int parse_array() {
  int l;
  cin >> l;
  int arr[] = new int[l];
  for(int i = 0; i < l; i++) {
    cin >> arr[i];
  }
  return arr;
}

int main() {
  auto array = parse_array();
  int arr[] = { 1, 7, 2, 8, 9, 6, 32, 85, 23, 19, 24 };
  int l = sizeof(arr) / sizeof(arr[0]);
  cout << l << "\n";
  merge_sort(arr, l);
  for(int i = 0; i < l; i++) {
    cout << arr[i] << " ";
  }
}
