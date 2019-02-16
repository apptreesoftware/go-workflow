using Microsoft.AspNetCore;
using Microsoft.AspNetCore.Hosting;

namespace StepCore {
    public class Server {
        public static void RunServer(string[] args)
        {
            CreateWebHostBuilder(args).Build().Run();
        }

        public static IWebHostBuilder CreateWebHostBuilder(string[] args) =>
            WebHost.CreateDefaultBuilder(args)
                .UseStartup<Startup>();
    }
}