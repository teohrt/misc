package ds;
public class App {

	public static void main(String[] args) {
		Queue q = new Queue(5);
		q.insert(100);
		q.insert(200);
		q.insert(300);
		q.insert(400);
		q.insert(500);
		q.insert(600);
		q.view();
		q.remove();
		q.view();
	}

}
