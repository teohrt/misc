package Sorting;

import java.util.Arrays;

public class mergeSort {

	public static void main(String[] args) {
		int[] arr = {1, 6, 3, 9, 6, 7, 2};
		
		System.out.println(Arrays.toString(arr));
		System.out.println(Arrays.toString(mergeSort(arr)));
	}
	
	public static int[] mergeSort(int[] arr) {
		if (arr.length == 1) {
			return arr;
		}
		int mid = arr.length/2;
		int[] left = Arrays.copyOfRange(arr, 0, mid);
		int[] right = Arrays.copyOfRange(arr, mid, arr.length);
		
		return merge(mergeSort(left), mergeSort(right));
	}
	
	public static int[] merge(int[] left, int[] right) {
		int[] result = new int[left.length + right.length];
		int leftIndex = 0, rightIndex = 0, resultIndex = 0;
		
		while (leftIndex < left.length && rightIndex < right.length) {
			if (left[leftIndex] <= right[rightIndex]) {
				result[resultIndex++] = left[leftIndex++];
			}
			else {
				result[resultIndex++] = right[rightIndex++];
			}
		}
		
		// add remaining elements to the results array
		if (leftIndex < left.length) {
			for (int i = leftIndex; i < left.length; i++) {
				result[resultIndex++] = left[i];
			}
		}
		else {
			for (int i = rightIndex; i < right.length; i++) {
				result[resultIndex++] = right[i];
			}
		}
		
		return result;
	}
}
