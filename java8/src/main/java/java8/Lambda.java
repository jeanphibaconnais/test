package java8;

import exemples.Personne;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;
import java.util.function.BiFunction;
import java.util.function.Consumer;
import java.util.function.Predicate;

public class Lambda {

	// Functional Interfaces :

	// Un Predicate prend un argument et retourne un booléen.
	// Un Consumer prend un argument mais ne retourne pas de résultat.
	// Les Comparator sont devenus des FunctionalInterface.
	// 		Comparator<Integer> ascending = (a, b) -> a.compareTo(b);
	// Un Supplier ne prend pas d’argument et produit un résultat.
	// BiFunction : prend deux arguments et retourne un résultat.
	// 		BiFunction<Integer, String, String> concat = (Integer i, String s) -> s + ": " + i;

	public static void main(String[] args) {

		testerComparatorList();

		testerThrad();

		testPredicate();

		testBiFunction();
	}

	private static void testerThrad() {
		System.out.println("\n=================== testerThrad ");
		Thread monThread = new Thread(() -> {
			System.out.println("Mon traitement");
		});
		monThread.start();
	}

	public static void testerComparatorList() {
		System.out.println("\n=================== testerComparatorList ");
		List<String> strings = Arrays.asList("hello", "world", "!");
		Collections.sort(strings, (s1, s2) -> s1.compareTo(s2));
		System.out.println(strings);
	}

	public static void testPredicate() {
		System.out.println("\n=================== testPredicate ");
		Personne p1 = new Personne("Thomas", 19);
		Personne p2 = new Personne("Paul", 25);

		List<Personne> personnes = new ArrayList<>();

		personnes.add(p1);
		personnes.add(p2);

		Predicate<Personne> plusVieuxQue = (personne) -> {
			return personne.getAge() >= 20;
		};
		
		Predicate<Personne> plusJeuneQue = (personne) -> {
			return personne.getAge() < 20;
		};

		Consumer<Personne> afficherPersonne = (personne) -> {
			personne.print();
		};

		printPersonnesWithPredicate(personnes, plusVieuxQue, afficherPersonne);
		printPersonnesWithPredicate(personnes, plusJeuneQue, afficherPersonne);
		
	}
	
	public static void printPersonnesWithPredicate(List<Personne> personnes, Predicate<Personne> tester, Consumer<Personne> consumer) {
	    for (Personne p : personnes) {
	        if (tester.test(p)) {
	        	consumer.accept(p);
	        }
	    }
	}

	public static void testBiFunction() {
		System.out.println("\n=================== testBiFunction ");
		BiFunction<Integer, String, String> concat = (Integer i, String s) -> s + ": " + i;

		System.out.println(concat.apply(10, "20"));
	}
}
