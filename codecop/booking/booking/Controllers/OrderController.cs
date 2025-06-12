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
    public class OrderController : Controller
    {
        private readonly IOrderService _orderService;

        public OrderController(IOrderService orderService)
        {
            _orderService = orderService;
        }

        [HttpPost]
        public async Task<ActionResult> CreateFlight([FromBody]OrderModel model)
        {
            await _orderService.Create(model);
            return Ok();
        }

        [HttpGet]
        public async Task<ActionResult> GetAllFlights()
        {
            var orders = await _orderService.GetAll(0, 0);
            return Ok(orders);
        }

        [HttpGet("{id}")]
        public async Task<ActionResult> GetFlight(string id)
        {
            var order = await _orderService.GetById(id);
            return Ok(order);
        }

        [HttpPut("{id}")]
        public async Task<ActionResult> UpdateFlight(string id, [FromBody]OrderModel model)
        {
            await _orderService.Update(id, model);
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult> DeleteFlight(string id)
        {
            await _orderService.Remove(id);
            return Ok();
        }
    }
}
