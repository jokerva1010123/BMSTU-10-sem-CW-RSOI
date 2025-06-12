using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.client.Abstract
{
    public interface IRepository <T>
    {
        void Add(T item);
        void Delete(String id);
        T Update(T item);
        T Get(String id);
        IEnumerable<T> GetAll();
    }
}
