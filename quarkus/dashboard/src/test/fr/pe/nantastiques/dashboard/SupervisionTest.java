package org.acme.quickstart;

import fr.pe.nantastiques.config.Config;
import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static org.hamcrest.CoreMatchers.is;

@QuarkusTest
public class SupervisionTest {

	@Test
	public void testgetVersion() {
		given()
				.when().get("/supervision")
				.then()
				.statusCode(200)
				.body(is(Config.VERSION));

	}

}