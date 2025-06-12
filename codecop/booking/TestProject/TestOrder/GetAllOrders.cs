using System.Linq;
using System.Collections.Generic;

using Microsoft.AspNetCore.Mvc;
using Xunit;
using Moq;


using booking.common.ViewModel;
using booking.order.Model;
using booking.order.Abstract;
using booking.order.Controllers;

namespace TestProject.TestOrders
{
    public class GetAllOrders
    {
        [Fact]
        public void TestGetAllOrders()
        {
            var testOrders = GetTestOrders();
            var mockRepo = new Mock<IOrderRepository>();

            mockRepo.Setup(c => c.GetAll())
               .Returns(testOrders);
            var controller = new OrderController(mockRepo.Object);

            // Act
            var result = controller.GetAllOrders(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<OrderModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Equal(3, (model.Value as IEnumerable<OrderModel>).Count());
        }

        [Fact]
        public void TestGetAllClientsPagination()
        {
            int page = 2;
            int amount = 1;
            var testOrders = GetTestOrders().Skip(page * (amount - 1)).Take(amount);
            
            var mockRepo = new Mock<IOrderRepository>();
            mockRepo.Setup(c => c.GetAll())
               .Returns(testOrders);
            var controller = new OrderController(mockRepo.Object);

            // Act
            var result = controller.GetAllOrders(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<OrderModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Single((model.Value as IEnumerable<OrderModel>));
        }

        private List<Order> GetTestOrders()
        {
            var orders = new List<Order>
            {
                
                new Order()
                {
                    Id = "1",
                    FlightId = "110",
                    ClientId = "210",
                    Summ = 1000,
                    Status = 0
                },
                new Order()
                {
                    Id = "2",
                    FlightId = "120",
                    ClientId = "220",
                    Summ = 2000,
                    Status = 0
                },
                new Order()
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

