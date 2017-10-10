package java8;

import java.util.Arrays;
import java.util.Collections;
import java.util.List;

public class Lambda {

	public static void main(String[] args) {

		testerComparatorList();
		
		testerThrad();
	}

	private static void testerThrad() {
		Thread monThread = new Thread(() -> { System.out.println("Mon traitement"); });
		monThread.start();
	}

	public static void testerComparatorList() {
		List<String> strings = Arrays.asList("hello", "world", "!");
		Collections.sort(strings, (s1, s2)-> s1.compareTo(s2));
		System.out.println(strings);
	}
}
