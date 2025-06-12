
using booking.common.ViewModel;
using booking.order.Abstract;
using booking.order.Controllers;
using booking.order.Model;
using Microsoft.AspNetCore.Mvc;
using Moq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace TestProject.TestOrder
{
    public class PutOrders
    {
        [Fact]
        public async Task TestPutOrderNotFoundResult()
        {
            // Arrange & Act
            var mockRepo = new Mock<IOrderRepository>();
            //var mockLogger = new Mock<ILogger<ConcertsController>>();
            var controller = new OrderController(mockRepo.Object);
            controller.ModelState.AddModelError("error", "some error");

            // Act
            var result =  controller.Put(id: "0", model: null);

            // Assert
            Assert.IsType<NotFoundResult>(result);
        }

        [Fact]
        public async Task TestPutOrderNotFoundResultId()
        {
            // Arrange
            String testId = "101";
            OrderModel order = GetTestOrders()[0];
            var mockRepo = new Mock<IOrderRepository>();
           
            var controller = new OrderController(mockRepo.Object);

            // Act
            var result =  controller.Put(testId, order);

            // Assert
            Assert.IsType<NotFoundResult>(result);
        }

        [Fact]
        public async Task TestPutOrderOkResult()
        {
            // Arrange 
            String testId = "1";
            OrderModel order = GetTestOrders()[0];
            var mockRepo = new Mock<IOrderRepository>();
            var controller = new OrderController(mockRepo.Object);

            // Act 
            var result = controller.Put(testId, order);

            // Assert 
            Assert.IsType<NotFoundResult>(result);
        }

        private List<OrderModel> GetTestOrders()
        {
            var orders = new List<OrderModel>
            {

                new OrderModel()
                {
                    Id = "1",
                    FlightId = "110",
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
