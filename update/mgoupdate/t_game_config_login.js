conn = new Mongo();
db = conn.getDB("test");

db.getCollection('t_game_config_login').drop();

db.getCollection('t_game_config_login').insert({
    "CurVersion": 5.0,
    "BaseDownloadUrl": ""
});
