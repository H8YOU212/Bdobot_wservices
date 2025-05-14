using System;
using DotNetEnv;
using Microsoft.VisualBasic;
using MongoDB.Bson;
using MongoDB.Driver;

namespace DB {
    public class Variable{
        public static string Uri { get; set; }
        public static string DbName { get; set; }
        public static string coll {get; set;}
        public static object client {get; set;}
        public static object db {get; set;}
    }

    class DefineENV{
        // Example of defining a variable using DotNetEnv
        public (String, String) LoadEnvironmentVariables() {
            DotNetEnv.Env.Load(@"DataLayer\.env");
            // Variable variable = new Variable();

            Variable.Uri = DotNetEnv.Env.GetString("uri");
            Variable.DbName = DotNetEnv.Env.GetString("dbname");

            if (string.IsNullOrEmpty(Variable.Uri) || string.IsNullOrEmpty(Variable.DbName)) {
                throw new ArgumentNullException("Переменные окружения 'uri' или 'dbname' не заданы.");
            }
            return (Variable.Uri, Variable.DbName);
        }
    }

    class Connection{
        public IMongoDatabase Conn() {
            DefineENV variables = new DefineENV();
            Variable variable = new Variable();

            (string Uri, string dbname) =  variables.LoadEnvironmentVariables();

            var client = new MongoClient(Uri);
            
            var db = client.GetDatabase(dbname);

            // var checkcoll = db.ListCollections().ToList();
            // foreach (var curcoll in checkcoll) {

            // }

            // Variable.coll = db.GetCollection<BsonDocument>("Items");
            return db;
        }
    }

    class Methods{
        private readonly IMongoDatabase _database;
        public Methods(IMongoDatabase database){
            _database = database;
        }
        public void Insert(string Coll, BsonDocument document) {
            var collection = _database.GetCollection<BsonDocument>(Coll);
            var collectionNames = _database.ListCollectionNames().ToList();
            foreach (var name in collectionNames)
            {
                Console.WriteLine(name);
            }
            collection.InsertOne(document);
        }

        public void Delete(string Coll, BsonDocument document) {
            var collection = _database.GetCollection<BsonDocument>(Coll);
        }

        public void Update(string Coll, BsonDocument document) {
            var collection = _database.GetCollection<BsonDocument>(Coll);
        }
        
        public List<BsonDocument> Search(string Coll, FilterDefinition<BsonDocument> filter) {
            var collection = _database.GetCollection<BsonDocument>(Coll);

            return collection.Find(filter).ToList();
        }
    }
}