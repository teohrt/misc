package ds;

public class App {

	public static void main(String[] args) {
		DoublyLinkedList list = new DoublyLinkedList();
		
		list.insertLast(3);
		list.insertLast(4);
		list.insertLast(9999);
		list.insertFirst(2);
		list.insertFirst(1);
		list.deleteLast();
		list.insertAfter(4, 5);
		list.insertFirst(0);
		list.deleteFirst();
		list.insertFirst(00000);
		list.deleteFirst();
		list.insertAfter(3, 6);
		list.deleteKey(6);
		
		list.displayForward();
		list.displayBackward();
	}
}
