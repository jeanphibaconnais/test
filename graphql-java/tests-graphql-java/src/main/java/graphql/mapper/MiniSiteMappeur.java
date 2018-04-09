package graphql.mapper;

import graphql.es.MinisiteES;
import org.bson.Document;

public class MiniSiteMappeur {

	public MinisiteES op2ES(Document doc) {
		MinisiteES es = null;

		if (null != doc) {
			es = new MinisiteES(
					doc.get("_id").toString(),
					doc.getString("idMiniSite"),
					doc.getString("statutMinisite")
			);
		}

		return es;
	}

}
