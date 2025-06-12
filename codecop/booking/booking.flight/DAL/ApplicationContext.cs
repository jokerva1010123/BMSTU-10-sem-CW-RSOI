using booking.flight.Model;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.flight.DAL
{
    public class ApplicationContext : DbContext
    {
        public DbSet<Flight> Flights { get; set; }

        public DbSet<Aircraft> Aircrafts { get; set; }

        public ApplicationContext(DbContextOptions options) : base(options)
        {
        }
    }

    
}
