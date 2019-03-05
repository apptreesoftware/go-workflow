using System;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Threading.Tasks;
using Newtonsoft.Json;
using YamlDotNet.Serialization;
using YamlDotNet.Serialization.NamingConventions;

namespace StepCore {
    public static class PackageManager {
        public static async Task<int> Run(string[] args) {
            bool serveMode = false;
            var port = 5000;
            for (var i = 0; i < args.Length; i++)
            {
                if (args[i] == "--serve")
                {
                    serveMode = true;
                    
                    if (args.Length > i + 1)
                    {
                        var portString = args[i + 1];
                        int.TryParse(portString, out port);
                    }
                } 
            }
            if (serveMode) {
                Server.RunServer(args, port);
            } else {
                await Run();
            }
            return 0;
        }

        public static async Task Run() {
            var environment = Util.GetEnvironment();
            var input = "";
            if (!string.IsNullOrEmpty(environment.InputFile)) {
                input = File.ReadAllText(environment.InputFile);
            } else {
                input = Console.ReadLine();
            }

            try {
                var output = await StepRunner.Run(environment, input);
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
                    Description = stepDef.Description ?? "",
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