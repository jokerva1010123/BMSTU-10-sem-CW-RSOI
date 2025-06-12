using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using booking.order.Abstract;
using booking.order.Model;
using booking.common.ViewModel;

namespace booking.order.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class OrderController : ControllerBase
    {
        private IOrderRepository orderRepository;

        public OrderController(IOrderRepository orderRepository)
        {
            this.orderRepository = orderRepository;
        }


        [HttpGet]
        public ActionResult<IEnumerable<OrderModel>> GetAllOrders([FromQuery]int page, [FromQuery]int amount)
        {
            var orders = orderRepository.GetAll();
            if (page != 0 && amount != 0)
            {
                orders = orders.Skip(page * (amount - 1)).Take(amount);
            }

            if (orders == null)
                return null;
            // return BadRequest();


            return Ok(orders.Select(x => new OrderModel()
            {
                Id = x.Id,
                ClientId = x.ClientId,
                FlightId = x.FlightId,
                Status = x.Status,
                Summ = x.Summ
            }).ToList());
        }

        // GET 
        [HttpGet("{id}")]
        public ActionResult<OrderModel> Get(string id)
        {
            var order = orderRepository.Get(id);
            if (order == null)
                return null;
            //return BadRequest();

            return Ok(new OrderModel()
            {
                Id = order.Id,
                ClientId = order.ClientId,
                FlightId = order.FlightId,
                Status = order.Status,
                Summ = order.Summ
            });
        }


        [HttpGet("[action]/{id}")]
        public ActionResult<OrderModel> GetByFlightId(string id)
        {
            var order = orderRepository.GetbyFlightId(id);
            if (order == null)
                return null;
            //  return BadRequest();

            return Ok(new OrderModel()
            {
                Id = order.Id,
                ClientId = order.ClientId,
                FlightId = order.FlightId,
                Status = order.Status,
                Summ = order.Summ
            });
        }


        // POST api/values
        [HttpPost]
        public ActionResult Post([FromBody] OrderModel model)
        {
            var order = new Order
            {
                ClientId = model.ClientId,
                FlightId = model.FlightId,
                Status = model.Status,
                Summ = model.Summ
            };

            orderRepository.Add(order);
            return Ok();
        }

        // обновить заказ
        [HttpPut("{id}")]
        public ActionResult Put(string id, [FromBody] OrderModel model)
        {
            var order = orderRepository.Get(id);
            if (order == null)
            {
                return NotFound();
            }
            order.ClientId = model.ClientId;
            order.FlightId = model.FlightId;
            order.Status = model.Status;
            order.Summ = model.Summ;

            orderRepository.Update(order);

            return Ok();
        }

        // DELETE 
        [HttpDelete("{id}")]
        public ActionResult Delete(string id)
        {
            orderRepository.Delete(id);
            return Ok();
        }

        [HttpDelete("[action]/{id}")]
        public ActionResult DeleteByFlightId(string id)
        {
            orderRepository.DeleteByFlightId(id);
            return Ok();
        }
    }
}
