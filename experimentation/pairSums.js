/*
Given an array A of size N. 
You need to find all pairs in the array that sum to a number K. 
If no such pair exists then output will be -1. 
The elements of the array are distinct and are in sorted order.
*/

// Brute force
// O(N^2)
const bruteForce = (arr, k) => {

    let paired = false;

    for (let i = 0; i < arr.length; i++) {
        for (let j = i+1; j < arr.length; j++) {
            if (arr[i] + arr[j] === k) {
                console.log(arr[i] + " " + arr[j]);
                paired = true;
            }
        }
    }
    if (!paired) console.log(-1);
}

// O(N log N)
const binarySearchLoop = (arr, k) => {
    let paired = false;

    // Object to store pairs
    let map = {};

    for (let i = 0; i < arr.length; i++) {
        let result = binarySearch(arr, (k - arr[i]));

        // If BinarySearch returns a useable value, if that value is not i and if the pair hasn't been mapped before
        if (result !== false && result !== i && map[arr[i]] == undefined) {
            map[arr[result]] = arr[i];
            console.log(arr[i] + " " + arr[result]);
            paired = true;
        }
    }
    if (!paired) console.log(-1);
}

// Iterative binary search
// Returns index of array with value k or false
// Binary search is O(log N)
const binarySearch = (arr, k) => {
    let low = 0;
    let high = arr.length - 1;
    
    while(high >= low) {
        let mid = Math.ceil((low + high)/2);

        if (arr[mid] === k) {
            return mid;
        }

        // k is in the right half
        else if (arr[mid] < k) {
            low = mid + 1;
        }

        // k is in the left half
        else if (arr[mid] > k){
            high = mid - 1;
        }
    }
    return false;
}


// Function prints pairs of numbers in sorted array 
// that add up to k.
// Runs in O(N) time.
const twoPointer = (arr, k) => {
    
    let p1 = 0;
    let p2 = arr.length-1;
    let paired = false;
    
    // While the pointers haven't crossed
    while(p1 < p2) {
        
        let x1 = arr[p1];
        let x2 = arr[p2];
        
        // Y = K - X
        // Y is the required pair for X to sum up to K
        let y1 = k - x1;
        let y2 = k - x2;
        
        // Pair found!
        if (y1 + y2 == k) {
            console.log(arr[p1] + " " + arr[p2]);
            paired = true;
            p1++;
            p2--;
        }
        else {
            // Neither have a pair
            if (y1 > x2 && y2 < x1) {
                p1++;
                p2--;
            }
            
            // p1 has no pair
            else if (y1 > x2) {
                p1++;
            }
            
            // p2 has no pair
            else if (y2 < x1) {
                p2--;
            }
        }
    }
}

const arr = [-5, 10, 20, 22, 33, 50, 60, 80];
const k = 55;

console.log(arr);
console.log("Pairs that sum to " + k + " are:");

twoPointer(arr, k);
bruteForce(arr, k);
binarySearchLoop(arr, k);