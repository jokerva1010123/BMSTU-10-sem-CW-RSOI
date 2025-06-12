using booking.order.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.order.Abstract
{
    public interface IOrderRepository : IRepository<Order>
    {
        void DeleteByFlightId(string id);

    }
}
