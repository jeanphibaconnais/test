package java8;

import java.time.Duration;
import java.time.Instant;
import java.time.LocalDateTime;
import java.time.Month;
import java.time.ZoneId;
import java.time.ZonedDateTime;
import java.time.format.DateTimeFormatter;
import java.time.temporal.ChronoUnit;

public class DateJSR310 {

	public static void main(String[] args) {
		try {
			Instant now = Instant.now();
			Thread.sleep(1000);
			Instant reNow = Instant.now();

			Duration elapsed = Duration.between(now, reNow);

			System.out.println(elapsed.get(ChronoUnit.SECONDS));

			Duration added = elapsed.plus(2L, ChronoUnit.SECONDS);

			System.out.println(added.get(ChronoUnit.SECONDS));

			testLocalTime();

		} catch (InterruptedException e) {
			e.printStackTrace();
		}

	}

	public static void testLocalTime() {
		DateTimeFormatter format = DateTimeFormatter.ofPattern("HH'h'mm, dd MMM yyyy");

		LocalDateTime ldt = LocalDateTime.of(2017, Month.OCTOBER, 19, 14, 30);
		System.out.println("LocalDateTime : " + format.format(ldt));

		ZonedDateTime parisDateTime = ldt.atZone(ZoneId.of("Europe/Paris"));
		System.out.println("Depart : " + format.format(parisDateTime));

		
		ZonedDateTime japanDateTime = parisDateTime.withZoneSameInstant(ZoneId.of("Asia/Tokyo"));
		System.out.println("Arrive : " + format.format(japanDateTime));

	}

}
