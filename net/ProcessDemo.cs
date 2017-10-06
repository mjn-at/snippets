using System;
using System.Diagnostics;

class ProcessDemo    
{
  static void Main()  
  {
    LaunchCommandLineApp();
  } 
 static void LaunchCommandLineApp() 
 {  
  ProcessStartInfo startInfo = new ProcessStartInfo();  
  startInfo.CreateNoWindow = false;  
  startInfo.UseShellExecute = false;  
  startInfo.FileName = "C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe";  
  startInfo.Arguments = "-c Write-Host SUCCESS -Fore Green";  
  try         
    {            
      using (Process exeProcess = Process.Start(startInfo))
      {
        exeProcess.WaitForExit();
      }
    }
    catch 
    {
      Console.WriteLine(startInfo.ToString());
    }
  }
}
