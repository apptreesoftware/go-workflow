
using System;
using System.Collections.Generic;
using System.Threading.Tasks;

namespace StepCore {
    public class StepRunner {
        public static async Task<Dictionary<string, object>>  Run(Core.Environment environment, string input) {
            var stepName = environment.StepName;
            var stepVersion = environment.StepVersion;
            if (string.IsNullOrEmpty(stepVersion)) {
                stepVersion = "1.0";
            }
            var step = PackageManager.GetStep(stepName, stepVersion);
            if (step == null) {
                throw new Exception("Step not found");
            }
            
            step.BindInputs(input);
            await step.ExecuteAsync();
            var output = step.GetOutputs();

            return output;
            
        }
    }
}