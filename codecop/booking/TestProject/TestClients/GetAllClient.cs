using System.Linq;
using System.Collections.Generic;
using Microsoft.Extensions.Logging;
using Microsoft.AspNetCore.Mvc;
using Xunit;
using Moq;

using booking.client.Controllers;
using booking.client.Model;
using booking.client.Abstract;
using booking.common.ViewModel;

namespace TestProject.TestControllers
{
    public class GetAllClient
    {
        [Fact]
        public void TestGetAllClients()
        {
            var testClients = GetTestClients();
            //var mockLogger = new Mock<ILogger<ClientController>>();
            var mockRepo = new Mock<IClientRepository>();

            mockRepo.Setup(c => c.GetAll())
               .Returns(testClients);
            var controller = new ClientController(mockRepo.Object);

            // Act
            var result = controller.GetAll(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<ClientModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Equal(3, (model.Value as IEnumerable<ClientModel>).Count());
        }

        [Fact]
        public void TestGetAllClientsPagination()
        {
            int page = 2;
            int amount = 1;
            var testClients = GetTestClients().Skip(page * (amount - 1)).Take(amount);
            //var mockLogger = new Mock<ILogger<ClientController>>();
            var mockRepo = new Mock<IClientRepository>();

            mockRepo.Setup(c => c.GetAll())
               .Returns(testClients);
            var controller = new ClientController(mockRepo.Object);

            // Act
            var result = controller.GetAll(0, 0);

            // Assert
            var actionResult = Assert.IsType<ActionResult<IEnumerable<ClientModel>>>(result);
            var model = Assert.IsType<OkObjectResult>(actionResult.Result);
            Assert.Single((model.Value as IEnumerable<ClientModel>));
        }

        private List<Client> GetTestClients()
        {
            var clients = new List<Client>
            {
                new Client()
                {
                    Id = "100",
                    Firstname = "Анна",
                    Middlename = "Михайловна",
                    Lastname = "Козакова",
                    Age = 43
                },
                new Client()
                {
                    Id = "101",
                    Firstname = "Макар",
                    Middlename = "Брониславович",
                    Lastname = "Румянцев",
                    Age = 24
                },
                new Client()
                {
                    Id = "102",
                    Firstname = "Наталия",
                    Middlename = "Мироновна",
                    Lastname = "Васильева",
                    Age = 30
                }
            };
            return clients;
        }

    }
}

