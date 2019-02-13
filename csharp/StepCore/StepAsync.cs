using System;
using System.Collections.Generic;
using System.Reflection;
using System.Threading.Tasks;
using Grpc.Core;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace StepCore {
    public abstract class StepAsync {
        public abstract Task ExecuteAsync();

        private Cache _cache;

        public void BindInputs(string input) {
            var jObject = JObject.Parse(input);
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

        public Cache GetCache() {
            if (_cache != null) {
                return _cache;
            }

            var conn = System.Environment.GetEnvironmentVariable("WORKFLOW_CACHE_CONNECTION");
            Console.WriteLine(conn);
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