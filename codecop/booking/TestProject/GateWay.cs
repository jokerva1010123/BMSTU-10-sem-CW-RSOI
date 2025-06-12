using booking.common.ViewModel;
using booking.Controllers;
using booking.Services;
using Microsoft.AspNetCore.Mvc;
using Moq;
using System;
using System.Collections.Generic;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;
using Xunit;
using static booking.Controllers.BookingController;

namespace TestProject
{
    public class GateWay
    {
        [Fact]
        public async Task TestGetAll()
        {
            int testId = 1;

            //var mockLogger = new Mock<ILogger<ConcertsController»(); 
            var mockClientFactory = new Mock<IHttpClientFactory>();
            var mockServiceClient = new Mock<IClientService>();
            var mockServiceOrder = new Mock<IOrderService>();
            var mockServiceFlight = new Mock<IFlightService>();
            var mockServiceAircraft = new Mock<IAircraftService>();

            mockServiceClient.Setup(c => c.GetAll(1, 1));
            //.Returns(Task.FromResult(GetTestClients()[0])); 
            mockServiceOrder.Setup(c => c.GetAll(1, 1));
            mockServiceFlight.Setup(c => c.GetAll(1, 1));
            mockServiceAircraft.Setup(c => c.GetAll(1, 1));


            var controller = new BookingController(mockClientFactory.Object,
            mockServiceOrder.Object, mockServiceClient.Object, mockServiceFlight.Object, mockServiceAircraft.Object);

            // Act 
            var result = await controller.GetAll();

            // Assert 
            var requestResult = Assert.IsType<OkObjectResult>(result);
            var model = Assert.IsType<Res>(requestResult.Value);
            Assert.NotNull(model);
        }

        [Fact]
        public async Task AddOrderGatewayActionResult()
        {
            String testId = "100";
            bool success = true;
            var mockClientFactory = new Mock<IHttpClientFactory>();
            var mockServiceClient = new Mock<IClientService>();
            var mockServiceOrder = new Mock<IOrderService>();
            var mockServiceFlight = new Mock<IFlightService>();
            var mockServiceAircraft = new Mock<IAircraftService>();
            OrderModel order = GetTestOrders()[0];
            mockServiceOrder.Setup(c => c.Create(GetTestOrders()[0]));
            //.ReturnsAsync((true, GetTestConcerts()[0])); 
            mockServiceFlight.Setup(c => c.Update(testId, GetTestFlights()[0]));
            mockServiceFlight.Setup((c) => c.GetById(testId))
            .Returns(Task.FromResult(GetTestFlights()[0]));
            //mockServiceAircraft.Setup(c => c.GetAll(1, 1)); 


            var controller = new BookingController(mockClientFactory.Object,
            mockServiceOrder.Object, mockServiceClient.Object, mockServiceFlight.Object, mockServiceAircraft.Object);

            // Act 
            var result = await controller.AddOrder(order);

            // Assert 
            var requestResult = Assert.IsType<OkResult>(result);

        }




        [Fact]
        public async Task DeleteFlightGatewayBadrequestResult()
        {
            string testId = "100";
            var mockClientFactory = new Mock<IHttpClientFactory>();
            var mockServiceClient = new Mock<IClientService>();
            var mockServiceOrder = new Mock<IOrderService>();
            var mockServiceFlight = new Mock<IFlightService>();
            var mockServiceAircraft = new Mock<IAircraftService>();
            OrderModel order = GetTestOrders()[0];
            mockServiceOrder.Setup(c => c.RemoveByFlightId(testId))
            .Returns(Task.FromResult(0));
            mockServiceFlight.Setup(c => c.Remove(testId))
            .Returns(Task.FromResult(0));

            var controller = new BookingController(mockClientFactory.Object,
            mockServiceOrder.Object, mockServiceClient.Object, mockServiceFlight.Object, mockServiceAircraft.Object);

            // Act 
            var result = await controller.DeleteFlight(testId);

            // Assert 
            var requestResult = Assert.IsType<OkResult>(result);
        }

        private List<ClientModel> GetTestClients()
        {
            var clients = new List<ClientModel>();
            clients.Add(new ClientModel()
            {
                Id = "100",
                Firstname = "Анна",
                Middlename = "Михайловна",
                Lastname = "Козакова",
                Age = 43
            });
            clients.Add(new ClientModel()
            {
                Id = "101",
                Firstname = "Макар",
                Middlename = "Брониславович",
                Lastname = "Румянцев",
                Age = 24
            });
            clients.Add(new ClientModel()
            {
                Id = "102",
                Firstname = "Наталия",
                Middlename = "Мироновна",
                Lastname = "Васильева",
                Age = 30
            });

            return clients;
        }

        private List<FlightModel> GetTestFlights()
        {
            var flights = new List<FlightModel>
            {
                new FlightModel()
                {
                    Id = "100",
                    Number = "SU102",
                    AircraftId = "1000",
                    FreeSeats = 1,
                    Sum = 1000
                },
                new FlightModel()
                {
                    Id = "101",
                    Number = "AR111",
                    AircraftId = "1100",
                    FreeSeats = 10,
                    Sum = 2000
                },
                new FlightModel()
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


        private List<AircraftModel> GetTestAircrafts()
        {
            var aircrafts = new List<AircraftModel>
            {
                new AircraftModel()
                {
                    Id = "1100",
                    Name = "SU102",
                    NumberOfSeats = 1,
                },
                new AircraftModel()
                {
                    Id = "1200",
                    Name = "SU102",
                    NumberOfSeats = 15,
                },
                new AircraftModel()
                {
                    Id = "1300",
                    Name = "SU102",
                    NumberOfSeats = 60,
                }
            };
            return aircrafts;
        }

        private List<OrderModel> GetTestOrders()
        {
            var orders = new List<OrderModel>
            {

                new OrderModel()
                {
                    Id = "1",
                    FlightId = "100",
                    ClientId = "210",
                    Summ = 1000,
                    Status = 0
                },
                new OrderModel()
                {
                    Id = "2",
                    FlightId = "120",
                    ClientId = "220",
                    Summ = 2000,
                    Status = 0
                },
                new OrderModel()
                {
                    Id = "3",
                    FlightId = "130",
                    ClientId = "230",
                    Summ = 3000,
                    Status = 0
                }
            };
            return orders;
        }
    }
}
