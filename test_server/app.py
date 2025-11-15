import os
from flask import Flask, Response

app = Flask(__name__)

status_code = int(os.getenv("HEALTH_CODE", "200"))

@app.route("/health")
def health():
    if status_code == 200:
        return "OK", 200
    else:
        return Response(f"Error {status_code}", status=status_code)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=80)
