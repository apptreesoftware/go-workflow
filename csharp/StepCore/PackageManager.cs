using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Threading.Tasks;
using Google.Protobuf.Collections;
using Newtonsoft.Json;
using Newtonsoft.Json.Serialization;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace StepCore {
    public static class PackageManager {

        public static async Task<int> Run(string[] args) {
            var environment = Util.GetEnvironment();
            var stepName = environment.StepName;
            var stepVersion = environment.StepVersion;
            if (string.IsNullOrEmpty(stepVersion)) {
                stepVersion = "1.0.0";
            }
            if (string.IsNullOrEmpty(stepName)) {
                var yaml = GeneratePackageInfoYaml();
                Console.WriteLine(yaml);
                return 0;
            }


            var stdin = Console.ReadLine();
            var step = GetStep(stepName, stepVersion);
            if (step == null) {
                Console.WriteLine($"Could not find step {stepName}@{stepVersion}");
                return -1;
            }

            try {
                step.BindInputs(stdin);
                await step.ExecuteAsync();
                var output = step.GetOutputs();
                
                var outputLoc = System.Environment.GetEnvironmentVariable("WORKFLOW_OUTPUT");
                if (output != null && output.Count > 0) {
                    var outString = JsonConvert.SerializeObject(output);
                    if (string.IsNullOrEmpty(outputLoc)) {
                        Console.WriteLine(outString);
                    } else {
                        File.WriteAllText(outputLoc, outString);   
                    }
                }
            } catch (Exception e) {
                Console.WriteLine($"Step failed with exception: {e}");
            }
            return 0;
        }
        
        public static StepAsync GetStep(string name, string version) {
            return GetStepInAssembly(Assembly.GetEntryAssembly(), name, version);
        }

        public static StepAsync GetStepInAssembly(Assembly assembly, string name, string version) {
            var type = assembly.GetTypes()
                .FirstOrDefault(
                    t => {
                        var description = t.GetCustomAttribute<StepDescription>();
                        if (description == null) return false;
                        return description.Name == name && description.Version == version;
                    });
            if (type == null) {
                return null;
            }

            return Activator.CreateInstance(type) as StepAsync;
        }


        public static Package GeneratePackageInfoFromAssembly(Assembly assembly) {
            var packageDef = assembly.GetTypes()
                .FirstOrDefault(t => Attribute.IsDefined(t, typeof(PackageDefinition)));
            if (packageDef == null)
                throw new Exception("No Package attribute defined in this project");

            var packageAttribute = packageDef.GetCustomAttribute<PackageDefinition>();
            var package = new Package {
                Name = packageAttribute.Name,
                Lang = "NETCORE21",
                Exec = Path.GetFileName(Assembly.GetEntryAssembly().Location)
            };

            var stepClass = assembly.GetTypes()
                .Where(t => Attribute.IsDefined(t, typeof(StepDescription)))
                .ToList();

            foreach (var type in stepClass) {
                var stepDef = type.GetCustomAttribute<StepDescription>();
                var stepId = $"{stepDef.Name}@{stepDef.Version}";
                var packageStep = new PackageStep {
                    Description = stepDef.Description,
                };

                var properties = type.GetProperties();
                foreach (var propertyInfo in properties) {
                    var inputAttribute = propertyInfo.GetCustomAttribute<Input>();
                    if (inputAttribute != null) {
                        var key = inputAttribute.Key ?? propertyInfo.Name;
                        packageStep.Inputs[key] = new InputInfo {
                            Required = inputAttribute.Required,
                            Description = inputAttribute.Description ?? "",
                        };
                    }

                    var outputAttribute = propertyInfo.GetCustomAttribute<Output>();
                    if (outputAttribute != null) {
                        var key = outputAttribute.Key ?? propertyInfo.Name;
                        packageStep.Outputs[key] = new OutputInfo {
                            Description = outputAttribute.Description ?? ""
                        };
                    }
                }

                package.Steps[stepId] = packageStep;
            }

            return package;
        }

        public static Package GeneratePackageInfo() {
            return GeneratePackageInfoFromAssembly(Assembly.GetEntryAssembly());
        }

        public static string GeneratePackageInfoYaml() {
            var serializer = new SerializerBuilder()
                .DisableAliases()
                .WithNamingConvention(new UnderscoredNamingConvention())
                .Build();
            return serializer.Serialize(GeneratePackageInfo());
        }
        
    }
   
}