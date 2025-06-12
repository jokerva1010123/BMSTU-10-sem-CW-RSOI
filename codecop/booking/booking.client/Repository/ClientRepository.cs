using booking.client.Abstract;
using booking.client.DAL;
using booking.client.Model;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace booking.client.Repository
{
    public class ClientRepository : IClientRepository
    {
        private readonly ApplicationContext context;

        public ClientRepository (ApplicationContext context)
        {
            this.context = context;
        }

        public void Add(Client item)
        {
            item.Id = Guid.NewGuid().ToString();
            context.Clients.Add(item);
            context.SaveChanges();
        }

        public void Delete(String id)
        {
            var client = context.Clients.FirstOrDefault(x => x.Id == id);
            if (client != null)
            {
                context.Clients.Remove(client);
                context.SaveChanges();
            }
            
        }

        public Client Get(string id)
        {
            return context.Clients.FirstOrDefault(x => x.Id == id);
        }

        public IEnumerable<Client> GetAll()
        {
            return context.Clients.ToList();
        }

        public Client Update(Client item)
        {
            var client = context.Clients.FirstOrDefault(x => x.Id == item.Id);
            if (client != null)
            {
                client.Firstname = item.Firstname;
                client.Middlename = item.Middlename;
                client.Lastname = item.Lastname;
                client.Age = item.Age;
                context.Clients.Update(client);
                context.SaveChanges();
            }
            return client;
        }
    }
}
