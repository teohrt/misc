package ds;
public class App {

	public static void main(String[] args) {
		Stack s = new Stack(4);
		
		s.push(10);
		s.push(20);
		s.push(30);
		s.push(40);
		s.push(50);
		
		while(!s.isEmpty()) {
			long val = s.pop();
			System.out.println(val);
		}

	}

}
