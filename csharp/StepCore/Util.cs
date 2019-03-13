using Core;

namespace StepCore {
    public static class Util {
        public static Environment GetEnvironment() {
            return new Environment {
                Project = System.Environment.GetEnvironmentVariable("PROJECT") ?? "",
                RunId = System.Environment.GetEnvironmentVariable("RUN_ID") ?? "",
                Workflow = System.Environment.GetEnvironmentVariable("WORKFLOW_ID") ?? "",
                StepName = System.Environment.GetEnvironmentVariable("STEP_NAME") ?? "",
                StepVersion = System.Environment.GetEnvironmentVariable("STEP_VERSION") ?? "",
                InputFile = System.Environment.GetEnvironmentVariable("INPUT_FILE") ?? ""
            };
        }
    }
}