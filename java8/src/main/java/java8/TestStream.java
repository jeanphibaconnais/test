package java8;

import java.util.Arrays;
import java.util.List;
import java.util.stream.Stream;

enum Couleur {
	MARRON("marron"),
	BLEU("bleu"),
	VERT("vert"),
	VERRON("verron"),
	INCONNU("non déterminé"),
	ROUGE("rouge mais j'avais piscine...");

	private String name = "";

	Couleur(String n){name = n;}
	public String toString() {return name;}
}

class Personne {

	public Double taille = 0.0d, poids = 0.0d;
	public String nom = "", prenom = "";
	public Couleur yeux = Couleur.INCONNU;
	public Personne() {	}
	public Personne(double taille, double poids, String nom, String prenom, Couleur yeux) {
		super();
		this.taille = taille;
		this.poids = poids;
		this.nom = nom;
		this.prenom = prenom;
		this.yeux = yeux;
	}

	public String toString() {
		String s = "Je m'appelle " + nom + " " + prenom;
		s += ", je pèse " + poids + " Kg";
		s += ", et je mesure " + taille + " cm.";
		return s;
	}
	public Double getTaille() {return taille;}
	public void setTaille(Double taille) {this.taille = taille;}
	public Double getPoids() {return poids;}
	public void setPoids(Double poids) {this.poids = poids;}
	public String getNom() {return nom;}
	public void setNom(String nom) {this.nom = nom;}
	public String getPrenom() {return prenom;}
	public void setPrenom(String prenom) {this.prenom = prenom;}
	public Couleur getYeux() {return yeux;}
	public void setYeux(Couleur yeux) {this.yeux = yeux;}
}

public class TestStream {
	public static void main(String[] args) {
		List<Personne> listP = Arrays.asList(
				new Personne(1.80, 70, "A", "Nicolas", Couleur.BLEU),
				new Personne(1.56, 50, "B", "Nicole", Couleur.VERRON),
				new Personne(1.75, 65, "C", "Germain", Couleur.VERT),
				new Personne(1.68, 50, "D", "Michel", Couleur.ROUGE),
				new Personne(1.96, 65, "E", "Cyrille", Couleur.BLEU),
				new Personne(2.10, 120, "F", "Denis", Couleur.ROUGE),
				new Personne(1.90, 90, "G", "Olivier", Couleur.VERRON)
		);

		Stream<Personne> personnes = listP.stream();
		personnes.forEach(System.out::println);

		System.out.println("");

		personnes = listP.stream();
		personnes.filter(x -> x.getPoids() > 65).forEach(System.out::println);

		System.out.println("");

		personnes = listP.stream();
		personnes.filter(x -> x.getPoids() > 65).map(x -> x.getPrenom()).forEach(System.out::println);


	}
}
