from fastapi import FastAPI
import httpx

app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello World"}

@app.get("/ping/{target}")
async def ping(target: str):
    response = httpx.get(f"http://{target}:8080/")
    return {"target": target, "response": response.text}
