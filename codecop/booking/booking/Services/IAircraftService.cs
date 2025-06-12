using booking.common.ViewModel;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.Services
{
    public interface IAircraftService
    {
        Task Create(AircraftModel model);

        Task Update(string id, AircraftModel model);

        Task<IEnumerable<AircraftModel>> GetAll(int page, int size);

        Task<AircraftModel> GetById(string id);

        Task Remove(string id);
    }
}
