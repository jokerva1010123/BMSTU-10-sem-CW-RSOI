using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using System.Net.Http;
using Newtonsoft.Json;
using booking.client.Model;
using booking.flight.Model;
using booking.Services;
using booking.common.ViewModel;

namespace booking.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class BookingController : ControllerBase
    {
        private readonly IHttpClientFactory clientFactory;
        private readonly IOrderService _orderService;
        private readonly IClientService _clientService;
        private readonly IFlightService _flightService;
        private readonly IAircraftService _aircraftService;

        public BookingController(IHttpClientFactory clientFactory,
            IOrderService orderService,
            IClientService clientService,
            IFlightService flightService,
            IAircraftService aircraftService)
        {
            this.clientFactory = clientFactory;
            _orderService = orderService;
            _clientService = clientService;
            _flightService = flightService;
            _aircraftService = aircraftService;
        }

        public class Res
        {
            public IEnumerable<ClientModel> Clients { get; set; }
            public IEnumerable<AircraftModel> Aircrafts { get; set; }
            public IEnumerable<FlightModel> Flights { get; set; }
            public IEnumerable<OrderModel> Orders { get; set; }

        }

        //вывести все данные
        [HttpGet("[action]")]
        public async Task<ActionResult> GetAll()
        {
            var res = new Res();
            res.Clients = await _clientService.GetAll(0, 0);
            res.Aircrafts = await _aircraftService.GetAll(0, 0);
            res.Flights = await _flightService.GetAll(0, 0);
            res.Orders = await _orderService.GetAll(0, 0);
            return Ok(res);
        }


        /*1. Get заказов по полю id рейса.
        2.Если не ноль, то запоминаем id заказа. (как тут? десериализовать строку и взять поле?) не совсем понятны действия
        3.Удаляем заказ
        4.Поиск заказа наверное надо зациклить, пока не 0
        5. Get рейса по id, если не 0, то update статус для жтого id
*/

        //http://localhost:5000/api/booking/<id>
        [HttpDelete("{id}")]
        public async Task<ActionResult> DeleteFlight(string id)
        {
            //1. Get заказов по полю id рейса.

            await _orderService.RemoveByFlightId(id);
            await _flightService.Remove(id);

            return Ok();
        }

        //
        [HttpPost("[action]")]
        public async Task<ActionResult> AddOrder([FromBody]OrderModel model)
        {
            var flight = await _flightService.GetById(model.FlightId);

            //1. Get flight по ID , проверка что есть Freeseats
            if (flight == null || flight.FreeSeats == 0)
                return BadRequest();//такого рейса нет

            if (flight.FreeSeats > 0)
            {
                flight.FreeSeats--;
                await _flightService.Update(model.FlightId, flight);
            }

            await _orderService.Create(model);

            return Ok();
        }
    }
}
