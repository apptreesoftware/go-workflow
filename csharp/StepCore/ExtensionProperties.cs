using System.Collections.Generic;

namespace Steps {
    public static class ExtensionProperties {
        private static Dictionary<object, Dictionary<string, object>> PropertyValues =
            new Dictionary<object, Dictionary<string, object>>();

        // Set a property value for the item.
        private static void SetValue(this object item, string name, object value) {
            // If we don't have a dictionary for this item yet,
            // make one.
            if (!PropertyValues.ContainsKey(item))
                PropertyValues[item] = new Dictionary<string, object>();

            // Set the value in the item's dictionary.
            PropertyValues[item][name] = value;
        }

        // Return a property value for the item.
        private static object GetValue(this object item, string name, object default_value) {
            // If we don't have a dictionary for
            // this item yet, return the default value.
            if (!PropertyValues.ContainsKey(item)) return default_value;

            // If the value isn't in the dictionary,
            // return the default value.
            if (!PropertyValues[item].ContainsKey(name))
                return default_value;

            // Return the saved value.
            return PropertyValues[item][name];
        }

        // Remove the property.
        private static void RemoveValue(this object item, string name) {
            // If we don't have a dictionary for this item, do nothing.
            if (!PropertyValues.ContainsKey(item)) return;

            // If the value isn't in the dictionary, do nothing.
            if (!PropertyValues[item].ContainsKey(name)) return;

            // Remove the value.
            PropertyValues[item].Remove(name);

            // If the dictionary is empty, remove it.
            if (PropertyValues[item].Count == 0)
                PropertyValues.Remove(PropertyValues[item]);
        }

        // Remove the object's property dictionary.
        private static void RemoveAllValues(this object item) {
            // If we have a dictionary for this item, remove it.
            if (PropertyValues.ContainsKey(item))
                PropertyValues.Remove(PropertyValues[item]);
        }

        public static void SetCacheId(this object item, string sourceId) {
            item.SetCacheProperty("id", sourceId);
        }

        public static string GetCacheId(this object item) {
            return item.GetCacheProperty<string>("id");
        }

        public static Dictionary<string, object> CacheProperties(this object item) {
            return item.GetValue("_cache", item) as Dictionary<string, object>;
        }

        public static T GetCacheProperty<T>(this object item, string key) {
            var cacheDict = item.GetValue("_cache", item) as Dictionary<string, object>;
            var val       = cacheDict?.GetValueOrDefault(key, defaultValue: default(T));
            return (T) val;
        }

        public static void SetCacheProperty<T>(this object item, string key, T value) {
            var cacheDict = item.GetValue("_cache", item) as Dictionary<string, object> ??
                            new Dictionary<string, object>();
            cacheDict[key] = value;
            item.SetValue("_cache", cacheDict);
        }
    }
}