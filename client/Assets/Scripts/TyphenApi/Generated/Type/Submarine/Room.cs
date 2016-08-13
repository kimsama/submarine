// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class Room : TyphenApi.TypeBase<Room>
    {
        protected static readonly SerializationInfo<Room, long> id = new SerializationInfo<Room, long>("id", false, (x) => x.Id, (x, v) => x.Id = v);
        public long Id { get; set; }
        protected static readonly SerializationInfoForList<Room, TyphenApi.Type.Submarine.User> members = new SerializationInfoForList<Room, TyphenApi.Type.Submarine.User>("members", false, (x) => x.Members, (x, v) => x.Members = v);
        public List<TyphenApi.Type.Submarine.User> Members { get; set; }
        protected static readonly SerializationInfoForList<Room, TyphenApi.Type.Submarine.Bot> bots = new SerializationInfoForList<Room, TyphenApi.Type.Submarine.Bot>("bots", true, (x) => x.Bots, (x, v) => x.Bots = v);
        public List<TyphenApi.Type.Submarine.Bot> Bots { get; set; }
    }
}
