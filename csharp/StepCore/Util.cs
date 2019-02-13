namespace StepCore {
    public class Util {
        public static Environment GetEnvironment() {
            return new Environment {
                App = System.Environment.GetEnvironmentVariable("APP") ?? "",
                RunId = System.Environment.GetEnvironmentVariable("RUN_ID") ?? "",
                WorkflowId = System.Environment.GetEnvironmentVariable("WORKFLOW_ID") ?? "",
                StepName = System.Environment.GetEnvironmentVariable("STEP_NAME") ?? "",
                StepVersion = System.Environment.GetEnvironmentVariable("STEP_VERSION") ?? "",
                InputFile = System.Environment.GetEnvironmentVariable("INPUT_FILE") ?? ""
            };
        }
    }
}