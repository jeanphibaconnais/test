package java8.api;

public interface TestInterface {

	public default void testToString() {
		System.out.println("testToString TestInterface");
	}
}
