using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.order.Abstract
{
    public interface IRepository <T>
    {
        void Add(T item);
        void Delete(String Id);
        T Update(T item);
        T Get(String id);
        T GetbyFlightId(String flightId);
        IEnumerable<T> GetAll();
    }
}
