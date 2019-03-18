using System;
using System.Threading.Tasks;
using Core;
using Google.Protobuf;
using Grpc.Core;
using Newtonsoft.Json;

namespace StepCore {
    public class Server : StepHost.StepHostBase {

        public static void Run(int port) {
            var server = new Grpc.Core.Server
            {
                Services = { StepHost.BindService(new Server()) },
                Ports = { new ServerPort("localhost", port, ServerCredentials.Insecure) }
            };
            server.Start();

            Console.WriteLine("Step package listening on port " + port);
            Console.WriteLine("Press `Return` key to stop the server...");
            Console.ReadLine();
            server.ShutdownAsync().Wait();
        }
        
        public override async Task<StepOutput> RunStep(RunStepRequest request, ServerCallContext context) {

            var inputBytes = request.Input;
            var inputStr = inputBytes.ToStringUtf8();

            try {
                var output = await StepRunner.Run(request.Environment, inputStr);

                ByteString byteString = null;

                if (output != null) {
                    var str = JsonConvert.SerializeObject(output);
                    byteString = ByteString.CopyFromUtf8(str);
                }

                return new StepOutput {
                    Output = byteString
                };
            } catch (Exception e) {
                throw new RpcException(new Status(StatusCode.Internal, e.Message));
            }
            
        }

        public override Task<Package> GetPackageInfo(EmptyMessage request, ServerCallContext context) {
            var package = PackageManager.GeneratePackageInfo();
            return Task.FromResult(package);
        }
    }
}