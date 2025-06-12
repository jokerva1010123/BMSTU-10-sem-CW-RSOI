using booking.order.Abstract;
using booking.order.Controllers;
using booking.order.Model;
using booking.order.Repository;
using booking.common.ViewModel;
using booking.order.Model;
using Microsoft.AspNetCore.Mvc;
using Microsoft.EntityFrameworkCore;
using Microsoft.Extensions.Logging;
using Moq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace TestProject.TestOrders
{
    public class CreateOrder
    {
        
        [Fact]
        public async Task TestCreateOrder()
        {
            // Arrange
            string testId = "1";
            Order order = GetTestOrders().FirstOrDefault(p => p.Id == testId);
            OrderModel ordermodel = GetTestOrdersModels().FirstOrDefault(p => p.Id == testId);

            var mockRepo = new Mock<IOrderRepository>();
            mockRepo.Setup(c => c.Add(order));
            
            var controller = new OrderController(mockRepo.Object);

            // Act
            var result =  controller.Post(ordermodel);

            // Assert
            var actionResult = Assert.IsType<OkResult>(result);
           
            var model = Assert.IsType<OkResult>(actionResult);
            
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

        private List<OrderModel> GetTestOrdersModels()
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
