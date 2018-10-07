package ds;

public class App {

	public static void main(String[] args) {
		SinglyLinkedList list = new SinglyLinkedList();
		
		list.insertFirst(100);
		list.insertFirst(200);
		list.insertFirst(300);
		list.insertFirst(400);
		list.insertFirst(500);
		list.insertLast(1000000);
		
		list.displayList();
		
		
	}
}
