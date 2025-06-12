using booking.common.ViewModel;
using System;
using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System.Linq;
using System.Threading.Tasks;

namespace booking.order.Model
{
    public class Order
    {
        [Key]
        public string Id { get; set; }
        public string FlightId { get; set; }
        public string ClientId { get; set; }
        public decimal Summ { get; set; }
        public OrderStatus Status { get; set; }
    }
}
