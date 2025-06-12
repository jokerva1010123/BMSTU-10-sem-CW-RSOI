using booking.client.Abstract;
using booking.client.Controllers;
using booking.client.Model;
using booking.common.ViewModel;
using Microsoft.AspNetCore.Mvc;
using Moq;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using Xunit;

namespace TestProject.TestLients
{
    public class PutClient
    {
        [Fact]
        public async Task TestPutClientNotFoundResult()
        {
            // Arrange & Act
            var mockRepo = new Mock<IClientRepository>();
            //var mockLogger = new Mock<ILogger<ConcertsController>>();
            var controller = new ClientController(mockRepo.Object);
            controller.ModelState.AddModelError("error", "some error");

            // Act
            var result =  controller.Put(id: "0", model: null);

            // Assert
            Assert.IsType<NotFoundResult>(result);
        }

        [Fact]
        public async Task TestPutClientNotFoundResultId()
        {
            // Arrange
            String testId = "101";
            ClientModel clientmodel = GetTestClientsModels()[0];
            Client client = GetTestClients()[0];
            var mockRepo = new Mock<IClientRepository>();
           
            var controller = new ClientController(mockRepo.Object);

            // Act
            var result =  controller.Put(testId, clientmodel);

            // Assert
            Assert.IsType<NotFoundResult>(result);
        }

        [Fact]
        public async Task TestPutClientOkResult()
        {
            // Arrange 
            String testId = "100";
            ClientModel clientmodel = GetTestClientsModels()[0];
            Client client = GetTestClients()[0];
            var mockRepo = new Mock<IClientRepository>();
            mockRepo = new Mock<IClientRepository>();
            mockRepo.Setup(c => c.Update(client))
            .Returns(GetTestClients()[0]);
            mockRepo.Setup(c => c.Get(testId))
            .Returns(new Client() { Id = testId });

            var controller = new ClientController(mockRepo.Object);

            // Act 
            var result = controller.Put(testId, clientmodel);

            // Assert 
            Assert.IsType<OkResult>(result);
        }



        private List<ClientModel> GetTestClientsModels()
        {
            var clients = new List<ClientModel>
            {
                new ClientModel()
                {
                    Id = "100",
                    Firstname = "Анна",
                    Middlename = "Михайловна",
                    Lastname = "Козакова",
                    Age = 43
                },
                new ClientModel()
                {
                    Id = "101",
                    Firstname = "Макар",
                    Middlename = "Брониславович",
                    Lastname = "Румянцев",
                    Age = 24
                },
                new ClientModel()
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
