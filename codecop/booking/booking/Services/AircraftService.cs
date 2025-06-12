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
    public class AircraftService : IAircraftService
    {
        private readonly HttpClient _httpClient;
        private readonly ILogger<AircraftService> _logger;
        private readonly UrlHosts _urls;


        public AircraftService(HttpClient httpClient, ILogger<AircraftService> logger, IOptions<UrlHosts> config)
        {
            _httpClient = httpClient;
            _logger = logger;
            _urls = config.Value;
        }

        public async Task Create(AircraftModel model)
        {
            await _httpClient.PostAsJsonAsync(_urls.Flight + "/api/aircraft", model);
        }

        public async Task<IEnumerable<AircraftModel>> GetAll(int page, int size)
        {
            var data = await _httpClient.GetStringAsync(_urls.Flight + $"/api/aircraft?page={page}&size={size}");
            var aircrafts = !string.IsNullOrEmpty(data)
                ? JsonConvert.DeserializeObject<IEnumerable<AircraftModel>>(data)
                : null;
            return aircrafts;
        }

        public async Task<AircraftModel> GetById(string id)
        {
            var data = await _httpClient.GetStringAsync(_urls.Flight + $"/api/aircraft/{id}");
            var aircraft = !string.IsNullOrEmpty(data) ? JsonConvert.DeserializeObject<AircraftModel>(data) : null;
            return aircraft;
        }

        public async Task Remove(string id)
        {
            await _httpClient.DeleteAsync(_urls.Flight + $"/api/aircraft/{id}");
        }

        public async Task Update(string id, AircraftModel model)
        {
            await _httpClient.PutAsJsonAsync(_urls.Flight + $"/api/aircraft/{id}", model);
        }
    }
}
