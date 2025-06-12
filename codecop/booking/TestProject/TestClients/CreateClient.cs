using booking.client.Abstract;
using booking.client.Controllers;
using booking.client.Model;
using booking.client.Repository;
using booking.common.ViewModel;
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

namespace TestProject.TestClients
{
    public class CreateClient
    {
        
        [Fact]
        public async Task TestCreateClient()
        {
            // Arrange
            string testId = "100";
            Client client = GetTestClients().FirstOrDefault(p => p.Id == testId);
            ClientModel clientmodel = GetTestClientsModels().FirstOrDefault(p => p.Id == testId);

            var mockRepo = new Mock<IClientRepository>();
            mockRepo.Setup(c => c.Add(client));
           
            var controller = new ClientController(mockRepo.Object);

            // Act
            var result =  controller.Post(clientmodel);

            // Assert
            var actionResult = Assert.IsType<OkResult>(result);
           
            var model = Assert.IsType<OkResult>(actionResult);
            /*
            Assert.Equal(client.Id, model.Id);
            Assert.Equal(client.Firstname, model.Firstname);
            Assert.Equal(client.Middlename, model.Middlename);
            Assert.Equal(client.Lastname, model.Lastname);
            Assert.Equal(client.Age, model.Age);
            */
        }
        


        private List<Client> GetTestClients()
        {
            var clients = new List<Client>();
            clients.Add(new Client()
            {
                Id = "100",
                Firstname = "Анна",
                Middlename = "Михайловна",
                Lastname = "Козакова",
                Age = 43
            });
            clients.Add(new Client()
            {
                Id = "101",
                Firstname = "Макар",
                Middlename = "Брониславович",
                Lastname = "Румянцев",
                Age = 24
            });
            clients.Add(new Client()
            {
                Id = "102",
                Firstname = "Наталия",
                Middlename = "Мироновна",
                Lastname = "Васильева",
                Age = 30
            });

            return clients;
        }

        private List<ClientModel> GetTestClientsModels()
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

    }
}
