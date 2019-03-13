using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using Newtonsoft.Json;

namespace StepCore {
    [Route("/")]
    [ApiController]
    public class TriggerController : ControllerBase {
        [HttpPost]
        public async Task<ActionResult<object>> Post(StepRequest request) {
            try {
                var output = await StepRunner.Run(
                    request.Environment, JsonConvert.SerializeObject(request.Inputs));
                return output;
            } catch (Exception e) {
                Console.WriteLine(e.Message);
                Console.WriteLine(e.StackTrace);
                return new BadRequestObjectResult(e.ToString());
            }
        }

        [HttpGet]
        public ActionResult<string> Get() {
            var result = PackageManager.GeneratePackageInfoYaml();
            return result;
        }
    }

    public class StepRequest {
        public Core.Environment Environment { get; set; }
        public Dictionary<string, object> Inputs { get; set; }
    }
}