using booking.client.Model;

using Microsoft.AspNetCore.Mvc;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Newtonsoft.Json;
using booking.common.ViewModel;
using System.Collections;
using booking.Services;

namespace booking.Controllers
{
    [Route("api/[controller]")]
    [ApiController]
    public class ClientController : ControllerBase
    {
        private readonly IClientService _clientService;

        public ClientController(IClientService clientService)
        {
            _clientService = clientService;
        }

        [HttpPost]
        public async Task<ActionResult> Create([FromBody]ClientModel model)
        {
            await _clientService.Create(model);
            return Ok();
        }

        [HttpGet("{id}")]
        public async Task<ActionResult> GetbyId(string id)
        {
            var client = await _clientService.GetById(id);
            return Ok(client);
        }        

        [HttpGet]
        public async Task<ActionResult> GetAll()
        {
            var clients = await _clientService.GetAll(0, 0);
            return Ok(clients);
        }

        [HttpPut("{id}")]
        public async Task<ActionResult> Update(string id, [FromBody]ClientModel model)
        {
            await _clientService.Update(id, model);
            return Ok();
        }

        [HttpDelete("{id}")]
        public async Task<ActionResult> Delete(string id)
        {
            await _clientService.Remove(id);
            return Ok();
        }

        //[HttpGet]
        //public async Task<ActionResult<int>> GetCount()
        //{
        //    //try
        //    //{
        //    //    var request = new HttpRequestMessage(HttpMethod.Get, "http://localhost:5010/api/client/count");

        //    //    var client = clientFactory.CreateClient();
        //    //    var response = await client.SendAsync(request);
        //    //    var count = await response.Content.ReadAsAsync<int>();
        //    //    return Ok(count);
        //    //}
        //    //catch
        //    //{
        //    //    return BadRequest();
        //    //}
        //}
    }
}
