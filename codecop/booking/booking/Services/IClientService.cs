using booking.common.ViewModel;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.Services
{
    public interface IClientService 
    {
        Task Create(ClientModel model);

        Task Update(string id, ClientModel model);

        Task<IEnumerable<ClientModel>> GetAll(int page, int size);

        Task<ClientModel> GetById(string id);

        Task Remove(string id);
    }
}
