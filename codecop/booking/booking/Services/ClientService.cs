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
    public class ClientService : IClientService
    {
        private readonly HttpClient _httpClient;
        private readonly ILogger<ClientService> _logger;
        private readonly UrlHosts _urls;


        public ClientService(HttpClient httpClient, ILogger<ClientService> logger, IOptions<UrlHosts> config)
        {
            _httpClient = httpClient;
            _logger = logger;
            _urls = config.Value;
        }

        public async Task Create(ClientModel model)
        {
            await _httpClient.PostAsJsonAsync(_urls.Client + "/api/client", model);
        }

        public async Task<IEnumerable<ClientModel>> GetAll(int page, int size)
        {
            var data = await _httpClient.GetStringAsync(_urls.Client + $"/api/client?page={page}&size={size}");
            var clients = !string.IsNullOrEmpty(data)
                ? JsonConvert.DeserializeObject<IEnumerable<ClientModel>>(data)
                : null;
            return clients;
        }

        public async Task<ClientModel> GetById(string id)
        {
            var data = await _httpClient.GetStringAsync(_urls.Client + $"/api/client/{id}");
            var client = !string.IsNullOrEmpty(data) ? JsonConvert.DeserializeObject<ClientModel>(data) : null;
            return client;
        }

        public async Task Remove(string id)
        {
            await _httpClient.DeleteAsync(_urls.Client + $"/api/client/{id}");
        }

        public async Task Update(string id, ClientModel model)
        {
            await _httpClient.PutAsJsonAsync(_urls.Client + $"/api/client/{id}", model);
        }
    }
}
