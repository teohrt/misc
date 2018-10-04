package ds;
public class Stack {
	
	private int maxSize;
	private long[] stackArray;
	private int top;
	
	public Stack(int size) {
		this.maxSize = size;
		this.stackArray = new long[maxSize];
		this.top = -1;
	}
	
	public void push(long j) {
		if (isFull()) {
			System.out.println(" This stack is already full.");
		}
		else {
			top++;
			stackArray[top] = j;
		}
	}
	
	public long pop() {
		if (isEmpty()) {
			System.out.println(" This stack is already empty");
			return -1;
		}
		else {
			top--;
			return stackArray[top + 1];
		}
	}
	
	public long peek() {
		return stackArray[top];
	}
	
	public boolean isEmpty() {
		return (top == -1);
	}
	
	public boolean isFull() {
		return (maxSize-1 == top);
	}
	
}
