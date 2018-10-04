package ds;
public class App {

	public static void main(String[] args) {
		String input = "Reverse this string by using a stack!";
		
		Stack s = new Stack(input.length());
		
		for (int i = 0; i < input.length(); i++) {
			s.push(input.charAt(i));
		}
		
		while(!s.isEmpty()) {
			char val = s.pop();
			System.out.print(val);
		}
	}

}
