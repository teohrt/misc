package Sorting;

import java.util.Arrays;

public class quickSort {

	public static void main(String[] args) {
		int [] arr = {15, 3, 2, 1, -30, 9999, 9, 5, 7, 8, 6};
		System.out.println(Arrays.toString(arr));
		qs(arr, 0, arr.length-1);
		System.out.println(Arrays.toString(arr));
	}
	
	public static void qs(int[] arr, int left, int right) {
		if (left >= right) return;
		
		int pivot = arr[(right + left) / 2];
		int index = partition(arr, left, right, pivot);
		
		qs(arr, left, index-1);
		qs(arr, index, right);
	}
	
	public static int partition(int[] arr, int left, int right, int pivot) {
		while (left <= right) {
			while(arr[left ] < pivot) left++;
			while(arr[right] > pivot) right--;
			
			// Swap!
			if (left <= right) {
				int temp = arr[left];
				arr[left] = arr[right];
				arr[right] = temp;
				left++;
				right--;
			}
		}
		return left;
	}
}
