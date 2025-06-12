using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.ComponentModel.DataAnnotations;

namespace booking.flight.Model
{
    public class Flight
    {
        [Key]
        public string Id { get; set; }
        public string Number { get; set; }
        public string AircraftId { get; set; }
        public int FreeSeats { get; set; }
        public decimal Sum { get; set; }
        public DateTime Date { get; set; }
    }
}
