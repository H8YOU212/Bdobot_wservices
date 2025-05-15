using System.Collections.ObjectModel;
using MongoDB.Bson;
using MongoDB.Driver;

namespace DB{
    class CheckExistColl{
        private readonly IMongoDatabase _database;
        public CheckExistColl(IMongoDatabase db){
            _database = db;
        }

        public void CheckExist(string NameColl) {
            List<String> ListColl = _database.ListCollectionNames().ToList();
            if (!ListColl.Contains(NameColl)) {
                _database.CreateCollection(NameColl);
                Console.WriteLine($"The collection'{NameColl}' has been created");
            } else {
                Console.WriteLine($"The collection '{NameColl}' is exists");
            }
        }

    }
}