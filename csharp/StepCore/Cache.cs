using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Google.Protobuf;
using Newtonsoft.Json;
using Newtonsoft.Json.Bson;

namespace StepCore {
    public class Cache {
        private global::Cache.CacheClient _client;
        
        public Cache(global::Cache.CacheClient client) {
            _client = client;
        }

        public async Task PutRecord(string id, object record, string cacheName = null, Dictionary<string, object> meta = null) {
            var memoryStream = new MemoryStream();
            using (BsonDataWriter writer = new BsonDataWriter(memoryStream)) {
                var serializer = new JsonSerializer();
                serializer.Serialize(writer, record);
            }

            var req = new CachePushRequest();
            req.Record = GetBytes(record);
            if (meta != null) {
                req.Metadata = GetBytes(meta);    
            }

            req.Id = id;
            req.Environment = Util.GetEnvironment();
            req.CacheName = cacheName;
            await _client.PushAsync(req);
        }

        public async Task<(T, Dictionary<string, object>)>PullRecord<T>(string id, string cacheName = null) {
            var req = new CachePullRequest {
                Id = id, Environment = Util.GetEnvironment(), CacheName = cacheName
            };

            var response = await _client.PullAsync(req);
            var record = Deserialize<T>(response.Record);
            var meta = response.Metadata.Length > 0
                ? Deserialize<Dictionary<string, object>>(response.Metadata) :
                  new Dictionary<string, object>();

            return (record, meta);
        }

        private T Deserialize<T>(ByteString byteString) {
            var ms = new MemoryStream(byteString.ToByteArray());
            using (var reader = new BsonDataReader(ms))
            {
                var serializer = new JsonSerializer();
                return serializer.Deserialize<T>(reader);
            }
        }

        private ByteString GetBytes(object record) {
            var memoryStream = new MemoryStream();
            using (var writer = new BsonDataWriter(memoryStream)) {
                var serializer = new JsonSerializer();
                serializer.Serialize(writer, record);
            }
            return ByteString.FromStream(memoryStream);
        }
        
    }
}