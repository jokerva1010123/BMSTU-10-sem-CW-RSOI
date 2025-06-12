using booking.flight.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.flight.Abstract
{
    public interface IFlightRepository: IRepository <Flight>
    {
    }
}
