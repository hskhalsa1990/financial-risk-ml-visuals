
import json, sys
from pathlib import Path
import joblib
import pandas as pd
from flask import Flask, jsonify, request

PREFIX = sys.argv[1] if len(sys.argv) > 1 else "loan_score_lr_v1"
model = joblib.load(f"{PREFIX}_model.joblib")
meta = json.loads(Path(f"{PREFIX}_metadata.json").read_text())
FEATURES = meta["features"]

app = Flask(__name__)

@app.get("/health")
def health():
    return jsonify(status="ok", model=meta["name"], version=meta["version"])

@app.get("/score")
def score():
    try:
        values = [int(request.args[f]) for f in FEATURES]
    except (KeyError, ValueError):
        return jsonify(error="integer query parameters required: " + ", ".join(FEATURES)), 400
    if any(v < 0 for v in values):
        return jsonify(error="values must be non-negative"), 400
    p = float(model.predict_proba(pd.DataFrame([values], columns=FEATURES))[0, 1])
    return jsonify(score=round(p, 4), model_version=meta["version"])

if __name__ == "__main__":
    app.run(host="127.0.0.1", port=5000)
