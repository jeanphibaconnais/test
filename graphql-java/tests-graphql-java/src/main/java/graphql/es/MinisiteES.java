package graphql.es;

public class MinisiteES {

	private final String id;
	private final String idMinisite;
	private final String statutMinisite;

	public MinisiteES(String id, String idMinisite, String statutMinisite) {
		this.id = id;
		this.idMinisite = idMinisite;
		this.statutMinisite = statutMinisite;
	}

	public String getId() {
		return id;
	}

	public String getIdMinisite() {
		return idMinisite;
	}

	public String getStatutMinisite() {
		return statutMinisite;
	}
}
