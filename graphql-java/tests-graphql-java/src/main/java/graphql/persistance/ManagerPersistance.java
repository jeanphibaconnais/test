package graphql.persistance;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

import com.mongodb.BasicDBObject;
import com.mongodb.client.FindIterable;
import com.mongodb.client.MongoCollection;
import graphql.es.MinisiteES;
import graphql.mapper.MiniSiteMappeur;
import org.bson.Document;

import static com.mongodb.client.model.Filters.eq;

public class ManagerPersistance {

	private MiniSiteMappeur mappeur;

	private final MongoCollection<Document> minisitesDocument;

	public ManagerPersistance(MongoCollection<Document> minisites) {
		this.minisitesDocument = minisites;
	}

	public List<MinisiteES> getAllMinisite() {
		List<MinisiteES> minisites = new ArrayList<>();
		for (Document doc : minisitesDocument.find().limit(10)) {
			minisites.add(mappeur.op2ES(doc));
		}
		return minisites;
	}


	public MinisiteES lireMinisite(String id) {
		mappeur = new MiniSiteMappeur();
		FindIterable<Document> docs = minisitesDocument.find(eq("_id", Integer.parseInt(id)));
		Iterator iterator = docs.iterator();
		Document result = (Document) iterator.next();

		return mappeur.op2ES(result);
	}
}
