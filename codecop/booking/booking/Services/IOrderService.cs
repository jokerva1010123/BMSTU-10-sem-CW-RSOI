using booking.common.ViewModel;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.Services
{
    public interface IOrderService
    {
        Task Create(OrderModel model);

        Task Update(string id, OrderModel model);

        Task<IEnumerable<OrderModel>> GetAll(int page, int size);

        Task<OrderModel> GetById(string id);


        Task<OrderModel> GetByFlightId(string id);

        Task Remove(string id);

        Task RemoveByFlightId(string id);
    }
}
