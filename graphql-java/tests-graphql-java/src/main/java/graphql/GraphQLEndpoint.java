package graphql;

import com.coxautodev.graphql.tools.SchemaParser;
import javax.servlet.annotation.WebServlet;

import com.mongodb.*;
import com.mongodb.client.MongoDatabase;
import graphql.persistance.ManagerPersistance;
import graphql.schema.GraphQLSchema;
import graphql.servlet.SimpleGraphQLServlet;

import java.util.ArrayList;
import java.util.List;

@WebServlet(urlPatterns = "/graphql")
public class GraphQLEndpoint extends SimpleGraphQLServlet {

	private static final ManagerPersistance managerPersistance;

	static {
		//Change to `new MongoClient("mongodb://<host>:<port>/hackernews")`
		//if you don't have Mongo running locally on port 27017
		MongoDatabase mongo = new MongoClient().getDatabase("tests-graphql");
		

		managerPersistance = new ManagerPersistance(mongo.getCollection("MiniSite"));
	}

	public GraphQLEndpoint() {
		super(buildSchema());
	}

	private static GraphQLSchema buildSchema() {

		return SchemaParser.newParser()
				.file("schema.graphqls")
				.resolvers(new Query(managerPersistance))
				.build()
				.makeExecutableSchema();
	}
}