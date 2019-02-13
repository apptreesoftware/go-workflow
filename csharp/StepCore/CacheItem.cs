using System;
using System.Collections.Generic;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace StepCore {
   
    
    public class CacheItem<T>  where T : class {
        public object Record { get; set; }
        public string Id { get; set; }
        public Dictionary<string, object> Meta { get; set; }
        
        public CacheItem(string id, T record, Dictionary<string, object> meta) {
            Record = record ?? throw new ArgumentNullException(nameof(record));
            Id = id ?? throw new ArgumentNullException(nameof(id));
            Meta = meta ?? new Dictionary<string, object>();
        }
    }
}