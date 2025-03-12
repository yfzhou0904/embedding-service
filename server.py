from flask import Flask, request, jsonify
from transformers import AutoTokenizer, AutoModel
import torch
import os
from dotenv import load_dotenv

# Load environment variables from .env file if it exists
load_dotenv()

app = Flask(__name__)

tokenizer = AutoTokenizer.from_pretrained('sentence-transformers/all-MiniLM-L6-v2')
model = AutoModel.from_pretrained('sentence-transformers/all-MiniLM-L6-v2')

@app.route('/embedding/text', methods=['POST'])
def get_text_embedding():
    data = request.json
    text = data['text']
    inputs = tokenizer(text, return_tensors='pt', padding=True, truncation=True, max_length=512)
    with torch.no_grad():
        outputs = model(**inputs)
    embedding = outputs.last_hidden_state.mean(dim=1).squeeze().tolist()
    return jsonify({'embedding': embedding})

if __name__ == '__main__':
    listen_address = os.environ.get('LISTEN_ADDRESS', '0.0.0.0:80')
    host, port_str = listen_address.split(':')
    port = int(port_str)
    app.run(host=host, port=port)
