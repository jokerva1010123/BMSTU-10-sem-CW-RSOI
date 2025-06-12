using booking.flight.Abstract;
using booking.flight.DAL;
using booking.flight.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.flight.Repository
{
    public class FlightRepository : IFlightRepository
    {
        private ApplicationContext context;

        public FlightRepository(ApplicationContext context)
        {
            this.context = context;
        }

        public void Add(Flight item)
        {
            context.Flights.Add(item);
            context.SaveChanges();
        }

        public void Delete(String id)
        {
            var flight = context.Flights.FirstOrDefault(x => x.Id == id);
            if (flight != null)
            {
                context.Flights.Remove(flight);
                context.SaveChanges();
            }
        }

        public Flight Get(string id)
        {
            return context.Flights.FirstOrDefault(x => x.Id == id);
        }

        public IEnumerable<Flight> GetAll()
        {
            return context.Flights.ToList();
        }

        public Flight Update(Flight item)
        {
            var flight = context.Flights.FirstOrDefault(x => x.Id == item.Id);
            if (flight != null)
            {
                flight.AircraftId = item.AircraftId;
                flight.Date = item.Date;
                flight.Number = item.Number;
                context.Flights.Update(flight);
                context.SaveChanges();
            }
            return flight;
        }
    }
}
