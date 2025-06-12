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
    public class FlightController : ControllerBase
    {
        private IFlightRepository flightRepository;

        public IAircraftRepository aircraftRepository { get; }

        public FlightController(IFlightRepository flightRepository, IAircraftRepository aircraftRepository)
        {
            this.flightRepository = flightRepository;
            this.aircraftRepository = aircraftRepository;
        }

        [HttpGet]
        public ActionResult<IEnumerable<FlightModel>> GetAllFlights([FromQuery]int page, [FromQuery]int amount)
        {
            var flights = flightRepository.GetAll();
            if (page != 0 && amount != 0)
            {
                flights = flights.Skip(page * (amount - 1)).Take(amount);
            }

            if (flights == null)
                return BadRequest();


            return Ok(flights.Select( x=> new FlightModel()
            {
                Id = x.Id,
                AircraftId = x.AircraftId,
                Date = x.Date,
                FreeSeats = x.FreeSeats,
                Number = x.Number,
                Sum = x.Sum
            }).ToList());
        }

        // GET
        [HttpGet("{id}")]
        public ActionResult<FlightModel> Get(string id)
        {
            var flight = flightRepository.Get(id);
            if (flight == null)
                return null;
                //return BadRequest();


            return Ok(new FlightModel()
            {
                Id = flight.Id,
                AircraftId = flight.AircraftId,
                Date = flight.Date,
                FreeSeats = flight.FreeSeats,
                Number = flight.Number,
                Sum = flight.Sum
            });
        }

        // добавить рейс
        [HttpPost]
        public ActionResult Post([FromBody] FlightModel model)
        {
            var flight = new Flight
            {
                AircraftId = model.AircraftId,
                Date = model.Date,
                FreeSeats = model.FreeSeats,
                Sum = model.Sum,
                Number = model.Number
            };

            flightRepository.Add(flight);
            return Ok();
        }

        // обновить рейс
        [HttpPut("{id}")]
        public ActionResult Put(string id, [FromBody] FlightModel model)
        {
            var flight = flightRepository.Get(id);
            if (flight == null)
            {
                return NotFound();
            }

            flight.AircraftId = model.AircraftId;
            flight.Date = model.Date;
            flight.FreeSeats = model.FreeSeats;
            flight.Sum = model.Sum;
            flight.Number = model.Number;

            flightRepository.Update(flight);

            return Ok();

        }

        // DELETE api/values/5
        [HttpDelete("{id}")]
        public ActionResult Delete(String id)
        {
            flightRepository.Delete(id);
            return Ok();
        }
    }
}
