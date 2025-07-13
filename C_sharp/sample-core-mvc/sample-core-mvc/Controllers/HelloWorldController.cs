using Microsoft.AspNetCore.Mvc;
using System.Text.Encodings.Web;

namespace MvcMovie.Controllers;

//Controller is base class from Microsoft.AspNetCore.Mvc
public class HelloWorldController : Controller
{
    //all public method in Controller class is handled as a callable http endpoint
    //Index() can be accessed with /HelloWorld
    public string Index()
    {
        return "This is my defaul action...";
    }

    //Welocme can be accessed with /HelloWorld/Welocme
    //name and numtimes are query parameter
    public string Welcome(string name, int numTimes=1)
    {
        return HtmlEncoder.Default.Encode($"Hello {name}, Numtime is ; {numTimes}");
    }

    //program.cs accecpts both Welcome/3?name=Rick&numtimes=4 and Welcome?name=Rick&numtimes=4
}