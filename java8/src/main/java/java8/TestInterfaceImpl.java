package java8;

import java8.impl.TestImplement1;
import java8.impl.TestImplement2;

public class TestInterfaceImpl {

	public static void main(String[] args) {
		TestImplement1 t1 = new TestImplement1();
		TestImplement2 t2 = new TestImplement2();
		
		System.out.println("----- Test interface ----");
		t1.testToString();
		t2.testToString();
	}

}
