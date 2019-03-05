using Microsoft.AspNetCore;
using Microsoft.AspNetCore.Hosting;

namespace StepCore {
    public class Server {
        public static void RunServer(string[] args, int port)
        {
            CreateWebHostBuilder(args, port).Build().Run();
        }

        public static IWebHostBuilder CreateWebHostBuilder(string[] args, int port) =>
            WebHost.CreateDefaultBuilder(args)
                .UseUrls($"http://0.0.0.0:{port}")
                .UseStartup<Startup>();
    }
}