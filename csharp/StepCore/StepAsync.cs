using System;
using System.Collections.Generic;
using System.IO;
using System.Reflection;
using System.Threading.Tasks;
using Grpc.Core;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace StepCore {
    public abstract class StepAsync {
        public abstract Task ExecuteAsync();

        private ICache _cache;

        public void BindInputs(string input)
        {
            
            //var jObject = JObject.Parse(input);
            var jsonReader = new JsonTextReader(new StringReader(input));
            jsonReader.DateParseHandling = DateParseHandling.None;
            var jObject = JObject.Load(jsonReader);
            var type = GetType();

            foreach (var propertyInfo in type.GetProperties()) {
                var attr = propertyInfo.GetCustomAttribute<Input>();
                if (attr == null) continue;
                var key = attr.Key ?? propertyInfo.Name;
                var jsValue = jObject[key];
                var obj = jsValue?.ToObject(propertyInfo.PropertyType);
                propertyInfo.SetValue(this, obj);
            }
        }

        public ICache GetCache() {
            if (_cache != null) {
                return _cache;
            }

            var conn = System.Environment.GetEnvironmentVariable("WORKFLOW_CACHE_CONNECTION");
            if (string.IsNullOrEmpty(conn)) {
                Console.WriteLine("Not connected to cache service. Returning a dummy cache");
                _cache = new NoOpCache();
                return _cache;
            }
           
            var channel = new Channel(conn, ChannelCredentials.Insecure);
            var client = new global::Cache.CacheClient(channel);
            _cache = new Cache(client);
            return _cache;
        }

        public Dictionary<string, object> GetOutputs() {
            var dictionary = new Dictionary<string, object>();
            var type = GetType();
            foreach (var propertyInfo in type.GetProperties()) {
                var attr = propertyInfo.GetCustomAttribute<Output>();
                if (attr == null) continue;
                var key = attr.Key ?? propertyInfo.Name;
                var value = propertyInfo.GetValue(this);
                if (value == null) continue;
                dictionary[key] = propertyInfo.GetValue(this);
            }

            return dictionary;
        }
    }

    public abstract class Step : StepAsync {
        public abstract void Execute();

        public override Task ExecuteAsync() {
            Execute();
            return Task.CompletedTask;
        }
    }
}