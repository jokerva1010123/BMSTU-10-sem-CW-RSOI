using System.Linq;
using System.Collections.Generic;
using Microsoft.Extensions.Logging;
using Microsoft.AspNetCore.Mvc;
using Xunit;
using Moq;

using booking.flight.Controllers;
using booking.flight.Model;
using booking.flight.Abstract;
using booking.common.ViewModel;

namespace TestProject.TestControllers
{
    public class GetAllFlights
    {
        [Fact]
        public void TestGetAllFlights()
        {
            var testFlight = GetTestFlights();
            var testAircraft = GetTestAircrafts();
            
            var mockRepoFlight = new Mock<IFlightRepository>();
            var mockRepoAircraft = new Mock<IAircraftRepository>();

            mockRepoFlight.Setup(c => c.GetAll())
               .Returns(testFlight);
            mockRepoAircraft.Setup(c => c.GetAll());
            var controller = new FlightController(mockRepoFlight.Object, mockRepoAircraft.Object);

            // Act
            var result = controller.GetAllFlights(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<FlightModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Equal(3, (model.Value as IEnumerable<FlightModel>).Count());
        }

        [Fact]
        public void TestGetAllAircrafts()
        {
            var testFlight = GetTestFlights();
            var testAircraft = GetTestAircrafts();
            
            
            var mockRepoAircraft = new Mock<IAircraftRepository>();
            var mockRepoFlight = new Mock<IFlightRepository>();
            mockRepoFlight.Setup(c => c.GetAll())
               .Returns(testFlight);
            mockRepoAircraft.Setup(c => c.GetAll())
                .Returns(testAircraft);
            var controller = new AircraftController(mockRepoAircraft.Object);

            // Act
            var result = controller.GetAll(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<AircraftModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Equal(3, (model.Value as IEnumerable<AircraftModel>).Count());
        }

        [Fact]
        public void TestGetAllAircraftsPagination()
        {
            int page = 2;
            int amount = 1;
            var testFlight = GetTestFlights();
            var testAircraft = GetTestAircrafts().Skip(page * (amount - 1)).Take(amount); ;
            
            
            var mockRepoAircraft = new Mock<IAircraftRepository>();
            var mockRepoFlight = new Mock<IFlightRepository>();
            mockRepoFlight.Setup(c => c.GetAll())
               .Returns(testFlight);
            mockRepoAircraft.Setup(c => c.GetAll())
                .Returns(testAircraft); ;
            var controller = new AircraftController(mockRepoAircraft.Object);

            // Act
            var result = controller.GetAll(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<AircraftModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Single((model.Value as IEnumerable<AircraftModel>));
        }

        [Fact]
        public void TestGetAllFlightsPagination()
        {
            int page = 2;
            int amount = 1;
            var testFlight = GetTestFlights().Skip(page * (amount - 1)).Take(amount);
            
            var testAircraft = GetTestAircrafts();
           
            var mockRepoFlight = new Mock<IFlightRepository>();
            var mockRepoAircraft = new Mock<IAircraftRepository>();

            mockRepoFlight.Setup(c => c.GetAll())
               .Returns(testFlight);
            mockRepoAircraft.Setup(c => c.GetAll());
            var controller = new FlightController(mockRepoFlight.Object, mockRepoAircraft.Object);

            // Act
            var result = controller.GetAllFlights(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<FlightModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            
            Assert.Single((model.Value as IEnumerable<FlightModel>));
        }

        private List<Flight> GetTestFlights()
        {
            var flights = new List<Flight>
            {
                new Flight()
                {
                    Id = "100",
                    Number = "SU102",
                    AircraftId = "1000",
                    FreeSeats = 1,
                    Sum = 1000
                },
                new Flight()
                {
                    Id = "101",
                    Number = "AR111",
                    AircraftId = "1100",
                    FreeSeats = 10,
                    Sum = 2000
                },
                new Flight()
                {
                    Id = "102",
                    Number = "SU115",
                    AircraftId = "1200",
                    FreeSeats = 55,
                    Sum = 3000
                }
            };
            return flights;
        }


        private List<Aircraft> GetTestAircrafts()
        {
            var aircrafts = new List<Aircraft>
            {
                new Aircraft()
                {
                    Id = "1100",
                    Name = "SU102",
                    NumberOfSeats = 1,
                },
                new Aircraft()
                {
                    Id = "1200",
                    Name = "SU102",
                    NumberOfSeats = 15,
                },
                new Aircraft()
                {
                    Id = "1300",
                    Name = "SU102",
                    NumberOfSeats = 60,
                }
            };
            return aircrafts;
        }
    }
}

