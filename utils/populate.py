import pymongo
import requests
import pickle
import time
import os
from dotenv import load_dotenv
import tqdm

load_dotenv()

SPOTIFY_CLIENT_ID = os.getenv("SPOTIFY_CLIENT_ID")
SPOTIFY_CLIENT_SECRET = os.getenv("SPOTIFY_CLIENT_SECRET")
SPOTIFY_ACCESS_TOKEN = None

client = pymongo.MongoClient(os.getenv("MONGO_URI"))
db = client[os.getenv("DB_NAME")]

def refresh_access_token():
    r = requests.post("https://accounts.spotify.com/api/token", {
        "grant_type": "client_credentials",
        "client_id": SPOTIFY_CLIENT_ID,
        "client_secret": SPOTIFY_CLIENT_SECRET
    }).json()
    global SPOTIFY_ACCESS_TOKEN
    SPOTIFY_ACCESS_TOKEN = r["access_token"]

print("Connected to database")

with open("../model/artifacts/embeddings.pkl", "rb") as f:
    embeddings = pickle.load(f)

embed_ids = list(embeddings.keys())

refresh_access_token()

for i in tqdm.tqdm(range(0, (len(embed_ids) // 50) + 1)):
    search_ids = embed_ids[50 * i: (50 * i) + 50]
    search = ",".join(search_ids)
    r = requests.get(f"https://api.spotify.com/v1/tracks?ids={search}", headers={
        "Authorization": f"Bearer {SPOTIFY_ACCESS_TOKEN}"
    }).json()
    objs = []
    for j in range(len(r["tracks"])):
        song_obj = dict()
        song_obj["songid"] = search_ids[j]
        song_obj["likes"] = 0
        song_obj["songname"] = r["tracks"][j]["name"]
        song_obj["image"] = r["tracks"][j]["album"]["images"][1]["url"]
        song_obj["songartists"] = ",".join(list(map(lambda x: x["name"], r["tracks"][j]["artists"])))
        objs.append(song_obj)
    new_song = db["songs"].insert_many(objs)
    if i % 89 == 0:
        time.sleep(30)
        refresh_access_token()
        
print("Inserted all values")
