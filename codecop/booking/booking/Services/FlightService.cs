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
    public class FlightService : IFlightService
    {
        private readonly HttpClient _httpClient;
        private readonly ILogger<FlightService> _logger;
        private readonly UrlHosts _urls;


        public FlightService(HttpClient httpClient, ILogger<FlightService> logger, IOptions<UrlHosts> config)
        {
            _httpClient = httpClient;
            _logger = logger;
            _urls = config.Value;
        }

        public async Task Create(FlightModel model)
        {
            await _httpClient.PostAsJsonAsync(_urls.Flight + "/api/flight", model);
        }

        public async Task<IEnumerable<FlightModel>> GetAll(int page, int size)
        {
            var data = await _httpClient.GetStringAsync(_urls.Flight + $"/api/flight?page={page}&size={size}");
            var flights = !string.IsNullOrEmpty(data)
                ? JsonConvert.DeserializeObject<IEnumerable<FlightModel>>(data)
                : null;
            return flights;
        }

        public async Task<FlightModel> GetById(string id)
        {
            var data = await _httpClient.GetStringAsync(_urls.Flight + $"/api/flight/{id}");
            var flight = !string.IsNullOrEmpty(data) ? JsonConvert.DeserializeObject<FlightModel>(data) : null;
            return flight;
        }

        public async Task Remove(string id)
        {
            await _httpClient.DeleteAsync(_urls.Flight + $"/api/flight/{id}");
        }

        public async Task Update(string id, FlightModel model)
        {
            await _httpClient.PutAsJsonAsync(_urls.Flight + $"/api/flight/{id}", model);
        }
    }
}
