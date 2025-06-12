using booking.common.ViewModel;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.Services
{
    public interface IFlightService
    {
        Task Create(FlightModel model);

        Task Update(string id, FlightModel model);

        Task<IEnumerable<FlightModel>> GetAll(int page, int size);

        Task<FlightModel> GetById(string id);

        Task Remove(string id);
    }
}
