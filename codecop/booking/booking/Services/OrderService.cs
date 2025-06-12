using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Threading.Tasks;
using booking.common.ViewModel;
using booking.Infrustructure;
using Microsoft.Extensions.Logging;
using Microsoft.Extensions.Options;
using Newtonsoft.Json;

namespace booking.Services
{
    public class OrderService : IOrderService
    {
        private readonly HttpClient _httpClient;
        private readonly ILogger<OrderService> _logger;
        private readonly UrlHosts _urls;


        public OrderService(HttpClient httpClient, ILogger<OrderService> logger, IOptions<UrlHosts> config)
        {
            _httpClient = httpClient;
            _logger = logger;
            _urls = config.Value;
        }

        public async Task Create(OrderModel model)
        {
            await _httpClient.PostAsJsonAsync(_urls.Order + "/api/order", model);
        }

        public async Task<IEnumerable<OrderModel>> GetAll(int page, int size)
        {
            var data = await _httpClient.GetStringAsync(_urls.Order + $"/api/order?page={page}&size={size}");
            var orders = !string.IsNullOrEmpty(data)
                ? JsonConvert.DeserializeObject<IEnumerable<OrderModel>>(data)
                : null;
            return orders;
        }

        public async Task<OrderModel> GetByFlightId(string id)
        {
            var data = await _httpClient.GetStringAsync(_urls.Order + $"/api/order/getbyflightid/{id}");
            var order = !string.IsNullOrEmpty(data) ? JsonConvert.DeserializeObject<OrderModel>(data) : null;
            return order;
        }

        public async Task<OrderModel> GetById(string id)
        {
            var data = await _httpClient.GetStringAsync(_urls.Order + $"/api/order/{id}");
            var order = !string.IsNullOrEmpty(data) ? JsonConvert.DeserializeObject<OrderModel>(data) : null;
            return order;
        }

        public async Task Remove(string id)
        {
            await _httpClient.DeleteAsync(_urls.Order + $"/api/order/{id}");
        }

        public async Task RemoveByFlightId(string id)
        {
            await _httpClient.DeleteAsync(_urls.Order + $"/api/order/DeleteByFlightId/{id}");
        }

        public async Task Update(string id, OrderModel model)
        {
            await _httpClient.PutAsJsonAsync(_urls.Order + $"/api/order/{id}", model);
        }
    }
}
