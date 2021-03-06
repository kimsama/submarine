// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine.Configuration
{
    public partial class ServerDatabaseObject : TyphenApi.TypeBase<ServerDatabaseObject>
    {
        protected static readonly SerializationInfo<ServerDatabaseObject, string> host = new SerializationInfo<ServerDatabaseObject, string>("host", false, (x) => x.Host, (x, v) => x.Host = v);
        public string Host { get; set; }
        protected static readonly SerializationInfo<ServerDatabaseObject, long> port = new SerializationInfo<ServerDatabaseObject, long>("port", false, (x) => x.Port, (x, v) => x.Port = v);
        public long Port { get; set; }
        protected static readonly SerializationInfo<ServerDatabaseObject, string> user = new SerializationInfo<ServerDatabaseObject, string>("user", false, (x) => x.User, (x, v) => x.User = v);
        public string User { get; set; }
        protected static readonly SerializationInfo<ServerDatabaseObject, string> password = new SerializationInfo<ServerDatabaseObject, string>("password", false, (x) => x.Password, (x, v) => x.Password = v);
        public string Password { get; set; }
    }
}
