package graphql;

import com.coxautodev.graphql.tools.GraphQLRootResolver;
import graphql.es.MinisiteES;
import graphql.persistance.ManagerPersistance;

import java.util.List;

class Query implements GraphQLRootResolver {

	private final ManagerPersistance manager;

	public Query(ManagerPersistance manager) {
		this.manager = manager;
	}

	public List<MinisiteES> allMinisites() {
		return manager.getAllMinisite();
	}

	public MinisiteES getMinisite(final String id) {
		return manager.lireMinisite(id);
	}
}
