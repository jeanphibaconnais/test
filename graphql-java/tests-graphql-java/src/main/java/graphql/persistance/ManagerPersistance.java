package graphql.persistance;

import java.util.ArrayList;
import java.util.List;

import com.mongodb.client.MongoCollection;
import graphql.es.MinisiteES;
import org.bson.Document;

public class ManagerPersistance {

	private final MongoCollection<Document> minisitesDocument;

	public ManagerPersistance(MongoCollection<Document> minisites) {
		this.minisitesDocument = minisites;
	}

	public List<MinisiteES> getAllMinisite() {
		List<MinisiteES> minisites = new ArrayList<>();
		for (Document doc : minisitesDocument.find().limit(10)) {
			minisites.add(op2ES(doc));
		}
		return minisites;
	}

	private MinisiteES op2ES(Document doc) {
		return new MinisiteES(
				doc.get("_id").toString(),
				doc.getString("idMiniSite"),
				doc.getString("statutMinisite")
		);
	}

}
