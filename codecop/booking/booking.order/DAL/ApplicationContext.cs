using booking.order.Model;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.order.DAL
{
    public class ApplicationContext : DbContext
    {
        public DbSet<Order> Orders { get; set; }

        public ApplicationContext(DbContextOptions options) : base(options)
        {
        }
    }
}
