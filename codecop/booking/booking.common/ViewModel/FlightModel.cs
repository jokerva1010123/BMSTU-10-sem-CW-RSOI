using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.common.ViewModel
{
    public class FlightModel
    {
        public string Id;
        public string Number { get; set; }
        public string AircraftId { get; set; }
        public int FreeSeats { get; set; }
        public decimal Sum { get; set; }
        public DateTime Date { get; set; }
    }

}
