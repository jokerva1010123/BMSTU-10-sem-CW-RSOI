using booking.flight.Abstract;
using booking.flight.DAL;
using booking.flight.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.flight.Repository
{
    public class AircraftRepository : IAircraftRepository
    {
        private ApplicationContext context;

        public AircraftRepository(ApplicationContext context)
        {
            this.context = context;
        }

        public void Add(Aircraft item)
        {
            context.Aircrafts.Add(item);
            context.SaveChanges();
        }

        public void Delete(String id)
        {
            var aircraft = context.Aircrafts.FirstOrDefault(x => x.Id == id);
            if (aircraft != null)
            {
                context.Aircrafts.Remove(aircraft);
                context.SaveChanges();
            }
        }

        public Aircraft Get(string id)
        {
            return context.Aircrafts.FirstOrDefault(x => x.Id == id);
        }

        public IEnumerable<Aircraft> GetAll()
        {
            return context.Aircrafts.ToList();
        }

        public Aircraft Update(Aircraft item)
        {
            var aircraft = context.Aircrafts.FirstOrDefault(x => x.Id == item.Id);
            if (aircraft != null)
            {
                aircraft.Name = item.Name;
                aircraft.NumberOfSeats = item.NumberOfSeats;
                context.Aircrafts.Update(aircraft);
                context.SaveChanges();
            }
            return aircraft;
        }
    }
}
