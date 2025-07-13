using Microsoft.AspNetCore.Mvc;
using System.Text.Encodings.Web;

namespace MvcMovie.Controllers;

//Controller is base class from Microsoft.AspNetCore.Mvc
public class HelloWorldController : Controller
{
    //all public method in Controller class is handled as a callable http endpoint
    //Index() can be accessed with /HelloWorld
    //public string Index()
    public IActionResult Index()
    {
        //return "This is my defaul action...";
        return View();  
    }

    //Welocme can be accessed with /HelloWorld/Welocme
    //name and numtimes are query parameter
    public string Welcome(string name, int numTimes=1)
    {
        return HtmlEncoder.Default.Encode($"Hello {name}, Numtime is ; {numTimes}");
    }

    public ActionResult Welcome2()
    {
        return View();
    }

    //program.cs accecpts both Welcome/3?name=Rick&numtimes=4 and Welcome?name=Rick&numtimes=4
}