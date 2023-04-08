from fastapi import FastAPI
from pydantic import BaseModel
from numpy import dot
from numpy.linalg import norm
import pickle
import random
import requests
import warnings
import tensorflow as tf
import os
from dotenv import load_dotenv

warnings.filterwarnings("ignore")
load_dotenv()

SPOTIFY_CLIENT_ID = os.getenv("SPOTIFY_CLIENT_ID")
SPOTIFY_CLIENT_SECRET = os.getenv("SPOTIFY_CLIENT_SECRET")
SPOTIFY_ACCESS_TOKEN = None
ACTUAL_SONGS_RATIO = 0.6

class Songs(BaseModel):
    song_ids: list[str]
    song_counts: int

class AddSong(BaseModel):
    song_id: str

app = FastAPI()

@app.post("/recommend")
async def recommend(songs: Songs):
    with open("./artifacts/embeddings.pkl", "rb") as f:
        embeddings = pickle.load(f)
    song_ids = songs.song_ids
    similarities = []
    for encoding_id in list(embeddings.keys()):
        if encoding_id in song_ids:
            continue
        sim = 0
        for song_id in song_ids:
            sim += dot(embeddings[song_id], embeddings[encoding_id]) / (norm(embeddings[song_id]) * norm(embeddings[encoding_id]))
        similarities.append([sim, encoding_id])
    similarities.sort(key=lambda x: x[0], reverse=True)
    recommendations = [similarities[i][1] for i in range(min(int(songs.song_counts * ACTUAL_SONGS_RATIO), len(similarities)))]
    recommendations.extend([similarities[random.randint(0, len(similarities) - 1)][1] for _ in range(max(songs.song_counts - int(songs.song_counts * ACTUAL_SONGS_RATIO), 0))])
    return {"data": recommendations}

@app.post("/add")
async def add(song: AddSong):
    with open("./artifacts/scalers.pkl", "rb") as f:
        scalers = pickle.load(f)
    refresh_access_token()
    song_id = song.song_id
    r = requests.get(f"https://api.spotify.com/v1/audio-features/{song_id}", headers={
        "Authorization": f"Bearer {SPOTIFY_ACCESS_TOKEN}"
    }).json()
    encoding = []
    for scaler in scalers:
        if scaler in r:
            encoding.append(scalers[scaler].transform([[r[scaler]]])[0][0])
    encoder = tf.keras.models.load_model("./artifacts/encoder.h5")
    encoding = encoder.predict([encoding])[0]
    with open("./artifacts/embeddings.pkl", "rb") as f:
        embeddings = pickle.load(f)
    embeddings[song_id] = encoding
    with open("./artifacts/embeddings.pkl", "wb") as f:
        pickle.dump(embeddings, f)
    return {"song_id": song_id}

def refresh_access_token():
    r = requests.post("https://accounts.spotify.com/api/token", {
        "grant_type": "client_credentials",
        "client_id": SPOTIFY_CLIENT_ID,
        "client_secret": SPOTIFY_CLIENT_SECRET
    }).json()
    global SPOTIFY_ACCESS_TOKEN
    SPOTIFY_ACCESS_TOKEN = r["access_token"]