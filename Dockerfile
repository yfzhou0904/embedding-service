FROM python:3.12
WORKDIR /app

COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

COPY server.py .

EXPOSE 80
CMD ["python3", "server.py"]
