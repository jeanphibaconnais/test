package java8;

import java.util.ArrayList;
import java.util.List;
import java.util.OptionalDouble;

import exemples.Personne;

public class Stream {

	public static void main(String[] args) {

		testAverageList();

		testFindOne();
	}

	public static void testAverageList() {
		List<Personne> personnes = new ArrayList<Personne>();

		personnes.add(new Personne("P1", 19));
		personnes.add(new Personne("P2", 25));
		personnes.add(new Personne("P3", 30));

		OptionalDouble moyenne = personnes.stream().filter(personne -> personne.getAge() > 20)
				.mapToInt(Personne::getAge).average();

		System.out.println(moyenne.getAsDouble());
	}

	public static void testFindOne() {
		List<Personne> personnes = new ArrayList<Personne>();

		personnes.add(new Personne("P1", 19));
		personnes.add(new Personne("P2", 25));
		personnes.add(new Personne("P3", 30));

		String nom = personnes.stream().filter(personne -> personne.getAge() > 20).findFirst().map(Personne::getNom)
				.orElse("Aucun r�sultat");

		System.out.println("Test + de 20 ans : " + nom);
		
		nom = personnes.stream().filter(personne -> personne.getAge() > 30).findFirst().map(Personne::getNom)
				.orElse("Aucun r�sultat");

		System.out.println("Test + de 30 ans : " + nom);
	}

}
