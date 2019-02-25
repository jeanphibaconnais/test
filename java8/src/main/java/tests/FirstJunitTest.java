package tests;

import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.CsvSource;

import java.util.function.BiFunction;

import static org.junit.jupiter.api.Assertions.assertEquals;

public class FirstJunitTest {

	@Test
	@DisplayName("First test junit 5 : 1 + 1 = 2")
	public void testOne() {
		assertEquals(2, 1 + 1);
	}

	@ParameterizedTest(name = "{0} + {1} = {2}")
	@CsvSource({
			"0,    1,   1",
			"1,    2,   3",
			"49,  51, 100",
			"1,  100, 101"
	})
	void add(int first, int second, int expectedResult) {
		BiFunction<Integer, Integer, Integer> somme = (x, y) -> x + y;

		assertEquals(expectedResult, (int) somme.apply(first, second),
				() -> first + " + " + second + " should equal " + expectedResult);

	}
}
