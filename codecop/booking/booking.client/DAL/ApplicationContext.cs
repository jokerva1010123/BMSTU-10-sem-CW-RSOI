using booking.client.Model;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.client.DAL
{
    public class ApplicationContext : DbContext
    {
        public DbSet<Client> Clients { get; set; }

        public ApplicationContext(DbContextOptions options) : base(options)
        {
        }
    }
}
