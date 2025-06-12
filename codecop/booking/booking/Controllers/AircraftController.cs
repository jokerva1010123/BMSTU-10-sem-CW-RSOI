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
    public class AircraftController : Controller
    {
        private readonly IAircraftService _aircraftService;

        public AircraftController(IAircraftService aircraftService)
        {
            _aircraftService = aircraftService;
        }

        [HttpPost]
        public async Task<ActionResult> Create([FromBody]AircraftModel model)
        {
            await _aircraftService.Create(model);
            return Ok();
        }

        [HttpGet]
        public async Task<ActionResult> GetAll()
        {
            var aircrafts = await _aircraftService.GetAll(0, 0);
            return Ok(aircrafts);
        }

        [HttpGet("{id}")]
        public async Task<ActionResult> Get(string id)
        {
            var aircraft = await _aircraftService.GetById(id);
            return Ok(aircraft);
        }

        [HttpPut("{id}")]
        public async Task<ActionResult> UpdateFlight(string id, [FromBody]AircraftModel model)
        {
            await _aircraftService.Update(id, model);
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult> DeleteFlight(string id)
        {
            await _aircraftService.Remove(id);
            return Ok();
        }
    }
}
