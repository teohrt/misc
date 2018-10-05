package ds;

public class Queue {
	private int maxSize;
	private long[] queArray;
	private int front;
	private int rear;
	private int nItems;

	public Queue(int size) {
		this.maxSize = size;
		this.queArray = new long[size];
		front = 0;
		rear = -1;
		nItems = 0;
	}

	public void insert(long v) {
		// This logic makes it a circular queue
		if (rear == maxSize-1) {
			rear = -1;
		}
		
		rear++;
		queArray[rear] = v;
		nItems++;
	}

	public long remove() {
		long temp = queArray[front];
		front++;
		if(front == maxSize) {
			front = 0;
		}
		nItems--;
		return temp;
	}

	public long peekFront() {
		return queArray[front];
	}

	public boolean isFull() {
		return (nItems == 0);
	}

	public void view() {
		System.out.print("[ ");
		for(int i = front; i < queArray.length; i++) {
			System.out.print(queArray[i] + " ");
		}
		System.out.println("]");
	}
}
