using System;

namespace StepCore {
    [AttributeUsage(AttributeTargets.Property, AllowMultiple = false)]
    public class Input : Attribute {
        public string Key { get; set; }
        public string Description { get; set; }
        public bool Required { get; set; }
        public object Default { get; set; }
        public string Sample { get; set; }

        public Input() {
            Required = true;
        }
    }

    [AttributeUsage(AttributeTargets.Property)]
    public class Output : Attribute {
        public string Key { get; set; }
        public string Description { get; set; }
       
    }

    [AttributeUsage(AttributeTargets.Class)]
    public class StepDescription : Attribute {
        public string Description { get; set; }

        public string Name { get; }
        public string Version { get; }

        public StepDescription(string Name, string Version = "1.0") {
            this.Name = Name;
            this.Version = Version;
        }
    }

    [AttributeUsage(AttributeTargets.Class)]
    public class PackageDefinition : Attribute {
        public PackageDefinition(string name) {
            Name = name;
        }
        public string Name { get; }        
    }
}