# Publishing an new version of the SDK to NUGET

To perform the following, you must have the `dotnet` [cli tools installed](https://docs.microsoft.com/en-us/dotnet/core/tools/?tabs=netcore2x)

You must also have the AppTree Nuget publish API Key. 

### 1. Update Package Version
Modify StepCore Project to increment the `Nuget Version`

### 2. Package the SDK
`dotnet pack -c Release`

### 3. Publish the new version to NUGET

nuget push bin/Release/StepCore.1.0.0-alpha2.nupkg -Source https://api.nuget.org/v3/index.json