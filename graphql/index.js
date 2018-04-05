const { GraphQLServer } = require('graphql-yoga');
const { rsvp } = require('rsvp');

var Db = require('mongodb').Db,
	MongoClient = require('mongodb').MongoClient,
	Server = require('mongodb').Server,
	ReplSetServers = require('mongodb').ReplSetServers,
	ObjectID = require('mongodb').ObjectID,
	Binary = require('mongodb').Binary,
	GridStore = require('mongodb').GridStore,
	Grid = require('mongodb').Grid,
	Code = require('mongodb').Code,
	assert = require('assert');

let minisites = [
	{
		id: 'link-0',
		idMinisite: '2343',
		raisonSociale: 'Fullstack tutorial for GraphQL'
	}
];
let idCount = minisites.length;

var db = new Db('test', new Server('127.0.0.1', 27017));
console.log('db connected ? ' + db._state);

if (db._state == 'connected') {
	db.open(function(err, db) {
		assert.equal(null, err);

		db.on('close', test.done.bind(test));
		db.close();
	});
}

// 2
const resolvers = {
	Query: {
		minisites: () => minisites
	},
	// 3
	Minisite: {
		id: root => root.id,
		idMinisite: root => root.idMinisite,
		raisonSociale: root => root.raisonSociale
	},
	Mutation: {
		// 2
		post: (root, args) => {
			const minisite = {
				id: `link-${idCount++}`,
				raisonSociale: args.raisonSociale,
				idMinisite: args.idMinisite
			};
			minisites.push(minisite);
			return minisite;
		}
	}
};

// 3
const serverGraphQL = new GraphQLServer({
	typeDefs: './schema.graphql',
	resolvers
});
serverGraphQL.start(() => console.log(`Server is running on http://localhost:4000`));
