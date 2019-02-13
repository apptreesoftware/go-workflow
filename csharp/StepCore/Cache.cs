using System;
using System.Collections.Generic;
using System.IO;
using System.Threading.Tasks;
using Google.Protobuf;
using Newtonsoft.Json;
using Newtonsoft.Json.Bson;

namespace StepCore {
    public interface ICache {
        Task PutRecord<T>(
            CacheRecord<T> record,
            string cacheName = "");

        Task<CacheRecord<T>> PullRecord<T>(string id, string cacheName = "");
    }

    public class Cache : ICache {
        private global::Cache.CacheClient _client;

        public Cache(global::Cache.CacheClient client) {
            _client = client;
        }

        public async Task PutRecord<T>(
            CacheRecord<T> cacheRecord,
            string cacheName = "") {
            var req = new CachePushRequest();
            req.Record = GetBytes(cacheRecord.Record);
            if (cacheRecord.Metadata != null) {
                req.Metadata = GetBytes(cacheRecord.Metadata);
            }
            req.Id = cacheRecord.Id;
            req.Environment = Util.GetEnvironment();
            req.CacheName = cacheName;
            await _client.PushAsync(req);
        }

        public async Task<CacheRecord<T>> PullRecord<T>(string id, string cacheName = "") {
            if (id == null) {
                Console.WriteLine("WARNING: Request record from cache with a null id");
                return null;
            }

            Console.WriteLine($"{id} - {cacheName} - {Util.GetEnvironment()}");
            var req = new CachePullRequest {
                Id = id, Environment = Util.GetEnvironment(), CacheName = cacheName
            };

            var response = await _client.PullAsync(req);
            if (response.NotFound) {
                return null;
            }

            if (response.Record == null) {
                return null;
            }

            var record = Deserialize<T>(response.Record);
            var meta = response.Metadata.Length > 0
                ? Deserialize<Dictionary<string, object>>(response.Metadata)
                : new Dictionary<string, object>();

            return new CacheRecord<T> {
                Id = id,
                Record = record,
                Metadata = meta,
            };
        }

        private T Deserialize<T>(ByteString byteString) {
            var ms = new MemoryStream(byteString.ToByteArray());
            using (var reader = new BsonDataReader(ms)) {
                var serializer = new JsonSerializer();
                return serializer.Deserialize<T>(reader);
            }
        }

        private ByteString GetBytes(object record) {
            var memoryStream = new MemoryStream();
            using (var writer = new BsonDataWriter(memoryStream)) {
                var serializer = new JsonSerializer();
                serializer.Serialize(writer, record);
                return ByteString.CopyFrom(memoryStream.ToArray());
            }
        }
    }

    public class NoOpCache : ICache {
        public Task PutRecord<T>(
            CacheRecord<T> cacheRecord,
            string cacheName = null) {
            return Task.CompletedTask;
        }

        public Task<CacheRecord<T>> PullRecord<T>(string id, string cacheName = null) {
            return Task.FromResult<CacheRecord<T>>(null);
        }
    }

    public class CacheRecord<T> {
        public CacheRecord(string id, T record, Dictionary<string, object> metadata) {
            Id = id;
            Record = record;
            Metadata = metadata;
        }

        public CacheRecord() { }
        public string Id { get; set; }
        public T Record { get; set; }
        public Dictionary<string, object> Metadata { get; set; }
    }
}