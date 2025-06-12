using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using System.Net.Http;
using Newtonsoft.Json;
using booking.common.ViewModel;
using System.Text;
using booking.flight.Model;
using booking.Services;

// For more information on enabling Web API for empty projects, visit https://go.microsoft.com/fwlink/?LinkID=397860

namespace booking.Controllers
{
    [Route("api/[controller]")]
    public class FlightController : Controller
    {
        private readonly IFlightService _flightService;

        public FlightController(IFlightService flightService)
        {
            _flightService = flightService;
        }
        //http://localhost:5000/api/flight/
        [HttpPost]
        public async Task<ActionResult> CreateFlight([FromBody]FlightModel model)
        {
            await _flightService.Create(model);
            return Ok();
        }

        [HttpGet]
        public async Task<ActionResult> GetAllFlights()
        {
            var flights = await _flightService.GetAll(0, 0);
            return Ok(flights);
        }

        [HttpGet("{id}")]
        public async Task<ActionResult> GetFlight(string id)
        {
            var flight = await _flightService.GetById(id);
            return Ok(flight);
        }

        [HttpPut("{id}")]
        public async Task<ActionResult> UpdateFlight(string id, [FromBody]FlightModel model)
        {
            await _flightService.Update(id, model);
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult> DeleteFlight(string id)
        {
            await _flightService.Remove(id);
            return Ok();
        }
    }
}
