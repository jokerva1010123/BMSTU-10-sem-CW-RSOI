using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using booking.flight.Abstract;
using booking.flight.Model;
using booking.common.ViewModel;

namespace booking.flight.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class AircraftController : ControllerBase
    {
        private IAircraftRepository aircraftRepository;

        public AircraftController(IAircraftRepository aircraftRepository)
        {
            this.aircraftRepository = aircraftRepository;
        }

        [HttpGet]
        public ActionResult<IEnumerable<AircraftModel>> GetAll([FromQuery]int page, [FromQuery]int amount)
        {
            var aircrafts = aircraftRepository.GetAll();
            if (page != 0 && amount != 0)
            {
                aircrafts = aircrafts.Skip(page * (amount - 1)).Take(amount);
            }

            if (aircrafts == null)
                return BadRequest();


            return Ok(aircrafts.Select(x => new AircraftModel()
            {
                Id = x.Id,
                Name = x.Name,
                NumberOfSeats = x.NumberOfSeats
            }).ToList());
        }

        // GET
        [HttpGet("{id}")]
        public ActionResult<AircraftModel> Get(string id)
        {
            var aircraft = aircraftRepository.Get(id);
            if (aircraft == null)
                return BadRequest();


            return Ok(new AircraftModel()
            {
                Id = aircraft.Id,
                Name = aircraft.Name,
                NumberOfSeats = aircraft.NumberOfSeats
            });
        }

        // добавить рейс
        [HttpPost]
        public ActionResult Post([FromBody] AircraftModel model)
        {
            var aircraft = new Aircraft
            {
                Name = model.Name,
                NumberOfSeats = model.NumberOfSeats
            };

            aircraftRepository.Add(aircraft);
            return Ok();
        }

        // обновить рейс
        [HttpPut("{id}")]
        public ActionResult Put(string id, [FromBody] AircraftModel model)
        {
            var aircraft = aircraftRepository.Get(id);
            if (aircraft == null)
            {
                return NotFound();
            }
            aircraft.Name = model.Name;
            aircraft.NumberOfSeats = model.NumberOfSeats;

            aircraftRepository.Update(aircraft);

            return Ok();

        }

        // DELETE api/values/5
        [HttpDelete("{id}")]
        public ActionResult Delete(string id)
        {
            aircraftRepository.Delete(id);
            return Ok();
        }
    }
}
